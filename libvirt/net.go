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
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/sandlbn/libvirt-go"
)

var netMetricsTypes = []string{"rxbytes", "rxpackets", "rxerrs", "rxdrop",
	"txbytes", "txpackets", "txerrs", "txdrop"}

func interfaceStat(ns []string, dom libvirt.VirDomain) (*plugin.MetricType, error) {
	switch {
	case regexp.MustCompile(`^/libvirt/.*/net/.*/rxbytes`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.RxBytes,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/net/.*/rxpackets`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.RxPackets,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/net/.*/rxerrs`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.RxErrs,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/net/.*/rxdrop`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.RxDrop,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/net/.*/txbytes`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.TxBytes,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/net/.*/txpackets`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.TxPackets,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/net/.*/txerrs`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			fmt.Println(err)
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.TxErrs,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/net/.*/txdrop`).MatchString(joinNamespace(ns)):
		iface := ns[3]
		ifaceStat, err := dom.InterfaceStats(iface)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      ifaceStat.RxDrop,
			Timestamp_: time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func getNetMetricTypes(dom libvirt.VirDomain) ([]plugin.MetricType, error) {
	var mts []plugin.MetricType
	domXML, err := getDomainXML(dom)
	if err != nil {
		return nil, err
	}
	lXML := libvirtXML{domain: domXML}
	domainname, err := dom.GetName()
	if err != nil {
		return nil, err
	}

	for _, metric := range netMetricsTypes {

		for _, netInterface := range lXML.GetNics() {
			mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("libvirt", domainname, "net", netInterface, metric)})

		}
	}
	return mts, nil
}
