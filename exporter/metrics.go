package exporter

import (
	"fmt"

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
	pid := fmt.Sprintf("%s/%s", group, project)
	pipelines, _, err := client.Pipelines.ListProjectPipelines(pid, &gitlab.ListProjectPipelinesOptions{})

	if err != nil {
		return nil, err
	}

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

	return &Metrics{
		PipelineCount{
			Success: successCount,
			Failed:  failedCount,
			Pending: pendingCount,
		},
	}, nil
}
