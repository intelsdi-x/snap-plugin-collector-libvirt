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
	"io/ioutil"
	"testing"

	"github.com/beevik/etree"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLibirtPluginNet(t *testing.T) {

	Convey("List interfaces", t, func() {
		buf, err := ioutil.ReadFile("./test_domain.xml")
		if err != nil {
			panic(err)
		}
		domXmlStr := string(buf)
		domXml := etree.NewDocument()
		domXml.ReadFromString(domXmlStr)
		data := listInterfaces(domXml)
		So(data, ShouldResemble, []string{"tap88709cbd-90"})

	})
	Convey("List interfaces with no net interface", t, func() {
		buf, err := ioutil.ReadFile("./test_domain_2.xml")
		if err != nil {
			panic(err)
		}
		domXmlStr := string(buf)
		domXml := etree.NewDocument()
		domXml.ReadFromString(domXmlStr)
		data := listInterfaces(domXml)
		So(data, ShouldResemble, []string{})

	})
}
