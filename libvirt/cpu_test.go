//
// +build unit

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
	"testing"

	"github.com/sandlbn/libvirt-go"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibirtPluginCpu(t *testing.T) {
	Convey("Collect cpu Metrics", t, func() {
		conn, _ := libvirt.NewVirConnection("test:///default")
		hostname, _ := conn.GetHostname()
		namespace := []string{"libvirt", hostname, "test", "cpu", "cputime"}
		namespace2 := []string{"libvirt", hostname, "test", "cpu", "vcpu", "0", "cputime"}
		namespace3 := []string{"libvirt", hostname, "test", "cpu", "0", "cputime", "bad"}
		dom, err := conn.LookupDomainByName("test")
		metric, err := cpuTimes(namespace, dom)
		So(err, ShouldBeNil)
		So(metric.Data_, ShouldNotBeNil)
		metric, err = cpuTimes(namespace2, dom)
		So(err, ShouldBeNil)
		So(metric.Data_, ShouldNotBeNil)
		metric, err = cpuTimes(namespace3, dom)
		So(metric, ShouldBeNil)
		So(err.Error(), ShouldStartWith, "Unknown error processing")

	})
	Convey("Get cpu metrics types", t, func() {
		conn, _ := libvirt.NewVirConnection("test:///default")
		hostname := "test"
		Convey("Join namespace ", func() {
			dom, _ := conn.LookupDomainByName(hostname)
			So(dom, ShouldNotBeNil)
			cpu_types, err := getCpuMetricTypes(dom, "test")
			So(err, ShouldBeNil)
			So(3, ShouldResemble, len(cpu_types))
		})

	})
}
