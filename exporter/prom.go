package exporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

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

type GitLabCollector struct {
	group, project string
	client         *gitlab.Client
}

func (c *GitLabCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pipelineTotalDesc
	ch <- successDesc
}

func (c *GitLabCollector) Collect(ch chan<- prometheus.Metric) {
	metrics, err := GetMetrics(c.client, c.group, c.project)

	if err != nil {
		log.Printf("ERROR %s", err)
		c.sendProbeSuccess(ch, false)  // send.go
		return
	}

	c.sendPipelineCountByStatus(ch, metrics)  // send.go
	c.sendProbeSuccess(ch, true)  // send.go
}

type ProbeManager struct {
	Client *gitlab.Client
}

func (pm *ProbeManager) ProbeHandler(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	project := r.URL.Query().Get("project")

	if group == "" || project == "" {
		http.Error(w, "missing 'group' or 'project' parameter", http.StatusBadRequest)
		return
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(&GitLabCollector{group: group, project: project, client: pm.Client})

	log.Printf("Scraping GitLab pipelines for %s/%s", group, project)
	promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}
