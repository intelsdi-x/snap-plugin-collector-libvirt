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

package libvirtcollector

import (
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	libvirt "github.com/sandlbn/libvirt-go"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibirtPlugin(t *testing.T) {

	Convey("Create Libvirt Collector", t, func() {
		libvirtCol := LibvirtCollector{}
		Convey("So psCol should not be nil", func() {
			So(libvirtCol, ShouldNotBeNil)
		})
		Convey("So psCol should be of Libvirt type", func() {
			So(libvirtCol, ShouldHaveSameTypeAs, LibvirtCollector{})
		})
		Convey("libvirtCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := libvirtCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a plugin.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})
	Convey("Get Metric Types", t, func() {
		libvirtCol := LibvirtCollector{}
		cfg := plugin.Config{
			"uri":  "test:///default",
			"nova": false,
		}
		metrics, err := libvirtCol.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 38)
	})
	Convey("Collect Metrics", t, func() {
		libvirtCol := LibvirtCollector{}

		config := plugin.Config{
			"uri":  "test:///default",
			"nova": false,
		}
		mts := []plugin.Metric{}

		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElements("cpu", "cputime"), Config: config})

		metrics, err := libvirtCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 1)
		So(metrics[0].Data, ShouldNotBeNil)
		So(metrics[0].Namespace.Strings()[2], ShouldResemble, "test")

	})
	Convey("Collect Vcpu Metrics", t, func() {
		libvirtCol := LibvirtCollector{}

		config := plugin.Config{
			"uri":  "test:///default",
			"nova": false,
		}
		mts := []plugin.Metric{}

		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElement("cpu").AddDynamicElement("cpu_id", "id of vcpu").AddStaticElement("cputime"), Config: config})
		metrics, err := libvirtCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 2)
		So(metrics[0].Data, ShouldNotBeNil)
		secondLastElement := len(metrics[0].Namespace.Strings()) - 2
		So(metrics[0].Namespace.Strings()[secondLastElement], ShouldBeIn, []string{"0", "1"})
		So(metrics[0].Namespace.Strings()[2], ShouldResemble, "test")
		So(metrics[1].Namespace.Strings()[secondLastElement], ShouldBeIn, []string{"0", "1"})
		So(metrics[1].Namespace.Strings()[2], ShouldResemble, "test")

	})
	Convey("Check if metric is not collected twice", t, func() {
		libvirtCol := LibvirtCollector{}

		config := plugin.Config{
			"uri":  "test:///default",
			"nova": false,
		}
		mts := []plugin.Metric{}

		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElement("cpu").AddDynamicElement("cpu_id", "id of vcpu").AddStaticElement("cputime"), Config: config})
		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElement("cpu").AddDynamicElement("cpu_id", "id of vcpu").AddStaticElement("cputime"), Config: config})
		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElements("cpu", "cputime"), Config: config})
		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElements("cpu", "cputime"), Config: config})
		metrics, err := libvirtCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 3)

	})

	Convey("Merge two map[string]string", t, func() {
		one := make(map[string]string)
		two := make(map[string]string)

		one["test"] = "test"
		two["test1"] = "test"

		three := merge(one, two)

		result := make(map[string]string)
		result["test"] = "test"
		result["test1"] = "test"

		So(result, ShouldResemble, three)

	})
	Convey("copy namespace", t, func() {

		mt := plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElements("cpu", "cputime")}
		ns := copyNamespace(mt)
		result := []plugin.NamespaceElement{plugin.NamespaceElement{Value: "intel", Description: "", Name: ""}, plugin.NamespaceElement{Value: "libvirt", Description: "", Name: ""}, plugin.NamespaceElement{Value: "*", Description: "an id of libvirt domain", Name: "domain_id"}, plugin.NamespaceElement{Value: "cpu", Description: "", Name: ""}, plugin.NamespaceElement{Value: "cputime", Description: "", Name: ""}}

		So(ns, ShouldResemble, result)

	})

	Convey("filter Namespace", t, func() {
		mts := []plugin.Metric{
			plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElements("cpu", "cputime")},
		}

		countResult, result := filterNamespace("cpu", mts)
		So(mts, ShouldResemble, result)
		So(countResult, ShouldResemble, 1)

		countResult, result2 := filterNamespace("memory", mts)
		expectedType := []plugin.Metric{}
		So(result2, ShouldHaveSameTypeAs, expectedType)
		So(countResult, ShouldResemble, 0)

	})
	Convey("metric Stored", t, func() {
		mts := []plugin.Metric{
			plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin, "test", "cpu", "cputime")},
		}

		ns := plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElements("cpu", "cputime")}
		ns.Namespace[2].Value = "test"

		stored := metricStored(mts, ns.Namespace)
		So(stored, ShouldResemble, true)

	})
	Convey("GetInstances", t, func() {
		mts := []plugin.Metric{}
		conn, err := libvirt.NewVirConnection("test:///default")
		config := plugin.Config{
			"uri":  "test:///default",
			"nova": false,
		}
		So(err, ShouldBeNil)

		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddDynamicElement("domain_id", "an id of libvirt domain").AddStaticElement("cpu").AddDynamicElement("cpu_id", "id of vcpu").AddStaticElement("cputime"), Config: config})
		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddStaticElement("test").AddStaticElement("cpu").AddDynamicElement("cpu_id", "id of vcpu").AddStaticElement("cputime"), Config: config})
		result, err := getInstances(conn, mts)
		So(len(result), ShouldResemble, 1)
		So(err, ShouldBeNil)

	})
	Convey("Contains", t, func() {
		source := []string{"one", "two"}

		So(true, ShouldResemble, contains(source, "one"))
		So(false, ShouldResemble, contains(source, "three"))

	})

}
