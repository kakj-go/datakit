package man

import (
	"embed"
)

var (
	//go:embed manuals/*.md
	docs embed.FS
)

var (
	OtherDocs = map[string]bool{
		"datakit-sink-guide":       true,
		"datakit-sink-dev":         true,
		"datakit-sink-influxdb":    true,
		"datakit-sink-logstash":    true,
		"datakit-sink-m3db":        true,
		"datakit-sink-otel-jaeger": true,
		"apis":                     true,
		"changelog":                true,
		"datakit-arch":             true,
		"datakit-batch-deploy":     true,
		"datakit-conf":             true,
		"datakit-input-conf":       true,
		"datakit-daemonset-deploy": true,
		"datakit-daemonset-update": true,
		"datakit-daemonset-bp":     true,
		"datakit-dql-how-to":       true,
		"datakit-filter":           true,
		"datakit-logging-how":      true,
		"datakit-install":          true,
		"datakit-logging":          true,
		"datakit-monitor":          true,
		"datakit-offline-install":  true,
		"datakit-on-public":        true,
		"datakit-pl-how-to":        true,
		"datakit-pl-global":        true,
		"datakit-service-how-to":   true,
		"datakit-tools-how-to":     true,
		"datakit-tracing":          true,
		"datakit-tracing-struct":   true,
		"datakit-update":           true,
		"datatypes":                true,
		"dataway":                  true,
		"dca":                      true,
		"ddtrace-golang":           true,
		"ddtrace-java":             true,
		"ddtrace-python":           true,
		"ddtrace-php":              true,
		"ddtrace-nodejs":           true,
		"ddtrace-cpp":              true,
		"ddtrace-ruby":             true,
		"development":              true,
		"dialtesting_json":         true,
		"election":                 true,
		"k8s-config-how-to":        true,
		"kubernetes-prom":          true,
		"kubernetes-crd":           true,
		"kubernetes-x":             true,
		"logfwd":                   true,
		"logging-pipeline-bench":   true,
		"logging_socket":           true,
		"opentelemetry-go":         true,
		"opentelemetry-java":       true,
		"pipeline":                 true,
		"prometheus":               true,
		"rum":                      true,
		"sec-checker":              true,
		"telegraf":                 true,
		"why-no-data":              true,
		"git-config-how-to":        true,
	}
)
