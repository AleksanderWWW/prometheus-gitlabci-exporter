package exporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pipelineTotalDesc = prometheus.NewDesc(
		"gitlab_pipeline_total",
		"Total number of pipelines for the target project.",
		[]string{"group", "project", "status"},
		nil,
	)

type GitLabCollector struct {
	group, project string
	config exporterConfig
}

func (c *GitLabCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pipelineTotalDesc
}

func (c *GitLabCollector) Collect(ch chan<- prometheus.Metric) {
	metrics, _ := GetMetrics(c.config, c.group, c.project)

	// send successful pipeline count
	ch <- prometheus.MustNewConstMetric(
			pipelineTotalDesc, 
			prometheus.CounterValue, 
			float64(metrics.PipelineSuccessCount), 
			c.group, 
			c.project, 
			"success",
		)

	// send failed pipeline count
	ch <- prometheus.MustNewConstMetric(
			pipelineTotalDesc, 
			prometheus.CounterValue, 
			float64(metrics.PipelineFailuedCount), 
			c.group, 
			c.project, 
			"failed",
		)
}

type ProbeManager struct {
	Config exporterConfig
}

func (pm *ProbeManager) ProbeHandler(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	project := r.URL.Query().Get("project")

	if group == "" || project == "" {
		http.Error(w, "missing 'group' or 'project' parameter", http.StatusBadRequest)
		return
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(&GitLabCollector{group: group, project: project, config: pm.Config})

	log.Printf("Scraping GitLab pipelines for %s/%s", group, project)
	promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}
