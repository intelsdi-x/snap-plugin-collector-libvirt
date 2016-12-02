/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2016 Intel Corporation

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
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/sandlbn/libvirt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLibirt(t *testing.T) {
	Convey("Get Interface Name when Interface don't exist", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 1)
		interf, err := GetDomainInterfaces(domains[0])
		So(err, ShouldNotBeNil)
		So(len(interf), ShouldResemble, 0)
	})
	Convey("Get Interface Name when Interface exist", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		var lXML Domain
		domXMLStr := string(buf)
		xml.Unmarshal([]byte(domXMLStr), &lXML)
		So(lXML.Devices.Interface[0].Target.AttrDev, ShouldResemble, "tap88709cbd-90")
	})
	Convey("Get Interface Name when Multiple Interfaces exist", t, func() {
		buf, err := ioutil.ReadFile("./test_domain_multiple_interfaces.xml")
		if err != nil {
			panic(err)
		}
		var lXML Domain
		domXMLStr := string(buf)
		xml.Unmarshal([]byte(domXMLStr), &lXML)
		So(len(lXML.Devices.Interface), ShouldResemble, 3)
	})
	Convey("Get Disk Name when disk don't exist", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 1)
		disk, err := GetDomainDisks(domains[0])
		So(err, ShouldBeNil)
		So(len(disk), ShouldResemble, 0)
	})
	Convey("Get Cpu Statistics", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 1)
		cpu, err := GetCPUStatistics(domains[0])
		So(err, ShouldBeNil)
		So(cpu, ShouldBeGreaterThanOrEqualTo, 0)
	})
	Convey("Get vCpu Statistics", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 1)
		cpu, err := GetVCPUStatistics(domains[0])
		So(err, ShouldBeNil)
		expectedType := make(map[string]int64)
		So(cpu, ShouldHaveSameTypeAs, expectedType)
	})
	Convey("Get Memory Statistics on Test Driver", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 1)
		cpu, err := GetMemoryStatistics(domains[0], "swap_in", "free")
		So(err.Error(), ShouldContainSubstring, "this function is not supported by the connection driver: virDomainMemoryStats")
		emptyMap := make(map[string]int64)
		So(cpu, ShouldResemble, emptyMap)
	})
	Convey("Get Nova Metadata", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		var lXML Domain
		domXMLStr := string(buf)
		xml.Unmarshal([]byte(domXMLStr), &lXML)

		So(lXML.GetNovaDisk(), ShouldResemble, "20")
		So(lXML.GetNovaMemory(), ShouldResemble, "2048")
		So(lXML.GetNovaUUID(), ShouldResemble, "5a26891c-efb0-4c6a-8bef-b54c45296136")
		So(lXML.GetNovaFlavor(), ShouldResemble, "m1.small")
		So(lXML.GetNovaEphemeral(), ShouldResemble, "0")
		So(lXML.GetNovaName(), ShouldResemble, "demo_gui")
		So(lXML.GetNovaOwner(), ShouldResemble, "3aff4d1bd3b74afb8262411eaae47d7c")
		So(lXML.GetNovaPackage(), ShouldResemble, "2015.1.1")
		So(lXML.GetNovaSwap(), ShouldResemble, "0")
		So(lXML.GetNovaVcpus(), ShouldResemble, "1")

	})
	Convey("Get Nova Metadata when Metadata entity dosn't exist ", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.1.xml")
		if err != nil {
			panic(err)
		}
		var lXML Domain
		domXMLStr := string(buf)
		xml.Unmarshal([]byte(domXMLStr), &lXML)
		data := generateNovaMetadata(&lXML)
		expectedData := map[string]string{"nova_uuid": "5a26891c-efb0-4c6a-8bef-b54c45296136"}
		So(data, ShouldResemble, expectedData)
	})
}
