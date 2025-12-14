package internal

import "github.com/prometheus/client_golang/prometheus"

func Describe(ch chan<- *prometheus.Desc) {
	ch <- pipelineTotalDesc
	ch <- successDesc
	ch <- latestDurationDesc
	ch <- probeDurationDesc
}

var pipelineTotalDesc = prometheus.NewDesc(
	"gitlab_pipeline_total",
	"Total number of pipelines for the target project.",
	[]string{"group", "project", "status"},
	nil,
)

var successDesc = prometheus.NewDesc(
	"gitlab_probe_success",
	"Whether the probe was successful (1 - success, 0 - failure).",
	nil,
	nil,
)

var latestDurationDesc = prometheus.NewDesc(
	"gitlab_pipeline_last_duration_seconds",
	"Duration of the latest pipeline in seconds.",
	nil,
	nil,
)

var probeDurationDesc = prometheus.NewDesc(
	"exporter_probe_duration_seconds",
	"Duration in seconds of the probe",
	nil,
	nil,
)
