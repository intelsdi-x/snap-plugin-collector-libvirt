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

var disk_metrics_types = []string{"wrreq", "rdreq", "wrbytes", "rdbytes"}

func diskStat(ns []string, dom libvirt.VirDomain) (*plugin.PluginMetricType, error) {
	switch {
	case regexp.MustCompile(`^/libvirt/.*/.*/disk/.*/wrreq`).MatchString(joinNamespace(ns)):
		disk := ns[4]
		disk_stat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatInt(disk_stat.WrReq, 10),
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/disk/.*/rdreq`).MatchString(joinNamespace(ns)):
		disk := ns[4]
		disk_stat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatInt(disk_stat.RdReq, 10),
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/disk/.*/wrbytes`).MatchString(joinNamespace(ns)):
		disk := ns[4]
		disk_stat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatInt(disk_stat.WrBytes, 10),
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/disk/.*/rdbytes`).MatchString(joinNamespace(ns)):
		disk := ns[4]
		disk_stat, err := dom.BlockStats(disk)
		if err != nil {
			return nil, err
		}
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      strconv.FormatInt(disk_stat.RdBytes, 10),
			Timestamp_: time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("Unknown error processing %v", ns)
}

func listDisks(domXml *etree.Document) []string {
	disks := []string{}
	for _, t := range domXml.FindElements("//domain/devices/disk[@device='disk']/target") {
		for _, i := range t.Attr {
			if i.Key == "dev" {
				disks = append(disks, i.Value)
			}
		}
	}
	return disks
}

func getDiskMetricTypes(dom libvirt.VirDomain, hostname string) ([]plugin.PluginMetricType, error) {
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

	for _, metric := range disk_metrics_types {

		for _, disk := range listDisks(domXml) {
			mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"libvirt", hostname, domainname, "disk", disk, metric}})

		}
	}
	return mts, nil
}
