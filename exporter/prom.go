package exporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type GitLabCollector struct {
	client        *gitlab.Client
	metricsSender MetricsSender
	opts          *GitlabScrapeOpts
}

type GitlabScrapeOpts struct {
	group, project string
}

type MetricsSender interface {
	SendProbeFailure(ch chan<- prometheus.Metric)
	SendMetrics(ch chan<- prometheus.Metric, metrics *Metrics, opts *GitlabScrapeOpts)
}

func (c *GitLabCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pipelineTotalDesc
	ch <- successDesc
	ch <- latestDurationDesc
	ch <- probeDurationDesc
}

func (c *GitLabCollector) Collect(ch chan<- prometheus.Metric) {
	metrics, err := GetMetrics(c.client, c.opts)

	if err != nil {
		log.Printf("ERROR %s", err)
		c.metricsSender.SendProbeFailure(ch)
		return
	}

	c.metricsSender.SendMetrics(ch, metrics, c.opts)
}

type ProbeManager struct {
	Client *gitlab.Client
	Sender MetricsSender
}

func (pm *ProbeManager) ProbeHandler(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	project := r.URL.Query().Get("project")

	if group == "" || project == "" {
		http.Error(w, "missing 'group' or 'project' parameter", http.StatusBadRequest)
		return
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(&GitLabCollector{
		client:        pm.Client,
		metricsSender: pm.Sender,
		opts: &GitlabScrapeOpts{
			project: project,
			group:   group,
		},
	})

	log.Printf("Scraping GitLab pipelines for %s/%s", group, project)
	promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}
