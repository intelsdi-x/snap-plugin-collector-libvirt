package libvirtcollector

import (
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

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
