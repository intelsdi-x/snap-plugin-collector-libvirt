/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015-2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package libvirt

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
	"github.com/sandlbn/libvirt-go"
)

const (
	// Name of plugin
	Name = "libvirt"
	// Version of plugin
	Version = 7
	// Type of plugin
	Type = plugin.CollectorPluginType
)

// Meta declaration for plugin
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

// Libvirt type
type Libvirt struct {
}

// Default qemu libvirt URI
var defaultURI = "qemu:///system"

// NewLibvirtCollector returns new instance of collector
func NewLibvirtCollector() *Libvirt {
	return &Libvirt{}

}

func joinNamespace(ns []string) string {
	return "/" + strings.Join(ns, "/")
}

func addTags(hostname string, domainname string) map[string]string {
	var tags map[string]string
	tags = make(map[string]string)
	tags["hostname"] = hostname
	tags["domainame"] = domainname
	return tags

}

func metricReported(mts []plugin.MetricType, ns string) bool {
	for _, mt := range mts {
		if mt.Namespace().String() == ns {
			return true
		}
	}
	return false
}

// CollectMetrics returns collected metrics
func (p *Libvirt) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	metrics := []plugin.MetricType{}
	conn, err := libvirt.NewVirConnection(getHypervisorURI(mts[0].Config().Table()))

	if err != nil {
		return nil, err
	}
	defer conn.CloseConnection()

	for _, p := range mts {

		ns := p.Namespace()
		if ns.Strings()[1] == "*" {
			domains, err := conn.ListDomains()
			if err != nil {
				return metrics, err
			}
			for j := 0; j < domainCount(domains); j++ {
				dom, err := conn.LookupDomainById(domains[j])
				if err != nil {
					return metrics, err
				}
				defer dom.Free()
				metric, err := processMetric(ns.String(), dom, p)
				if err != nil {
					return metrics, err
				}
				if metricReported(metrics, metric.Namespace().String()) == false {
					hostname, err := conn.GetHostname()
					if err != nil {
						return metrics, err
					}
					domainname := p.Namespace()[1]
					metric.Tags_ = addTags(hostname, domainname.Value)
					metrics = append(metrics, metric)
				}

			}
		} else {

			if metricReported(metrics, p.Namespace().String()) == false {
				domainName, err := namespacetoDomain(p.Namespace().Strings())
				if err != nil {
					return nil, err
				}
				dom, err := conn.LookupDomainByName(domainName)
				if err != nil {
					return nil, err
				}
				defer dom.Free()
				metric, err := processMetric(ns.String(), dom, p)
				hostname, err := conn.GetHostname()
				if err != nil {
					return metrics, err
				}
				domainname := p.Namespace()[1]
				metric.Tags_ = addTags(hostname, domainname.Value)
				metrics = append(metrics, metric)
			}
		}

	}
	return metrics, err
}

func processMetric(ns string, dom libvirt.VirDomain, p plugin.MetricType) (plugin.MetricType, error) {
	cpure := regexp.MustCompile(`^/libvirt/.*/cpu/.*`)
	memre := regexp.MustCompile(`^/libvirt/.*/mem/.*`)
	netre := regexp.MustCompile(`^/libvirt/.*/net/.*`)
	diskre := regexp.MustCompile(`^/libvirt/.*/disk/.*`)

	switch {
	case memre.MatchString(ns):
		metric, err := memStat(p.Namespace().Strings(), dom)
		return *metric, err

	case cpure.MatchString(ns):
		metric, err := cpuTimes(p.Namespace().Strings(), dom)
		return *metric, err

	case netre.MatchString(ns):
		metric, err := interfaceStat(p.Namespace().Strings(), dom)
		return *metric, err

	case diskre.MatchString(ns):
		metric, err := diskStat(p.Namespace().Strings(), dom)
		return *metric, err

	}
	return plugin.MetricType{}, fmt.Errorf("Failed to process metric, unknown type %s", ns)
}

// GetConfigPolicy returns a config policy
func (p *Libvirt) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	config := cpolicy.NewPolicyNode()

	uri, err := cpolicy.NewStringRule("uri", false, "qemu:///system")
	handleErr(err)
	uri.Description = "Libvirt uri"
	config.Add(uri)

	cp.Add([]string{""}, config)
	return cp, nil

}

// GetMetricTypes returns metric types that can be collected
func (p *Libvirt) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {

	conn, err := libvirt.NewVirConnection(getHypervisorURI(cfg.Table()))
	defer conn.CloseConnection()

	if err != nil {
		handleErr(err)
	}

	var metrics []plugin.MetricType

	domains, err := conn.ListDomains()
	if err != nil {
		handleErr(err)
	}

	for j := 0; j < domainCount(domains); j++ {
		dom, err := conn.LookupDomainById(domains[j])
		if err != nil {
			handleErr(err)
		}
		defer dom.Free()

		netMts, err := getNetMetricTypes(dom)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, netMts...)
		cpuMts, err := getCPUMetricTypes(dom)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, cpuMts...)
		memMts, err := getMemMetricTypes(dom)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, memMts...)
		diskMts, err := getDiskMetricTypes(dom)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, diskMts...)
	}
	return metrics, nil
}

func domainCount(domains []uint32) int {
	return len(domains)
}
func namespacetoDomain(namespace []string) (string, error) {
	if len(namespace) > 2 {
		return namespace[1], nil
	}
	return "", fmt.Errorf("Namespace is too short")

}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getHypervisorURI(cfg map[string]ctypes.ConfigValue) string {
	if cfgURI, ok := cfg["uri"]; ok {
		return cfgURI.(ctypes.ConfigValueStr).Value

	}
	return defaultURI
}
