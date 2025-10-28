package exporter

import (
	"fmt"
	"net/url"

	"gitlab.com/gitlab-org/api/client-go"
)

type PipelineCount struct {
	Success int32
	Failed  int32
	Pending int32
}

type Metrics struct {
	Count PipelineCount
}

func GetMetrics(client *gitlab.Client, group, project string) (*Metrics, error) {
	pid := url.PathEscape(fmt.Sprintf("%s/%s", group, project))
	pipelines, _, _ := client.Pipelines.ListProjectPipelines(pid, &gitlab.ListProjectPipelinesOptions{})

	var (
		successCount, failedCount, pendingCount int32
	)

	for _, pipe := range pipelines {
		if pipe.Status == "pending" {
			pendingCount++
		}

		if pipe.Status == "success" {
			successCount++
		}

		if pipe.Status == "failed" {
			failedCount++
		}
	}

	// here a real call to Gitlab CI/CD API

	// TODO: gitlab.com/gitlab-org/api/client-go

	return &Metrics{
		PipelineCount{
			Success: successCount,
			Failed:  failedCount,
			Pending: pendingCount,
		},
	}, nil
}
