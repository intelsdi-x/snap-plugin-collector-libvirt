package libvirt

import (
	"encoding/xml"
	"errors"
	"strconv"

	"fmt"

	"github.com/sandlbn/libvirt-go"
)

var memInfo = struct {
	m map[string]int32
}{m: map[string]int32{
	"swap_in":        0,
	"swap_out":       1,
	"major_fault":    2,
	"min_fault":      3,
	"unused":         4,
	"available":      5,
	"actual_balloon": 6,
	"rss":            7,
	"nr":             8,
}}

// GetNetworkStatistics return GoBlockStatistics
func GetNetworkStatistics(domain libvirt.VirDomain, paths ...string) (map[string]libvirt.VirDomainInterfaceStats, error) {
	netStat := make(map[string]libvirt.VirDomainInterfaceStats)
	var err error
	if len(paths) == 0 {
		paths, err = GetDomainInterfaces(domain)
		if err != nil {
			return netStat, err
		}
	}
	for _, path := range paths {
		stat, err := domain.InterfaceStats(path)
		if err != nil {
			return netStat, err
		}
		netStat[path] = stat
	}
	return netStat, nil
}

// GetBlockStatistics return GoBlockStatistics
func GetBlockStatistics(domain libvirt.VirDomain, paths ...string) (map[string]libvirt.VirDomainBlockStats, error) {
	blockStat := make(map[string]libvirt.VirDomainBlockStats)
	var err error
	if len(paths) == 0 {
		paths, err = GetDomainDisks(domain)
		if err != nil {
			return blockStat, err
		}
	}
	for _, path := range paths {
		stat, err := domain.BlockStats(path)
		if err != nil {
			return blockStat, err
		}
		blockStat[path] = stat
	}
	return blockStat, nil
}

//GetDomainInterfaces returns slice of available network devices
func GetDomainInterfaces(domain libvirt.VirDomain) ([]string, error) {
	var interfaces []string
	domXML, err := getDomainXML(domain)
	if err != nil {
		return interfaces, err
	}
	if domXML.Devices == nil || domXML.Devices.Interface == nil {
		errMsg := fmt.Sprintf("Error getting interface info, from domain: %s", domXML.Name)
		return interfaces, errors.New(errMsg)
	}
	interfaces = append(interfaces, domXML.Devices.Interface.Target.AttrDev)
	return interfaces, nil
}

//GetCPUStatistics returns cpu  statistics
func GetCPUStatistics(domain libvirt.VirDomain) (int64, error) {
	info, err := domain.GetInfo()
	if err != nil {
		return 0, err
	}
	return int64(info.GetCpuTime()), nil
}

//GetVCPUStatistics returns vcpu statistics
func GetVCPUStatistics(domain libvirt.VirDomain) (map[string]int64, error) {
	vcpuStat := make(map[string]int64)

	info, err := domain.GetInfo()
	if err != nil {
		return vcpuStat, err
	}
	vcpus, err := domain.GetVcpus(int32(info.GetNrVirtCpu()))
	if err != nil {
		return vcpuStat, err
	}
	for k, v := range vcpus {
		vcpuStat[strconv.Itoa(k)] = int64(v.CpuTime)
	}
	return vcpuStat, nil
}

//GetMemoryStatistics returns Libvirt memory statistics
func GetMemoryStatistics(domain libvirt.VirDomain, tags ...string) (map[string]int64, error) {
	retValue := make(map[string]int64)
	info, err := domain.MemoryStats(5, 0)
	if err != nil {
		return retValue, err
	}

	for _, v := range tags {
		retValue[v] = parseMemStats(info, memInfo.m[v])
	}
	return retValue, nil
}

//GetDomainDisks returns slice of available disk devices
func GetDomainDisks(domain libvirt.VirDomain) ([]string, error) {
	var disks []string
	domXML, err := getDomainXML(domain)
	if err != nil {
		return disks, err
	}
	if domXML.Devices != nil {
		for _, disk := range domXML.Devices.Disk {
			if disk.AttrDevice == "disk" {
				disks = append(disks, disk.Target.AttrDev)
			}
		}
	}

	return disks, nil
}

func parseMemStats(memstat []libvirt.VirDomainMemoryStat, nr int32) int64 {

	for i := 0; i < len(memstat); i++ {
		if memstat[i].Tag == nr {
			return int64(memstat[i].Val)
		}
	}
	return 0

}

func getDomainXML(domain libvirt.VirDomain) (*Domain, error) {
	var libvirtXML Domain
	xmlDesc, err := domain.GetXMLDesc(0)
	if err != nil {
		return &Domain{}, err
	}
	xml.Unmarshal([]byte(xmlDesc), &libvirtXML)
	return &libvirtXML, nil
}

// GetNovaMetadata return Nova Metadata info
func GetNovaMetadata(domain libvirt.VirDomain) (map[string]string, error) {
	domXML, err := getDomainXML(domain)
	if err != nil {
		return nil, err
	}
	novaMetadata := generateNovaMetadata(domXML)

	return novaMetadata, nil
}

func generateNovaMetadata(domXML *Domain) map[string]string {
	novaMetadata := make(map[string]string)
	if domXML.NovaMetadataExist() {

		novaMetadata["nova_flavor"] = domXML.GetNovaFlavor()
		novaMetadata["nova_disk_mb"] = domXML.GetNovaDisk()
		novaMetadata["nova_ephemeral_mb"] = domXML.GetNovaEphemeral()
		novaMetadata["nova_memory_mb"] = domXML.GetNovaMemory()
		novaMetadata["nova_swap_mb"] = domXML.GetNovaSwap()
		novaMetadata["nova_vcpus"] = domXML.GetNovaVcpus()
		novaMetadata["nova_name"] = domXML.GetNovaName()
		novaMetadata["nova_owner"] = domXML.GetNovaOwner()
		novaMetadata["nova_package"] = domXML.GetNovaPackage()

	}

	if domXML.UUID != nil {
		novaMetadata["nova_uuid"] = domXML.UUID.Text
	}
	return novaMetadata
}
