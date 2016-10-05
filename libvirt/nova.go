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
	"fmt"

	"github.com/beevik/etree"
	"github.com/sandlbn/libvirt-go"
)

type libvirtXML struct {
	domain *etree.Document
}

// GetUUID returns uuid of the Nova Instance
func (l libvirtXML) GetUUID() (string, error) {
	entries := l.domain.FindElements("//domain/sysinfo/system/entry[@name='uuid']")
	if len(entries) == 1 {
		return entries[0].Text(), nil
	}
	return "", fmt.Errorf("Can't find nova uuid")
}

// GetDisks returns list of the disk devices
func (l libvirtXML) GetDisks() []string {
	disks := []string{}
	for _, t := range l.domain.FindElements("//domain/devices/disk[@device='disk']/target") {
		for _, i := range t.Attr {
			if i.Key == "dev" {
				disks = append(disks, i.Value)
			}
		}
	}
	return disks
}

// GetNovaName returns name of the Nova Instance
func (l libvirtXML) GetNovaName() string {
	entry := l.domain.FindElement(".//nova:name")
	return entry.Text()
}

// GetNovaName returns memory size of the Nova instance flavor
func (l libvirtXML) GetNovaFlavorMemory() string {
	entry := l.domain.FindElement(".//nova:memory")
	return entry.Text()
}

// GetNovaFlavorDisk returns disk size of the Nova instance flavor
func (l libvirtXML) GetNovaFlavorDisk() string {
	entry := l.domain.FindElement(".//nova:disk")
	return entry.Text()
}

// GetNovaFlavorSwap returns swap size of the Nova instance flavor
func (l libvirtXML) GetNovaFlavorSwap() string {
	entry := l.domain.FindElement(".//nova:swap")
	return entry.Text()
}

// GetNovaFlavorVcpus returns vcpus count of the Nova instance flavor
func (l libvirtXML) GetNovaFlavorVcpus() string {
	entry := l.domain.FindElement(".//nova:vcpus")
	return entry.Text()
}

// GetNovaFlavorEmphemeral returns emphemeral size of the Nova instance flavor
func (l libvirtXML) GetNovaFlavorEmphemeral() string {
	entry := l.domain.FindElement(".//nova:ephemeral")
	return entry.Text()
}

// GetNovaFlavorEmphemeral returns Nova instance flavor
func (l libvirtXML) GetNovaFlavor() map[string]string {
	tags := map[string]string{}
	tags["nova_memory"] = l.GetNovaFlavorMemory()
	tags["nova_disk"] = l.GetNovaFlavorDisk()
	tags["nova_swap"] = l.GetNovaFlavorSwap()
	tags["nova_emphemeral"] = l.GetNovaFlavorEmphemeral()
	tags["nova_vcpus"] = l.GetNovaFlavorVcpus()
	return tags

}

// GetDisks returns list of the nics
func (l libvirtXML) GetNics() []string {
	networkInterfaces := []string{}
	for _, t := range l.domain.FindElements("//domain/devices/interface/target") {
		for _, i := range t.Attr {
			networkInterfaces = append(networkInterfaces, i.Value)
		}

	}
	return networkInterfaces
}
func getDomainXML(dom libvirt.VirDomain) (*etree.Document, error) {
	domXMLDesc, err := dom.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	domXML := etree.NewDocument()
	domXML.ReadFromString(domXMLDesc)
	return domXML, nil
}
