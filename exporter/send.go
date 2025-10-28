package exporter

import "github.com/prometheus/client_golang/prometheus"

func (c *GitLabCollector) sendPipelineCountByStatus(ch chan<- prometheus.Metric, metrics *Metrics) {
	// send successful pipeline count
	ch <- prometheus.MustNewConstMetric(
		pipelineTotalDesc,
		prometheus.CounterValue,
		float64(metrics.Count.Success),
		c.group,
		c.project,
		"success",
	)

	// send failed pipeline count
	ch <- prometheus.MustNewConstMetric(
		pipelineTotalDesc,
		prometheus.CounterValue,
		float64(metrics.Count.Failed),
		c.group,
		c.project,
		"failed",
	)

	// send pending count
	ch <- prometheus.MustNewConstMetric(
		pipelineTotalDesc,
		prometheus.CounterValue,
		float64(metrics.Count.Pending),
		c.group,
		c.project,
		"pending",
	)
}
