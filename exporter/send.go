package exporter

import "github.com/prometheus/client_golang/prometheus"

type DefaultMetricsSender struct{}

func (df *DefaultMetricsSender) SendProbeFailure(ch chan<- prometheus.Metric) {
	SendProbeStatus(ch, false)
}

func (df *DefaultMetricsSender) SendMetrics(ch chan<- prometheus.Metric, metrics *Metrics, opts *GitlabScrapeOpts) {
	SendPipelineCountByStatus(ch, metrics, opts)
	SendProbeStatus(ch, true)
	SendLatestDuration(ch, metrics.LatestDuration)
	SendProbeDuration(ch, metrics.ProbeDuration)
}

func SendPipelineCountByStatus(ch chan<- prometheus.Metric, metrics *Metrics, opts *GitlabScrapeOpts) {
	counts := metrics.Count

	SendCount(ch, counts.Success, "success", opts)

	SendCount(ch, counts.Failed, "failed", opts)

	SendCount(ch, counts.Cancelled, "cancelled", opts)

	SendCount(ch, counts.Manual, "manual", opts)

	SendCount(ch, counts.Pending, "pending", opts)

	SendCount(ch, counts.Preparing, "preparing", opts)

	SendCount(ch, counts.Running, "running", opts)

	SendCount(ch, counts.Scheduled, "scheduled", opts)

	SendCount(ch, counts.Skipped, "skipped", opts)

	SendCount(ch, counts.WaitingForResource, "waiting_for_resource", opts)
}

func SendCount(ch chan<- prometheus.Metric, val float64, status string, opts *GitlabScrapeOpts) {
	ch <- prometheus.MustNewConstMetric(
		pipelineTotalDesc,
		prometheus.CounterValue,
		val,
		opts.group,
		opts.project,
		status,
	)
}

func SendProbeStatus(ch chan<- prometheus.Metric, success bool) {
	var value float64
	if success {
		value = 1
	} else {
		value = 0
	}

	ch <- prometheus.MustNewConstMetric(
		successDesc,
		prometheus.GaugeValue,
		value,
	)
}

func SendLatestDuration(ch chan<- prometheus.Metric, duration float64) {
	ch <- prometheus.MustNewConstMetric(
		latestDurationDesc,
		prometheus.GaugeValue,
		duration,
	)
}

func SendProbeDuration(ch chan<- prometheus.Metric, duration float64) {
	ch <- prometheus.MustNewConstMetric(
		probeDurationDesc,
		prometheus.GaugeValue,
		duration,
	)
}
