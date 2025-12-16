package internal

import "github.com/prometheus/client_golang/prometheus"

type DefaultMetricsSender struct{}

func (df *DefaultMetricsSender) SendProbeFailure(ch chan<- prometheus.Metric) {
	sendProbeStatus(ch, false)
}

func (df *DefaultMetricsSender) SendMetrics(ch chan<- prometheus.Metric, metrics *Metrics, opts *GitlabScrapeOpts) {
	sendPipelineCountByStatus(ch, metrics, opts)
	sendProbeStatus(ch, true)
	sendLatestDuration(ch, metrics.LatestDuration)
	sendProbeDuration(ch, metrics.ProbeDuration)
	sendGitlabHost(ch, metrics.GitlabHost)
}

func sendPipelineCountByStatus(ch chan<- prometheus.Metric, metrics *Metrics, opts *GitlabScrapeOpts) {
	counts := metrics.Count

	sendCount(ch, counts.Success, "success", opts)

	sendCount(ch, counts.Failed, "failed", opts)

	sendCount(ch, counts.Cancelled, "cancelled", opts)

	sendCount(ch, counts.Manual, "manual", opts)

	sendCount(ch, counts.Pending, "pending", opts)

	sendCount(ch, counts.Preparing, "preparing", opts)

	sendCount(ch, counts.Running, "running", opts)

	sendCount(ch, counts.Scheduled, "scheduled", opts)

	sendCount(ch, counts.Skipped, "skipped", opts)

	sendCount(ch, counts.WaitingForResource, "waiting_for_resource", opts)
}

func sendCount(ch chan<- prometheus.Metric, val float64, status string, opts *GitlabScrapeOpts) {
	ch <- prometheus.MustNewConstMetric(
		pipelineTotalDesc,
		prometheus.CounterValue,
		val,
		opts.Group,
		opts.Project,
		status,
	)
}

func sendProbeStatus(ch chan<- prometheus.Metric, success bool) {
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

func sendLatestDuration(ch chan<- prometheus.Metric, duration float64) {
	ch <- prometheus.MustNewConstMetric(
		latestDurationDesc,
		prometheus.GaugeValue,
		duration,
	)
}

func sendProbeDuration(ch chan<- prometheus.Metric, duration float64) {
	ch <- prometheus.MustNewConstMetric(
		probeDurationDesc,
		prometheus.GaugeValue,
		duration,
	)
}

func sendGitlabHost(ch chan<- prometheus.Metric, host string) {
	ch <- prometheus.MustNewConstMetric(
		gitlabHostDesc,
		prometheus.GaugeValue,
		1,
		host,
	)
}
