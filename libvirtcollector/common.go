package libvirtcollector

import (
	"bytes"
	"strings"
	"time"

	wrapper "github.com/intelsdi-x/snap-plugin-collector-libvirt/libvirt"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	libvirt "github.com/sandlbn/libvirt-go"
)

func getLibvirtURI(cfg plugin.Config) string {
	uri, err := cfg.GetString("uri")
	if err != nil {
		return defaultURI
	}
	return uri
}

func createMetric(ns plugin.Namespace) plugin.Metric {
	metricType := plugin.Metric{

		Namespace: ns,
		Version:   Version,
	}
	return metricType
}

func filterNamespace(metricType string, mts []plugin.Metric) (int, []plugin.Metric) {
	filteredMetrics := []plugin.Metric{}
	for _, m := range mts {
		if m.Namespace.Strings()[nsMetricPostion] == metricType {
			filteredMetrics = append(filteredMetrics, m)
		}
	}
	filteredMetrics = removeDuplicates(filteredMetrics)
	return len(filteredMetrics), filteredMetrics
}

func merge(maps ...map[string]string) (output map[string]string) {
	size := len(maps)
	if size == 0 {
		return output
	}
	if size == 1 {
		return maps[0]
	}
	output = make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			output[k] = v
		}
	}
	return output
}

func createNamespace(mt plugin.Metric, value interface{}, ns plugin.Namespace, meta map[string]string) plugin.Metric {

	return plugin.Metric{
		Timestamp: time.Now(),
		Namespace: ns,
		Data:      value,
		Tags:      merge(mt.Tags, meta),
		Config:    mt.Config,
		Version:   Version,
	}
}

func copyNamespace(mt plugin.Metric) []plugin.NamespaceElement {
	ns := make([]plugin.NamespaceElement, len(mt.Namespace))
	copy(ns, mt.Namespace)
	return ns
}

func copyNamespaceElements(ns []plugin.NamespaceElement) []plugin.NamespaceElement {
	newNs := make([]plugin.NamespaceElement, len(ns))
	copy(newNs, ns)
	return newNs
}

func removeDuplicates(elements []plugin.Metric) []plugin.Metric {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []plugin.Metric{}

	for _, v := range elements {
		ns := strings.Join(v.Namespace.Strings(), "/")
		if !encountered[ns] {
			encountered[ns] = true
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}

func metricStored(elements []plugin.Metric, newNamespace []plugin.NamespaceElement) bool {
	for _, v := range elements {
		ns := strings.Join(v.Namespace.Strings(), "")
		if ns == joinNamespaceElements(newNamespace) {
			return true
		}
	}
	return false
}

func joinNamespaceElements(ns []plugin.NamespaceElement) string {
	var buffer bytes.Buffer
	for _, v := range ns {
		buffer.WriteString(v.Value)
	}
	return buffer.String()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getInstances(conn libvirt.VirConnection, elements []plugin.Metric) ([]libvirt.VirDomain, error) {

	instances := []string{}

	for _, v := range elements {
		domain := v.Namespace.Strings()[nsDomainPosition]

		if domain == "*" {
			return wrapper.GetInstances(conn)
		}
		if !contains(instances, domain) {
			instances = append(instances, domain)
		}
	}
	return wrapper.GetRequestedInstances(conn, instances)

}
