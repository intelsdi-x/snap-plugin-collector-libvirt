//
// +build small

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
	"io/ioutil"
	"testing"

	"github.com/beevik/etree"
	"github.com/sandlbn/libvirt-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLibirtPluginNova(t *testing.T) {
	Convey("Get Nova Name", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		domXMLStr := string(buf)
		domXML := etree.NewDocument()
		domXML.ReadFromString(domXMLStr)
		lXML := libvirtXML{domain: domXML}
		So(lXML.GetNovaName(), ShouldResemble, "demo_gui")

	})
	Convey("Get Nova Flavor", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		domXMLStr := string(buf)
		domXML := etree.NewDocument()
		domXML.ReadFromString(domXMLStr)
		lXML := libvirtXML{domain: domXML}
		expectedFlavor := map[string]string{"nova_disk": "20", "nova_emphemeral": "0", "nova_memory": "2048", "nova_swap": "0", "nova_vcpus": "1"}
		So(lXML.GetNovaFlavor(), ShouldResemble, expectedFlavor)

	})
	Convey("Get Nova UUID", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		domXMLStr := string(buf)
		domXML := etree.NewDocument()
		domXML.ReadFromString(domXMLStr)
		lXML := libvirtXML{domain: domXML}
		expectedUUID := "5a26891c-efb0-4c6a-8bef-b54c45296136"
		uuid, err := lXML.GetUUID()

		So(uuid, ShouldResemble, expectedUUID)
		So(err, ShouldBeNil)

	})
	Convey("Get Nova Nics", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		domXMLStr := string(buf)
		domXML := etree.NewDocument()
		domXML.ReadFromString(domXMLStr)
		lXML := libvirtXML{domain: domXML}
		expectedNICS := []string{"tap88709cbd-90"}
		nics := lXML.GetNics()
		So(nics, ShouldResemble, expectedNICS)

	})

	Convey("Get Nova Disk", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		domXMLStr := string(buf)
		domXML := etree.NewDocument()
		domXML.ReadFromString(domXMLStr)
		lXML := libvirtXML{domain: domXML}
		expectedDisks := []string{"vda"}
		disks := lXML.GetDisks()
		So(disks, ShouldResemble, expectedDisks)

	})

	Convey("Get Nova XML", t, func() {
		conn, err := libvirt.NewVirConnection("test:///default")
		So(err, ShouldBeNil)

		defer conn.CloseConnection()

		dom, err := conn.LookupDomainById(1)
		So(err, ShouldBeNil)
		domXML, err := getDomainXML(dom)

		samepleXML := etree.NewDocument()

		So(err, ShouldBeNil)
		So(domXML, ShouldNotBeNil)
		So(domXML, ShouldHaveSameTypeAs, samepleXML)

	})

}
