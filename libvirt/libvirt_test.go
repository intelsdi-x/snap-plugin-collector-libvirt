//
// +build small

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
	"regexp"
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/cdata"
	"github.com/intelsdi-x/snap/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibirtPlugin(t *testing.T) {
	Convey("Meta should return metadata for the plugin", t, func() {
		meta := Meta()
		So(meta.Name, ShouldResemble, Name)
		So(meta.Version, ShouldResemble, Version)
		So(meta.Type, ShouldResemble, plugin.CollectorPluginType)
	})

	Convey("Create Libvirt Collector", t, func() {
		libvirtCol := NewLibvirtCollector()
		Convey("So psCol should not be nil", func() {
			So(libvirtCol, ShouldNotBeNil)
		})
		Convey("So psCol should be of Libvirt type", func() {
			So(libvirtCol, ShouldHaveSameTypeAs, &Libvirt{})
		})
		Convey("libvirtCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := libvirtCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a cpolicy.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
			})
		})
	})
	Convey("Join namespace ", t, func() {
		namespace1 := []string{"intel", "libvirt", "one"}
		namespace2 := []string{}
		Convey("So namespace should equal intel/libvirt/one", func() {
			So("/intel/libvirt/one", ShouldResemble, joinNamespace(namespace1))
		})
		Convey("So namespace should equal slash", func() {
			So("/", ShouldResemble, joinNamespace(namespace2))
		})

	})
	Convey("Get Domain from Namespace ", t, func() {
		namespace1 := []string{"libvirt", "t1000", "mem", "mem"}
		namespace2 := []string{"libvirt"}
		Convey("So should return t1000", func() {
			domain, err := namespacetoDomain(namespace1)
			So("t1000", ShouldResemble, domain)
			So(err, ShouldBeNil)
		})
		Convey("So should thrown an error", func() {
			domain, err := namespacetoDomain(namespace2)
			So("", ShouldResemble, domain)
			So(err, ShouldNotBeNil)
		})

	})
	Convey("Get Hyervisor URI ", t, func() {
		hypervisorURI := make(map[string]ctypes.ConfigValue)
		Convey("So should return tcp+ssh://test", func() {
			hypervisorURI["uri"] = ctypes.ConfigValueStr{Value: "tcp+ssh://test"}
			uri := getHypervisorURI(hypervisorURI)
			So("tcp+ssh://test", ShouldResemble, uri)
		})
		Convey("So should return empty string", func() {
			uri := getHypervisorURI(hypervisorURI)
			So(defaultURI, ShouldResemble, uri)
		})

	})
	Convey("Get Metrics ", t, func() {
		libvirtCol := NewLibvirtCollector()
		cfgNode := cdata.NewNode()
		cfgNode.AddItem("uri", ctypes.ConfigValueStr{Value: "test:///default"})
		var cfg = plugin.ConfigType{
			ConfigDataNode: cfgNode,
		}
		Convey("So should return 26 types of metrics", func() {
			metrics, err := libvirtCol.GetMetricTypes(cfg)
			So(26, ShouldResemble, len(metrics))
			So(err, ShouldBeNil)
		})
		Convey("So should check namespace", func() {
			metrics, err := libvirtCol.GetMetricTypes(cfg)
			vcpuNamespace := metrics[0].Namespace().String()
			vcpu := regexp.MustCompile(`^/libvirt/test/cpu/vcpu/0/cputime`)
			So(true, ShouldEqual, vcpu.MatchString(vcpuNamespace))
			So(err, ShouldBeNil)

			vcpuNamespace1 := metrics[1].Namespace().String()
			vcpu1 := regexp.MustCompile(`^/libvirt/test/cpu/vcpu/1/cputime`)
			So(true, ShouldEqual, vcpu1.MatchString(vcpuNamespace1))
			So(err, ShouldBeNil)

			cpuNamespace := metrics[2].Namespace().String()
			cpu := regexp.MustCompile(`^/libvirt/test/cpu/cputime`)
			So(true, ShouldEqual, cpu.MatchString(cpuNamespace))
			So(err, ShouldBeNil)

			memNamespace := metrics[3].Namespace().String()
			mem := regexp.MustCompile(`^/libvirt/test/mem/mem`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[4].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/max`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[5].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/swap_in`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[6].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/swap_out`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[7].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/major_fault`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[8].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/min_fault`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[9].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/unused`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[10].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/available`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[11].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/actual_balloon`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[12].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/rss`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)

			memNamespace = metrics[13].Namespace().String()
			mem = regexp.MustCompile(`^/libvirt/test/mem/nr`)
			So(true, ShouldEqual, mem.MatchString(memNamespace))
			So(err, ShouldBeNil)
		})

	})
	Convey("Collect Metrics", t, func() {
		libvirtCol := &Libvirt{}
		cfgNode := cdata.NewNode()
		cfgNode.AddItem("uri", ctypes.ConfigValueStr{Value: "test:///default"})
		cfgNode.AddItem("nova", ctypes.ConfigValueBool{Value: false})

		Convey("So should get cpu metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("libvirt", "test", "cpu", "vcpu", "0", "cputime"),
				Config_:    cfgNode,
			}}
			collect, err := libvirtCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should get vcpu metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("libvirt", "test", "cpu", "cputime"),
				Config_:    cfgNode,
			}}
			collect, err := libvirtCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType uint64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should return only one metric", func() {
			metrics := []plugin.MetricType{
				{
					Namespace_: core.NewNamespace("libvirt", "test", "cpu", "cputime"),
					Config_:    cfgNode,
				},
				{
					Namespace_: core.NewNamespace("libvirt", "*", "cpu", "cputime"),
					Config_:    cfgNode,
				},
			}
			collect, err := libvirtCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType uint64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})

	})
}
