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

import "encoding/xml"

// Acpi type Libvirt
type Acpi struct {
}

// Address type Libvirt
type Address struct {
	AttrBus      string `xml:" bus,attr"  json:",omitempty"`
	AttrDomain   string `xml:" domain,attr"  json:",omitempty"`
	AttrFunction string `xml:" function,attr"  json:",omitempty"`
	AttrSlot     string `xml:" slot,attr"  json:",omitempty"`
	AttrType     string `xml:" type,attr"  json:",omitempty"`
}

// Alias type Libvirt
type Alias struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
}

// Apic type Libvirt
type Apic struct {
}

// BackingStore type Libvirt
type BackingStore struct {
	AttrIndex    string        `xml:" index,attr"  json:",omitempty"`
	AttrType     string        `xml:" type,attr"  json:",omitempty"`
	BackingStore *BackingStore `xml:" backingStore,omitempty" json:"backingStore,omitempty"`
	Format       *Format       `xml:" format,omitempty" json:"format,omitempty"`
	Source       *Source       `xml:" source,omitempty" json:"source,omitempty"`
}

// Boot type Libvirt
type Boot struct {
	AttrDev string `xml:" dev,attr"  json:",omitempty"`
}

// Clock type Libvirt
type Clock struct {
	AttrOffset string   `xml:" offset,attr"  json:",omitempty"`
	Timer      []*Timer `xml:" timer,omitempty" json:"timer,omitempty"`
}

// Console type Libvirt
type Console struct {
	AttrType string  `xml:" type,attr"  json:",omitempty"`
	Alias    *Alias  `xml:" alias,omitempty" json:"alias,omitempty"`
	Source   *Source `xml:" source,omitempty" json:"source,omitempty"`
	Target   *Target `xml:" target,omitempty" json:"target,omitempty"`
}

// Controller type Libvirt
type Controller struct {
	AttrIndex string   `xml:" index,attr"  json:",omitempty"`
	AttrModel string   `xml:" model,attr"  json:",omitempty"`
	AttrType  string   `xml:" type,attr"  json:",omitempty"`
	Address   *Address `xml:" address,omitempty" json:"address,omitempty"`
	Alias     *Alias   `xml:" alias,omitempty" json:"alias,omitempty"`
}

// CPU type Libvirt
type CPU struct {
	AttrMode string    `xml:" mode,attr"  json:",omitempty"`
	Model    *Model    `xml:" model,omitempty" json:"model,omitempty"`
	Topology *Topology `xml:" topology,omitempty" json:"topology,omitempty"`
}

// CPUtune type Libvirt
type CPUtune struct {
	Shares *Shares `xml:" shares,omitempty" json:"shares,omitempty"`
}

// CurrentMemory type Libvirt
type CurrentMemory struct {
	AttrUnit string `xml:" unit,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

// Devices type Libvirt
type Devices struct {
	Console    *Console      `xml:" console,omitempty" json:"console,omitempty"`
	Controller []*Controller `xml:" controller,omitempty" json:"controller,omitempty"`
	Disk       []*Disk       `xml:" disk,omitempty" json:"disk,omitempty"`
	Emulator   *Emulator     `xml:" emulator,omitempty" json:"emulator,omitempty"`
	Graphics   *Graphics     `xml:" graphics,omitempty" json:"graphics,omitempty"`
	Input      []*Input      `xml:" input,omitempty" json:"input,omitempty"`
	Interface  []*Interface  `xml:" interface,omitempty" json:"interface,omitempty"`
	Memballoon *Memballoon   `xml:" memballoon,omitempty" json:"memballoon,omitempty"`
	Serial     []*Serial     `xml:" serial,omitempty" json:"serial,omitempty"`
	Video      *Video        `xml:" video,omitempty" json:"video,omitempty"`
}

// Disk type Libvirt
type Disk struct {
	AttrDevice   string        `xml:" device,attr"  json:",omitempty"`
	AttrType     string        `xml:" type,attr"  json:",omitempty"`
	Address      *Address      `xml:" address,omitempty" json:"address,omitempty"`
	Alias        *Alias        `xml:" alias,omitempty" json:"alias,omitempty"`
	BackingStore *BackingStore `xml:" backingStore,omitempty" json:"backingStore,omitempty"`
	Driver       *Driver       `xml:" driver,omitempty" json:"driver,omitempty"`
	Source       *Source       `xml:" source,omitempty" json:"source,omitempty"`
	Target       *Target       `xml:" target,omitempty" json:"target,omitempty"`
}

// Domain type Libvirt
type Domain struct {
	AttrID        string         `xml:" id,attr"  json:",omitempty"`
	AttrType      string         `xml:" type,attr"  json:",omitempty"`
	Clock         *Clock         `xml:" clock,omitempty" json:"clock,omitempty"`
	CPU           *CPU           `xml:" cpu,omitempty" json:"cpu,omitempty"`
	CPUtune       *CPUtune       `xml:" cputune,omitempty" json:"cputune,omitempty"`
	CurrentMemory *CurrentMemory `xml:" currentMemory,omitempty" json:"currentMemory,omitempty"`
	Devices       *Devices       `xml:" devices,omitempty" json:"devices,omitempty"`
	Features      *Features      `xml:" features,omitempty" json:"features,omitempty"`
	Memory        *Memory        `xml:" memory,omitempty" json:"memory,omitempty"`
	Metadata      *Metadata      `xml:" metadata,omitempty" json:"metadata,omitempty"`
	Name          string         `xml:" name,omitempty" json:"name,omitempty"`
	OnCrash       *OnCrash       `xml:" onCrash,omitempty" json:"onCrash,omitempty"`
	OnPoweroff    *OnPoweroff    `xml:" onPoweroff,omitempty" json:"onPoweroff,omitempty"`
	OnReboot      *OnReboot      `xml:" onReboot,omitempty" json:"onReboot,omitempty"`
	Os            *Os            `xml:" os,omitempty" json:"os,omitempty"`
	Resource      *Resource      `xml:" resource,omitempty" json:"resource,omitempty"`
	Seclabel      *Seclabel      `xml:" seclabel,omitempty" json:"seclabel,omitempty"`
	Sysinfo       *Sysinfo       `xml:" sysinfo,omitempty" json:"sysinfo,omitempty"`
	UUID          *UUID          `xml:" uuid,omitempty" json:"uuid,omitempty"`
	Vcpu          *Vcpu          `xml:" vcpu,omitempty" json:"vcpu,omitempty"`
}

// Driver type Libvirt
type Driver struct {
	AttrCache string `xml:" cache,attr"  json:",omitempty"`
	AttrName  string `xml:" name,attr"  json:",omitempty"`
	AttrType  string `xml:" type,attr"  json:",omitempty"`
}

// Emulator type Libvirt
type Emulator struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Entry type Libvirt
type Entry struct {
	AttrName string `xml:" name,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

// Features type Libvirt
type Features struct {
	Acpi *Acpi `xml:" acpi,omitempty" json:"acpi,omitempty"`
	Apic *Apic `xml:" apic,omitempty" json:"apic,omitempty"`
}

// Format type Libvirt
type Format struct {
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

// Graphics type Libvirt
type Graphics struct {
	AttrAutoport string  `xml:" autoport,attr"  json:",omitempty"`
	AttrKeymap   string  `xml:" keymap,attr"  json:",omitempty"`
	AttrListen   string  `xml:" listen,attr"  json:",omitempty"`
	AttrPort     string  `xml:" port,attr"  json:",omitempty"`
	AttrType     string  `xml:" type,attr"  json:",omitempty"`
	Listen       *Listen `xml:" listen,omitempty" json:"listen,omitempty"`
}

// Imagelabel type Libvirt
type Imagelabel struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Input type Libvirt
type Input struct {
	AttrBus  string `xml:" bus,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
	Alias    *Alias `xml:" alias,omitempty" json:"alias,omitempty"`
}

// Interface type Libvirt
type Interface struct {
	AttrType string   `xml:" type,attr"  json:",omitempty"`
	Address  *Address `xml:" address,omitempty" json:"address,omitempty"`
	Alias    *Alias   `xml:" alias,omitempty" json:"alias,omitempty"`
	Mac      *Mac     `xml:" mac,omitempty" json:"mac,omitempty"`
	Model    *Model   `xml:" model,omitempty" json:"model,omitempty"`
	Source   *Source  `xml:" source,omitempty" json:"source,omitempty"`
	Target   *Target  `xml:" target,omitempty" json:"target,omitempty"`
}

// Label type Libvirt
type Label struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Listen type Libvirt
type Listen struct {
	AttrAddress string `xml:" address,attr"  json:",omitempty"`
	AttrType    string `xml:" type,attr"  json:",omitempty"`
}

// Mac type Libvirt
type Mac struct {
	AttrAddress string `xml:" address,attr"  json:",omitempty"`
}

// Memballoon type Libvirt
type Memballoon struct {
	AttrModel string   `xml:" model,attr"  json:",omitempty"`
	Address   *Address `xml:" address,omitempty" json:"address,omitempty"`
	Alias     *Alias   `xml:" alias,omitempty" json:"alias,omitempty"`
	Stats     *Stats   `xml:" stats,omitempty" json:"stats,omitempty"`
}

// Memory type Libvirt
type Memory struct {
	AttrUnit string `xml:" unit,attr"  json:",omitempty"`
	Text     string `xml:",chardata" json:",omitempty"`
}

// Metadata type Libvirt
type Metadata struct {
	NovaInstance *NovaInstance `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 instance,omitempty" json:"instance,omitempty"`
}

// Model type Libvirt
type Model struct {
	AttrFallback string `xml:" fallback,attr"  json:",omitempty"`
	AttrHeads    string `xml:" heads,attr"  json:",omitempty"`
	AttrType     string `xml:" type,attr"  json:",omitempty"`
	AttrVram     string `xml:" vram,attr"  json:",omitempty"`
}

// Name type Libvirt
type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// OnCrash type Libvirt
type OnCrash struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// OnPoweroff type Libvirt
type OnPoweroff struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// OnReboot type Libvirt
type OnReboot struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Os type Libvirt
type Os struct {
	Boot   *Boot   `xml:" boot,omitempty" json:"boot,omitempty"`
	Smbios *Smbios `xml:" smbios,omitempty" json:"smbios,omitempty"`
	Type   *Type   `xml:" type,omitempty" json:"type,omitempty"`
}

// Partition type Libvirt
type Partition struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Resource type Libvirt
type Resource struct {
	Partition *Partition `xml:" partition,omitempty" json:"partition,omitempty"`
}

// Root type Libvirt
type Root struct {
	Domain *Domain `xml:" domain,omitempty" json:"domain,omitempty"`
}

// Seclabel type Libvirt
type Seclabel struct {
	AttrModel   string      `xml:" model,attr"  json:",omitempty"`
	AttrRelabel string      `xml:" relabel,attr"  json:",omitempty"`
	AttrType    string      `xml:" type,attr"  json:",omitempty"`
	Imagelabel  *Imagelabel `xml:" imagelabel,omitempty" json:"imagelabel,omitempty"`
	Label       *Label      `xml:" label,omitempty" json:"label,omitempty"`
}

// Serial type Libvirt
type Serial struct {
	AttrType string  `xml:" type,attr"  json:",omitempty"`
	Alias    *Alias  `xml:" alias,omitempty" json:"alias,omitempty"`
	Source   *Source `xml:" source,omitempty" json:"source,omitempty"`
	Target   *Target `xml:" target,omitempty" json:"target,omitempty"`
}

// Shares type Libvirt
type Shares struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Smbios type Libvirt
type Smbios struct {
	AttrMode string `xml:" mode,attr"  json:",omitempty"`
}

// Source type Libvirt
type Source struct {
	AttrBridge string `xml:" bridge,attr"  json:",omitempty"`
	AttrFile   string `xml:" file,attr"  json:",omitempty"`
	AttrPath   string `xml:" path,attr"  json:",omitempty"`
}

// Stats type Libvirt
type Stats struct {
	AttrPeriod string `xml:" period,attr"  json:",omitempty"`
}

// Sysinfo type Libvirt
type Sysinfo struct {
	AttrType string  `xml:" type,attr"  json:",omitempty"`
	System   *System `xml:" system,omitempty" json:"system,omitempty"`
}

// System type Libvirt
type System struct {
	Entry []*Entry `xml:" entry,omitempty" json:"entry,omitempty"`
}

// Target type Libvirt
type Target struct {
	AttrBus  string `xml:" bus,attr"  json:",omitempty"`
	AttrDev  string `xml:" dev,attr"  json:",omitempty"`
	AttrPort string `xml:" port,attr"  json:",omitempty"`
	AttrType string `xml:" type,attr"  json:",omitempty"`
}

// Timer type Libvirt
type Timer struct {
	AttrName       string `xml:" name,attr"  json:",omitempty"`
	AttrPresent    string `xml:" present,attr"  json:",omitempty"`
	AttrTickpolicy string `xml:" tickpolicy,attr"  json:",omitempty"`
}

// Topology type Libvirt
type Topology struct {
	AttrCores   string `xml:" cores,attr"  json:",omitempty"`
	AttrSockets string `xml:" sockets,attr"  json:",omitempty"`
	AttrThreads string `xml:" threads,attr"  json:",omitempty"`
}

// Type type Libvirt
type Type struct {
	AttrArch    string `xml:" arch,attr"  json:",omitempty"`
	AttrMachine string `xml:" machine,attr"  json:",omitempty"`
	Text        string `xml:",chardata" json:",omitempty"`
}

// UUID type Libvirt
type UUID struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// Vcpu type Libvirt
type Vcpu struct {
	AttrPlacement string `xml:" placement,attr"  json:",omitempty"`
	Text          string `xml:",chardata" json:",omitempty"`
}

// Video type Libvirt
type Video struct {
	Address *Address `xml:" address,omitempty" json:"address,omitempty"`
	Alias   *Alias   `xml:" alias,omitempty" json:"alias,omitempty"`
	Model   *Model   `xml:" model,omitempty" json:"model,omitempty"`
}

// NovaCreationTime type Libvirt
type NovaCreationTime struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 creationTime,omitempty" json:"creationTime,omitempty"`
}

// NovaDisk type Libvirt
type NovaDisk struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 disk,omitempty" json:"disk,omitempty"`
}

// NovaEphemeral type Libvirt
type NovaEphemeral struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 ephemeral,omitempty" json:"ephemeral,omitempty"`
}

// NovaFlavor type Libvirt
type NovaFlavor struct {
	AttrName      string         `xml:" name,attr"  json:",omitempty"`
	NovaDisk      *NovaDisk      `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 disk,omitempty" json:"disk,omitempty"`
	NovaEphemeral *NovaEphemeral `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 ephemeral,omitempty" json:"ephemeral,omitempty"`
	NovaMemory    *NovaMemory    `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 memory,omitempty" json:"memory,omitempty"`
	NovaSwap      *NovaSwap      `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 swap,omitempty" json:"swap,omitempty"`
	NovaVcpus     *NovaVcpus     `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 vcpus,omitempty" json:"vcpus,omitempty"`
	XMLName       xml.Name       `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 flavor,omitempty" json:"flavor,omitempty"`
}

// NovaInstance type Libvirt
type NovaInstance struct {
	AttrNova         string       `xml:"xmlns nova,attr"  json:",omitempty"`
	NovaCreationTime string       `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 creationTime,omitempty" json:"creationTime,omitempty"`
	NovaFlavor       *NovaFlavor  `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 flavor,omitempty" json:"flavor,omitempty"`
	NovaName         *NovaName    `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 name,omitempty" json:"name,omitempty"`
	NovaOwner        *NovaOwner   `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 owner,omitempty" json:"owner,omitempty"`
	NovaPackage      *NovaPackage `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 package,omitempty" json:"package,omitempty"`
	NovaRoot         *NovaRoot    `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 root,omitempty" json:"root,omitempty"`
	XMLName          xml.Name     `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 instance,omitempty" json:"instance,omitempty"`
}

// NovaMemory type Libvirt
type NovaMemory struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 memory,omitempty" json:"memory,omitempty"`
}

// NovaName type Libvirt
type NovaName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 name,omitempty" json:"name,omitempty"`
}

// NovaOwner type Libvirt
type NovaOwner struct {
	NovaProject *NovaProject `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 project,omitempty" json:"project,omitempty"`
	NovaUser    *NovaUser    `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 user,omitempty" json:"user,omitempty"`
	XMLName     xml.Name     `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 owner,omitempty" json:"owner,omitempty"`
}

// NovaPackage type Libvirt
type NovaPackage struct {
	AttrVersion string   `xml:" version,attr"  json:",omitempty"`
	XMLName     xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 package,omitempty" json:"package,omitempty"`
}

// NovaProject type Libvirt
type NovaProject struct {
	AttrUUID string   `xml:" uuid,attr"  json:",omitempty"`
	Text     string   `xml:",chardata" json:",omitempty"`
	XMLName  xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 project,omitempty" json:"project,omitempty"`
}

// NovaRoot type Libvirt
type NovaRoot struct {
	AttrType string   `xml:" type,attr"  json:",omitempty"`
	AttrUUID string   `xml:" uuid,attr"  json:",omitempty"`
	XMLName  xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 root,omitempty" json:"root,omitempty"`
}

// NovaSwap type Libvirt
type NovaSwap struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 swap,omitempty" json:"swap,omitempty"`
}

// NovaUser type Libvirt
type NovaUser struct {
	AttrUUID string   `xml:" uuid,attr"  json:",omitempty"`
	Text     string   `xml:",chardata" json:",omitempty"`
	XMLName  xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 user,omitempty" json:"user,omitempty"`
}

// NovaVcpus type Libvirt
type NovaVcpus struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://openstack.org/xmlns/libvirt/nova/1.0 vcpus,omitempty" json:"vcpus,omitempty"`
}

//NovaMetadataExist check whether Nova Metadata exist
func (domain *Domain) NovaMetadataExist() bool {
	if domain.Metadata != nil && domain.Metadata.NovaInstance != nil {
		return true
	}
	return false
}

//GetNovaFlavor returns Nova Flavor Name
func (domain *Domain) GetNovaFlavor() string {

	if domain.Metadata.NovaInstance.NovaFlavor != nil {
		if domain.Metadata.NovaInstance.NovaFlavor.AttrName != "" {
			return domain.Metadata.NovaInstance.NovaFlavor.AttrName
		}
	}
	return ""
}

//GetNovaDisk returns Nova Disk Size
func (domain *Domain) GetNovaDisk() string {
	if domain.Metadata.NovaInstance.NovaFlavor != nil && domain.Metadata.NovaInstance.NovaFlavor.NovaDisk != nil {
		return domain.Metadata.NovaInstance.NovaFlavor.NovaDisk.Text
	}
	return ""
}

//GetNovaEphemeral return Nova Ephemeral size
func (domain *Domain) GetNovaEphemeral() string {

	if domain.Metadata.NovaInstance.NovaFlavor != nil && domain.Metadata.NovaInstance.NovaFlavor.NovaEphemeral != nil {
		return domain.Metadata.NovaInstance.NovaFlavor.NovaEphemeral.Text
	}

	return ""
}

//GetNovaMemory returns Nova memory size
func (domain *Domain) GetNovaMemory() string {

	if domain.Metadata.NovaInstance.NovaFlavor.NovaMemory != nil {
		return domain.Metadata.NovaInstance.NovaFlavor.NovaMemory.Text
	}

	return ""
}

//GetNovaSwap return Nova memory swap size
func (domain *Domain) GetNovaSwap() string {

	if domain.Metadata.NovaInstance.NovaFlavor.NovaSwap != nil {
		return domain.Metadata.NovaInstance.NovaFlavor.NovaSwap.Text
	}

	return ""
}

//GetNovaVcpus returns Nova vcpus count
func (domain *Domain) GetNovaVcpus() string {

	if domain.Metadata.NovaInstance.NovaFlavor.NovaVcpus != nil {
		return domain.Metadata.NovaInstance.NovaFlavor.NovaVcpus.Text
	}

	return ""
}

//GetNovaOwner returns Owner of domain
func (domain *Domain) GetNovaOwner() string {

	if domain.Metadata.NovaInstance.NovaOwner != nil && domain.Metadata.NovaInstance.NovaOwner.NovaUser != nil {
		return domain.Metadata.NovaInstance.NovaOwner.NovaUser.AttrUUID
	}
	return ""
}

//GetNovaName returns name of the domain
func (domain *Domain) GetNovaName() string {

	if domain.Metadata.NovaInstance.NovaName != nil {
		return domain.Metadata.NovaInstance.NovaName.Text
	}
	return ""
}

//GetNovaPackage returns package version of Nova service
func (domain *Domain) GetNovaPackage() string {

	if domain.Metadata.NovaInstance.NovaPackage != nil {
		return domain.Metadata.NovaInstance.NovaPackage.AttrVersion
	}
	return ""
}

//GetNovaUUID returns NovaUUID
func (domain *Domain) GetNovaUUID() string {

	if domain.UUID != nil {
		return domain.UUID.Text
	}
	return ""
}
