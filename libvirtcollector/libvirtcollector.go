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

package libvirtcollector

import (
	wrapper "github.com/intelsdi-x/snap-plugin-collector-libvirt/libvirt"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/sandlbn/libvirt-go"
)

const (
	// Name of plugin
	Name = "libvirt"
	// Vendor  prefix
	Vendor = "intel"
	// Plugin plugin name
	Plugin = "libvirt"
	// Version of plugin
	Version            = 15
	nsDomainPosition   = 2
	nsMetricPostion    = 3
	nsDevicePosition   = 4
	nsSubMetricPostion = 5
)

var nsTypes = struct {
	cpu     []string
	memory  []string
	disk    []string
	network []string
	all     []string
}{
	cpu:     []string{"cputime"},
	memory:  []string{"mem", "max", "swap_in", "swap_out", "major_fault", "min_fault", "unused", "available", "actual_balloon", "rss", "nr"},
	disk:    []string{"wrreq", "rdreq", "wrbytes", "rdbytes"},
	network: []string{"rxbytes", "rxpackets", "rxerrs", "rxdrop", "txbytes", "txpackets", "txerrs", "txdrop"},
	all:     []string{"cpu", "memory", "disk", "network"},
}

// LibvirtCollector type
type LibvirtCollector struct {
}

// Default qemu libvirt URI
var defaultURI = "qemu:///system"

// CollectMetrics returns collected metrics
func (LibvirtCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	meta := make(map[string]string)
	uri, err := mts[0].Config.GetString("uri")
	if err != nil {
		return nil, err
	}
	nova, err := mts[0].Config.GetBool("nova")
	if err != nil {
		return nil, err
	}
	conn, err := libvirt.NewVirConnectionReadOnly(uri)
	if err != nil {
		return nil, err
	}
	defer conn.UnrefAndCloseConnection()

	ids, err := getInstances(conn, mts)
	if err != nil {
		return metrics, err
	}

	netCount, netMetrics := filterNamespace("network", mts)
	diskCount, diskMetrics := filterNamespace("disk", mts)
	_, cpuMetrics := filterNamespace("cpu", mts)
	_, memoryMetrics := filterNamespace("memory", mts)

	var netCounters map[string]libvirt.VirDomainInterfaceStats
	var diskCounters map[string]libvirt.VirDomainBlockStats

	for _, id := range ids {
		if nova {
			meta, err = wrapper.GetNovaMetadata(id)
			if err != nil {
				return metrics, err
			}
		}
		if netCount > 0 {
			netCounters, err = wrapper.GetNetworkStatistics(id)
			if err != nil {
				return metrics, err
			}
		}
		for _, mt := range netMetrics {
			ns := copyNamespace(mt)
			if ns[nsDomainPosition].IsDynamic() {
				ns[nsDomainPosition].Value, err = id.GetName()
			}

			for k, v := range netCounters {
				newNamespace := copyNamespaceElements(ns)
				if newNamespace[nsDevicePosition].IsDynamic() {
					newNamespace[nsDevicePosition].Value = k
				}
				if newNamespace[nsDevicePosition].Value == k {
					var value int64
					switch ns[nsSubMetricPostion].Value {
					case "rxbytes":
						value = v.RxBytes
					case "rxpackets":
						value = v.RxPackets
					case "rxerrs":
						value = v.RxErrs
					case "rxdrop":
						value = v.RxDrop
					case "txbytes":
						value = v.TxBytes
					case "txpackets":
						value = v.TxPackets
					case "txerrs":
						value = v.TxErrs
					case "txdrop":
						value = v.TxDrop
					}
					if !metricStored(metrics, newNamespace) {
						metrics = append(metrics, createNamespace(mt, value, newNamespace, meta))
					}
				}
			}

		}

		if diskCount > 0 {
			diskCounters, err = wrapper.GetBlockStatistics(id)
			if err != nil {
				return metrics, err
			}
		}

		for _, mt := range diskMetrics {
			ns := copyNamespace(mt)
			if ns[nsDomainPosition].IsDynamic() {
				ns[nsDomainPosition].Value, err = id.GetName()
				if err != nil {
					return metrics, err
				}
			}

			for k, v := range diskCounters {
				newNamespace := copyNamespaceElements(ns)
				if newNamespace[nsDevicePosition].IsDynamic() {
					newNamespace[nsDevicePosition].Value = k
				}
				if newNamespace[nsDevicePosition].Value == k {
					var value int64
					switch ns[nsSubMetricPostion].Value {
					case "wrreq":
						value = v.RdReq
					case "rdreq":
						value = v.RdReq
					case "wrbytes":
						value = v.WrBytes
					case "rdbytes":
						value = v.RdBytes
					}
					if !metricStored(metrics, newNamespace) {
						metrics = append(metrics, createNamespace(mt, value, newNamespace, meta))
					}
				}

			}
		}
		for _, mt := range memoryMetrics {
			ns := copyNamespace(mt)
			if ns[nsDomainPosition].IsDynamic() {
				ns[nsDomainPosition].Value, err = id.GetName()

				if err != nil {
					return metrics, err
				}
			}
			memKey := ns[nsDevicePosition].Value
			memoryStat, err := wrapper.GetMemoryStatistics(id, memKey)
			if err != nil {
				return metrics, err
			}
			if !metricStored(metrics, ns) {
				metrics = append(metrics, createNamespace(mt, memoryStat[memKey], ns, meta))
			}
		}
		for _, mt := range cpuMetrics {
			ns := copyNamespace(mt)
			if ns[nsDomainPosition].IsDynamic() {
				ns[nsDomainPosition].Value, err = id.GetName()
				if err != nil {
					return metrics, err
				}
			}
			secondlastElement := len(ns) - 2
			if ns[secondlastElement].IsDynamic() {
				vcpuTime, err := wrapper.GetVCPUStatistics(id)
				if err != nil {
					return metrics, err
				}
				for k, v := range vcpuTime {
					newNamespace := copyNamespaceElements(ns)
					newNamespace[secondlastElement].Value = k
					metrics = append(metrics, createNamespace(mt, v, newNamespace, meta))
				}
			} else {
				cpuTime, err := wrapper.GetCPUStatistics(id)
				if err != nil {
					return metrics, err
				}
				if !metricStored(metrics, ns) {
					metrics = append(metrics, createNamespace(mt, cpuTime, ns, meta))
				}

			}

		}
	}

	return metrics, nil

}

// GetConfigPolicy returns a config policy
func (LibvirtCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	policy.AddNewStringRule([]string{"intel", "libvirt"}, "uri", false, plugin.SetDefaultString(defaultURI))
	policy.AddNewBoolRule([]string{"intel", "libvirt"}, "nova", false, plugin.SetDefaultBool(false))

	return *policy, nil
}

// GetMetricTypes returns metric types that can be collected
func (LibvirtCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {

	var metrics []plugin.Metric

	uri := getLibvirtURI(cfg)

	conn, err := libvirt.NewVirConnection(uri)
	defer conn.CloseConnection()

	ids, err := wrapper.GetInstances(conn)
	if err != nil {
		return metrics, err
	}

	for _, domain := range ids {

		domainName, err := domain.GetName()
		if err != nil {
			return metrics, err
		}

		ns := plugin.NewNamespace(Vendor, Plugin, domainName)
		for _, value := range nsTypes.cpu {
			metrics = append(metrics, createMetric(ns.AddStaticElements("cpu", value)))
			metrics = append(metrics, createMetric(ns.AddStaticElement("cpu").AddDynamicElement("cpu_id", "id of vcpu").AddStaticElement(value)))

		}
		for _, value := range nsTypes.disk {
			//ignoring errors, domain don't need to have disk attached
			disks, _ := wrapper.GetDomainDisks(domain)
			for _, disk := range disks {
				metrics = append(metrics, createMetric(
					ns.AddStaticElements("disk", disk, value)))
			}
		}
		for _, value := range nsTypes.network {
			//ignoring errors, domain don't need to have network interface
			interfaces, _ := wrapper.GetDomainInterfaces(domain)
			for _, netInterface := range interfaces {
				metrics = append(metrics, createMetric(
					ns.AddStaticElements("network", netInterface, value)))
			}
		}
		for _, value := range nsTypes.memory {
			metrics = append(metrics, createMetric(ns.AddStaticElements("memory", value)))
		}

	}

	ns := plugin.NewNamespace(Vendor, Plugin).
		AddDynamicElement("domain_id", "an id of libvirt domain")

	for _, value := range nsTypes.cpu {
		metrics = append(metrics, createMetric(ns.AddStaticElements("cpu", value)))
		metrics = append(metrics, createMetric(ns.AddStaticElement("cpu").AddDynamicElement("cpu_id", "id of vcpu").AddStaticElement(value)))

	}
	for _, value := range nsTypes.disk {
		metrics = append(metrics, createMetric(
			ns.AddStaticElements("disk").AddDynamicElement("device_name", "a name of filesystem device").AddStaticElement(value)))
	}
	for _, value := range nsTypes.network {
		metrics = append(metrics, createMetric(
			ns.AddStaticElements("network").AddDynamicElement("network_interface", "a name of network interface").AddStaticElement(value)))
	}
	for _, value := range nsTypes.memory {
		metrics = append(metrics, createMetric(ns.AddStaticElements("memory").AddStaticElement(value)))
	}
	return metrics, nil
}
