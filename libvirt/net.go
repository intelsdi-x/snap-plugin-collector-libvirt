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
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/beevik/etree"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/sandlbn/libvirt-go"
)

var net_metrics_types = []string{"rxbytes", "rxpackets", "rxerrs", "rxdrop",
	"txbytes", "txpackets", "txerrs", "txdrop"}

func interfaceStat(ns []string, dom libvirt.VirDomain) (*plugin.PluginMetricType, error) {
	switch {
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/rxbytes`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatInt(iface_stat.RxBytes, 10),
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/rxpackets`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      iface_stat.RxPackets,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/rxerrs`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      iface_stat.RxErrs,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/rxdrop`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      iface_stat.RxDrop,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/txbytes`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatInt(iface_stat.TxBytes, 10),
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/txpackets`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      iface_stat.TxPackets,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/txerrs`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      iface_stat.TxErrs,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/net/.*/txdrop`).MatchString(joinNamespace(ns)):
		iface := ns[4]
		iface_stat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      iface_stat.RxDrop,
			Timestamp_: time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func listInterfaces(domXml *etree.Document) []string {
	network_interfaces := []string{}
	for _, t := range domXml.FindElements("//domain/devices/interface/target") {
		for _, i := range t.Attr {
			network_interfaces = append(network_interfaces, i.Value)
		}

	}
	return network_interfaces
}

func getNetMetricTypes(dom libvirt.VirDomain, hostname string) ([]plugin.PluginMetricType, error) {
	mts := make([]plugin.PluginMetricType, 0)
	domXmlDesc, err := dom.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	domXml := etree.NewDocument()
	domXml.ReadFromString(domXmlDesc)

	domainname, err := dom.GetName()
	if err != nil {
		return nil, err
	}

	for _, metric := range net_metrics_types {

		for _, net_interface := range listInterfaces(domXml) {
			mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"libvirt", hostname, domainname, "net", net_interface, metric}})

		}
	}
	return mts, nil
}
