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

import "github.com/sandlbn/libvirt-go"

// Libvirt interface for all high level libvirt functions
type Libvirt interface {
	GetInstancesIds() ([]string, error)
	GetInstances() ([]libvirt.VirDomain, error)
	GetNetworkStatistics(domain libvirt.VirDomain, paths ...string)
	GetBlockStatistics(domain libvirt.VirDomain, paths ...string)
	GetCPUStatistics(domain libvirt.VirDomain) (uint64, error)
	GetMemoryStatistics(domain libvirt.VirDomain, tags ...string) (map[string]uint64, error)
	GetDomainInterfaces(domain libvirt.VirDomain) ([]string, error)
	GetDomainDisks(domain libvirt.VirDomain) ([]string, error)
	GetVCPUStatistics(domain libvirt.VirDomain) (map[string]int64, error)
	GetRequestedInstances(conn libvirt.VirConnection, domainNames []string) ([]libvirt.VirDomain, error)
	GetInstanceByDomainName(conn libvirt.VirConnection, domainName string) (libvirt.VirDomain, error)
}

// GetInstanceIds return all names of active VirDomains
func GetInstanceIds(conn libvirt.VirConnection) ([]string, error) {
	var instanceIds []string
	domains, err := GetInstances(conn)

	if err != nil {
		return instanceIds, err
	}
	for _, domain := range domains {
		instanceName, err := domain.GetName()
		if err != nil {
			return instanceIds, err
		}
		instanceIds = append(instanceIds, instanceName)
	}
	return instanceIds, nil
}

// GetInstances return all active VirDomains
func GetInstances(conn libvirt.VirConnection) ([]libvirt.VirDomain, error) {
	var libvirtDomains []libvirt.VirDomain
	domains, err := conn.ListDomains()
	if err != nil {
		return libvirtDomains, err
	}
	for i := 0; i < len(domains); i++ {
		domain, err := conn.LookupDomainById(domains[i])
		if err != nil {
			return libvirtDomains, err
		}
		libvirtDomains = append(libvirtDomains, domain)
	}
	return libvirtDomains, nil
}

// GetRequestedInstances return all instances from domainNames slice
func GetRequestedInstances(conn libvirt.VirConnection, domainNames []string) ([]libvirt.VirDomain, error) {
	var libvirtDomains []libvirt.VirDomain
	for _, domainName := range domainNames {
		domain, err := GetInstanceByDomainName(conn, domainName)
		if err != nil {
			return libvirtDomains, err
		}
		libvirtDomains = append(libvirtDomains, domain)
	}
	return libvirtDomains, nil

}

// GetInstanceByDomainName return instance by DomainName
func GetInstanceByDomainName(conn libvirt.VirConnection, domainName string) (libvirt.VirDomain, error) {
	return conn.LookupDomainByName(domainName)
}
