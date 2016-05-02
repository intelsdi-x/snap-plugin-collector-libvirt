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
	"time"

	"github.com/beevik/etree"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/sandlbn/libvirt-go"
)

var diskMetricsTypes = []string{"wrreq", "rdreq", "wrbytes", "rdbytes"}

func diskStat(ns []string, dom libvirt.VirDomain) (*plugin.MetricType, error) {
	switch {
	case regexp.MustCompile(`^/libvirt/.*/disk/.*/wrreq`).MatchString(joinNamespace(ns)):
		disk := ns[3]
		diskStat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      diskStat.WrReq,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/disk/.*/rdreq`).MatchString(joinNamespace(ns)):
		disk := ns[3]
		diskStat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      diskStat.RdReq,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/disk/.*/wrbytes`).MatchString(joinNamespace(ns)):
		disk := ns[3]
		diskStat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      diskStat.WrBytes,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/disk/.*/rdbytes`).MatchString(joinNamespace(ns)):
		disk := ns[3]
		diskStat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      diskStat.RdBytes,
			Timestamp_: time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func listDisks(domXML *etree.Document) []string {
	disks := []string{}
	for _, t := range domXML.FindElements("//domain/devices/disk[@device='disk']/target") {
		for _, i := range t.Attr {
			if i.Key == "dev" {
				disks = append(disks, i.Value)
			}
		}
	}
	return disks
}

func getDiskMetricTypes(dom libvirt.VirDomain) ([]plugin.MetricType, error) {
	var mts []plugin.MetricType
	domXMLDesc, err := dom.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	domXML := etree.NewDocument()
	domXML.ReadFromString(domXMLDesc)

	domainname, err := dom.GetName()
	if err != nil {
		return nil, err
	}

	for _, metric := range diskMetricsTypes {

		for _, disk := range listDisks(domXML) {
			mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("libvirt", domainname, "disk", disk, metric)})

		}
	}
	return mts, nil
}
