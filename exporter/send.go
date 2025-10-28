package exporter

import "github.com/prometheus/client_golang/prometheus"

func (c *GitLabCollector) sendPipelineCountByStatus(ch chan<- prometheus.Metric, metrics *Metrics) {
	counts := metrics.Count

	c.sendCount(ch, counts.Success, "success")

	c.sendCount(ch, counts.Failed, "failed")

	c.sendCount(ch, counts.Cancelled, "cancelled")

	c.sendCount(ch, counts.Manual, "manual")

	c.sendCount(ch, counts.Pending, "pending")

	c.sendCount(ch, counts.Preparing, "preparing")

	c.sendCount(ch, counts.Running, "running")

	c.sendCount(ch, counts.Scheduled, "scheduled")

	c.sendCount(ch, counts.Skipped, "skipped")

	c.sendCount(ch, counts.WaitingForResource, "waiting_for_resource")
}

func (c *GitLabCollector) sendCount(ch chan<- prometheus.Metric, val float64, status string) {
	ch <- prometheus.MustNewConstMetric(
		pipelineTotalDesc,
		prometheus.CounterValue,
		val,
		c.group,
		c.project,
		status,
	)
}

func (c *GitLabCollector) sendProbeSuccess(ch chan<- prometheus.Metric, success bool) {
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

func (c *GitLabCollector) sendLatestDuration(ch chan<- prometheus.Metric, duration float64) {
	ch <- prometheus.MustNewConstMetric(
		latestDurationDesc,
		prometheus.GaugeValue,
		duration,
	)
}
