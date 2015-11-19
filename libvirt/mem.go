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

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/sandlbn/libvirt-go"
)

var memory_metrics_types = []string{"mem", "max", "swap_in", "swap_out", "major_fault",
	"min_fault", "unused", "available", "actual_balloon", "rss", "nr"}

func memStat(ns []string, dom libvirt.VirDomain) (*plugin.PluginMetricType, error) {
	info, err := dom.GetInfo()
	if err != nil {
		return nil, err
	}

	switch {
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/mem`).MatchString(joinNamespace(ns)):
		memory := strconv.FormatUint(info.GetMemory(), 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/max`).MatchString(joinNamespace(ns)):
		memory := strconv.FormatUint(info.GetMaxMem(), 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/swap_in`).MatchString(joinNamespace(ns)):
		swap_in, err := getMemoryInfo("swap_in", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(swap_in, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/swap_out`).MatchString(joinNamespace(ns)):
		swap_out, err := getMemoryInfo("swap_out", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(swap_out, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/min_fault`).MatchString(joinNamespace(ns)):
		min_fault, err := getMemoryInfo("min_fault", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(min_fault, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/major_fault`).MatchString(joinNamespace(ns)):
		major_fault, err := getMemoryInfo("major_fault", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(major_fault, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/unused`).MatchString(joinNamespace(ns)):
		unused, err := getMemoryInfo("unused", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(unused, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/available`).MatchString(joinNamespace(ns)):
		available, err := getMemoryInfo("available", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(available, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/actual_balloon`).MatchString(joinNamespace(ns)):
		actual_balloon, err := getMemoryInfo("actual_balloon", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(actual_balloon, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/rss`).MatchString(joinNamespace(ns)):
		rss, err := getMemoryInfo("rss", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(rss, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	case regexp.MustCompile(`^/libvirt/.*/.*/mem/nr`).MatchString(joinNamespace(ns)):
		nr, err := getMemoryInfo("nr", dom)
		if err != nil {
			fmt.Println(err)
		}
		memory := strconv.FormatUint(nr, 10)
		return &plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      memory,
			Timestamp_: time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("Unknown error processing %v", ns)

}

func getMemoryInfo(tag string, dom libvirt.VirDomain) (uint64, error) {
	meminfo, err := dom.MemoryStats(5, 0)
	if err != nil {
		return 0, err
	}
	switch tag {
	case "swap_in":
		return parseMemStats(meminfo, 0), nil
	case "swap_out":
		return parseMemStats(meminfo, 1), nil
	case "major_fault":
		return parseMemStats(meminfo, 2), nil
	case "min_fault":
		return parseMemStats(meminfo, 3), nil
	case "unused":
		return parseMemStats(meminfo, 4), nil
	case "available":
		return parseMemStats(meminfo, 5), nil
	case "actual_balloon":
		return parseMemStats(meminfo, 6), nil
	case "rss":
		return parseMemStats(meminfo, 7), nil
	case "nr":
		return parseMemStats(meminfo, 8), nil
	}
	return 0, nil
}

func parseMemStats(memstat []libvirt.VirDomainMemoryStat, nr int32) uint64 {

	var metric uint64 = 0
	for i := 0; i < len(memstat); i++ {
		fmt.Println(i, memstat[i].Tag, memstat[i].Val)
		if memstat[i].Tag == nr {
			return memstat[i].Val
		}
	}
	return metric

}
func getMemMetricTypes(dom libvirt.VirDomain, hostname string) ([]plugin.PluginMetricType, error) {
	mts := make([]plugin.PluginMetricType, 0)

	domainname, err := dom.GetName()
	if err != nil {
		return nil, err
	}

	for _, metric := range memory_metrics_types {

		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"libvirt", hostname, domainname, "mem", metric}})
	}
	return mts, nil
}
