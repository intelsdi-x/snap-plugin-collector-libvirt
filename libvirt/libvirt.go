/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

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
	"errors"
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
	Version = 2
	// Type of plugin
	Type = plugin.CollectorPluginType
)

func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

type Libvirt struct {
}

// Default qemu libvirt URI
var defaultURI = "qemu:///system"

func NewLibvirtCollector() *Libvirt {
	return &Libvirt{}

}

func joinNamespace(ns []string) string {
	return "/" + strings.Join(ns, "/")
}

func (p *Libvirt) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	cpure := regexp.MustCompile(`^/libvirt/.*/.*/cpu/.*`)
	memre := regexp.MustCompile(`^/libvirt/.*/.*/mem/.*`)
	netre := regexp.MustCompile(`^/libvirt/.*/.*/net/.*`)
	diskre := regexp.MustCompile(`^/libvirt/.*/.*/disk/.*`)
	metrics := make([]plugin.PluginMetricType, len(mts))
	conn, err := libvirt.NewVirConnection(getHypervisorUri(mts[0].Config().Table()))

	if err != nil {
		return nil, err
	}
	defer conn.CloseConnection()

	for i, p := range mts {

		domainName, err := namespacetoDomain(p.Namespace())
		if err != nil {
			return nil, err
		}
		dom, err := conn.LookupDomainByName(domainName)
		if err != nil {
			return nil, err
		}
		defer dom.Free()

		ns := joinNamespace(p.Namespace())
		switch {
		case memre.MatchString(ns):
			metric, err := memStat(p.Namespace(), dom)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric

		case cpure.MatchString(ns):
			metric, err := cpuTimes(p.Namespace(), dom)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		case netre.MatchString(ns):
			metric, err := interfaceStat(p.Namespace(), dom)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric
		case diskre.MatchString(ns):
			metric, err := diskStat(p.Namespace(), dom)
			if err != nil {
				return nil, err
			}
			metrics[i] = *metric

		}
		metrics[i].Source_, _ = conn.GetHostname()

	}
	return metrics, nil
}

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

func (p *Libvirt) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {

	conn, err := libvirt.NewVirConnection(getHypervisorUri(cfg.Table()))

	if err != nil {
		handleErr(err)
	}

	metrics := make([]plugin.PluginMetricType, 0)

	domains, err := conn.ListDomains()
	if err != nil {
		handleErr(err)
	}

	hostname, err := conn.GetHostname()
	if err != nil {
		handleErr(err)
	}

	defer conn.CloseConnection()
	for j := 0; j < domainCount(domains); j++ {
		dom, err := conn.LookupDomainById(domains[j])
		if err != nil {
			handleErr(err)
		}
		defer dom.Free()

		net_mts, err := getNetMetricTypes(dom, hostname)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, net_mts...)
		cpu_mts, err := getCpuMetricTypes(dom, hostname)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, cpu_mts...)
		mem_mts, err := getMemMetricTypes(dom, hostname)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, mem_mts...)
		disk_mts, err := getDiskMetricTypes(dom, hostname)
		if err != nil {
			handleErr(err)
		}
		metrics = append(metrics, disk_mts...)
	}
	return metrics, nil
}

func domainCount(domains []uint32) int {
	return len(domains)
}
func namespacetoDomain(namespace []string) (string, error) {
	if len(namespace) > 2 {
		return namespace[2], nil
	}
	err := errors.New("Namespace is too short")
	return "", err

}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getHypervisorUri(cfg map[string]ctypes.ConfigValue) string {
	if cfg_uri, ok := cfg["uri"]; ok {
		return cfg_uri.(ctypes.ConfigValueStr).Value

	} else {
		return defaultURI
	}
}
