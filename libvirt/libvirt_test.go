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
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/sandlbn/libvirt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLibvirtFundaments(t *testing.T) {
	Convey("Connect to default libvirt", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		defer conn.CloseConnection()
		Convey("connections should success", func() {
			So(err, ShouldBeNil)
		})

		domains, err := GetInstances(conn)
		Convey("domain should be available", func() {
			So(err, ShouldBeNil)
			Convey("and be the only one", func() {
				So(len(domains), ShouldResemble, 1)
			})
		})
		interf, err := GetDomainInterfaces(domains[0])
		Convey("no interface", func() {
			Convey("should be error", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("number of interfaces should be zero", func() {
				So(len(interf), ShouldResemble, 0)
			})

		})
		name, err := domains[0].GetName()
		Convey("Its name should be test", func() {
			So(name, ShouldEqual, "test")
			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
		id, err := domains[0].GetID()
		Convey("and its ID should be 1", func() {
			So(id, ShouldEqual, 1)
			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
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
	Convey("Get Cpu Statistics", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		defer conn.CloseConnection()
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
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 1)
		cpu, err := GetVCPUStatistics(domains[0])
		So(err, ShouldBeNil)
		expectedType := make(map[string]int64)
		So(cpu, ShouldHaveSameTypeAs, expectedType)
	})
	Convey("Get Network Statistics without network interface neither path", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		network, err := GetNetworkStatistics(domains[0])
		So(err, ShouldNotBeNil)
		emptyMap := make(map[string]libvirt.VirDomainInterfaceStats)
		So(network, ShouldResemble, emptyMap)
	})
	Convey("Get Network Statistics with one path", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		So(err, ShouldBeNil)
		_, err = conn.DomainCreateXMLFromFile("./test_domain_multiple_interfaces.xml", 0)
		So(err, ShouldBeNil)
		defer conn.CloseConnection()
		name := "instance-00000031"
		domain, err := GetInstanceByDomainName(conn, name)
		So(err, ShouldBeNil)
		network, err := GetNetworkStatistics(domain, "tap88709cbd-90")
		So(network["tap88709cbd-90"].RxBytes, ShouldNotBeNil)
		So(network["tap88709cbd-90"].RxPackets, ShouldNotBeNil)
		So(network["tap88709cbd-90"].RxDrop, ShouldNotBeNil)
		So(network["tap88709cbd-90"].RxErrs, ShouldNotBeNil)
		So(network["tap88709cbd-90"].TxBytes, ShouldNotBeNil)
		So(network["tap88709cbd-90"].TxDrop, ShouldNotBeNil)
		So(network["tap88709cbd-90"].TxErrs, ShouldNotBeNil)
		So(network["tap88709cbd-90"].TxPackets, ShouldNotBeNil)
		So(err, ShouldBeNil)
		emptyMap := make(map[string]libvirt.VirDomainInterfaceStats)
		So(network, ShouldHaveSameTypeAs, emptyMap)
		So(network, ShouldNotBeEmpty)
	})
	Convey("Get Network Statistics with error", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		defer conn.CloseConnection()
		conn.DomainCreateXMLFromFile("./test_domain_multiple_interfaces.xml", 0)
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		block, err := GetNetworkStatistics(domains[0], "wrong")
		emptyMap := make(map[string]libvirt.VirDomainInterfaceStats)
		So(block, ShouldResemble, emptyMap)
		So(err, ShouldNotBeNil)
	})
	Convey("Get Block Statistics", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		conn.DomainCreateXMLFromFile("test_domain.xml", 0)
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		name := "instance-00000031"
		domain, err := GetInstanceByDomainName(conn, name)
		block, err := GetBlockStatistics(domain)
		So(err, ShouldBeNil)
		emptyMap := make(map[string]libvirt.VirDomainBlockStats)
		So(block, ShouldHaveSameTypeAs, emptyMap)
	})
	Convey("Get Block Statistics with error", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		defer conn.CloseConnection()
		conn.DomainCreateXMLFromFile("./test_domain_multiple_interfaces.xml", 0)
		So(err, ShouldBeNil)
		name := "instance-00000031"
		domain, err := GetInstanceByDomainName(conn, name)
		So(err, ShouldBeNil)
		block, err := GetBlockStatistics(domain, "wrong")
		emptyMap := make(map[string]libvirt.VirDomainBlockStats)
		So(block, ShouldResemble, emptyMap)
		So(err, ShouldNotBeNil)
	})
	Convey("Get Memory Statistics on Test Driver", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		domains, err := GetInstances(conn)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 2)
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
	Convey("Get Nova Metadata when Metadata entity doesn't exist ", t, func() {
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
	Convey("Get Instances from a slice", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		instances := []string{"test"}
		So(err, ShouldBeNil)
		domains, err := GetRequestedInstances(conn, instances)
		So(err, ShouldBeNil)
		So(len(domains), ShouldResemble, 1)
	})
	Convey("Get Instances from a slice, when instance doesn't exist", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		instances := []string{"testi1"}
		So(err, ShouldBeNil)
		domains, err := GetRequestedInstances(conn, instances)
		So(err, ShouldNotBeNil)
		So(len(domains), ShouldResemble, 0)
	})
}

func TestLibvirt(t *testing.T) {
	Convey("Get Interface Name", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		conn.DomainCreateXMLFromFile("test_domain.xml", 0)
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		name := "instance-00000031"
		domain, err := GetInstanceByDomainName(conn, name)
		So(err, ShouldBeNil)
		interf, err := GetDomainInterfaces(domain)
		So(err, ShouldBeNil)
		So(len(interf), ShouldBeGreaterThan, 0)
	})
	Convey("Get Disk Name when disk don't exist", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		conn.DomainCreateXMLFromFile("test_domain.xml", 0)
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		name := "instance-00000031"
		domain, err := GetInstanceByDomainName(conn, name)
		disk, err := GetDomainDisks(domain)
		So(err, ShouldBeNil)
		So(len(disk), ShouldResemble, 1)
	})
	Convey("Get Nova Metadata info", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		conn.DomainCreateXMLFromFile("test_domain.xml", 0)
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		name := "instance-00000031"
		domain, err := GetInstanceByDomainName(conn, name)
		metadata, err := GetNovaMetadata(domain)
		So(err, ShouldBeNil)
		emptyMap := make(map[string]string)
		So(metadata, ShouldHaveSameTypeAs, emptyMap)
		So(len(metadata), ShouldResemble, 10)
		So(metadata["nova_memory_mb"], ShouldEqual, "2048")
		So(metadata["nova_owner"], ShouldEqual, "3aff4d1bd3b74afb8262411eaae47d7c")
		So(metadata["nova_uuid"], ShouldEqual, "5a26891c-efb0-4c6a-8bef-b54c45296136")
		So(metadata["nova_flavor"], ShouldEqual, "m1.small")
		So(metadata["nova_ephemeral_mb"], ShouldEqual, "0")
		So(metadata["nova_disk_mb"], ShouldEqual, "20")
		So(metadata["nova_swap_mb"], ShouldEqual, "0")
		So(metadata["nova_vcpus"], ShouldEqual, "1")
		So(metadata["nova_name"], ShouldEqual, "demo_gui")
		So(metadata["nova_package"], ShouldEqual, "2015.1.1")
	})
	Convey("ParseMemStats", t, func() {
		Array := make([]libvirt.VirDomainMemoryStat, 3)
		Array[0].Tag = 77
		Array[0].Val = 7
		Array[1].Tag = 33
		Array[1].Val = 3
		Array[2].Tag = 55
		Array[2].Val = 5

		var nr int32 = 33
		number := parseMemStats(Array, nr)
		var n int64 = 3
		So(number, ShouldEqual, n)
		So(number, ShouldHaveSameTypeAs, n)
	})
	Convey("ParseMemStats with empty array", t, func() {
		emptyArray := make([]libvirt.VirDomainMemoryStat, 0)
		var nr int32 = 33
		number := parseMemStats(emptyArray, nr)
		var n int64 = 0
		So(number, ShouldEqual, n)
		So(number, ShouldHaveSameTypeAs, n)
	})
	Convey("Get instances ID", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		defer conn.CloseConnection()
		So(err, ShouldBeNil)
		ids, err := GetInstanceIds(conn)
		fmt.Println(ids)
		fmt.Println(len(ids))
		So(err, ShouldBeNil)
		mockedIds := []string{"test", "instance-00000031"}
		emptyArray := make([]string, 0)
		So(ids, ShouldHaveSameTypeAs, emptyArray)
		So(ids[0], ShouldBeIn, mockedIds)
		So(ids[1], ShouldBeIn, mockedIds)
		So(ids[0], ShouldNotEqual, ids[1])
	})
	Convey("Get instances ID with error", t, func() {
		conn := libvirt.VirConnection{}
		ids, err := GetInstanceIds(conn)
		So(err, ShouldNotBeNil)
		So(ids, ShouldBeNil)
	})
}
