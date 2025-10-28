package exporter

import (
	"fmt"

	"gitlab.com/gitlab-org/api/client-go"
)

type PipelineCount struct {
	Success            float64
	Failed             float64
	Pending            float64
	Created            float64
	WaitingForResource float64
	Preparing          float64
	Running            float64
	Cancelled          float64
	Skipped            float64
	Scheduled          float64
	Manual             float64
}

type Metrics struct {
	Count          PipelineCount
	LatestDuration float64
}

func GetMetrics(client *gitlab.Client, group, project string) (*Metrics, error) {
	pid := fmt.Sprintf("%s/%s", group, project)
	pipelines, _, err := client.Pipelines.ListProjectPipelines(pid, &gitlab.ListProjectPipelinesOptions{})

	if err != nil {
		return nil, err
	}

	latestPipeline, _, err := client.Pipelines.GetLatestPipeline(pid, &gitlab.GetLatestPipelineOptions{})

	if err != nil {
		return nil, err
	}

	var duration float64
	if latestPipeline != nil {
		duration = float64(latestPipeline.Duration)
	}

	pc := PipelineCount{}

	for _, pipe := range pipelines {
		switch pipe.Status {
		case "success":
			pc.Success++
		case "failed":
			pc.Failed++
		case "pending":
			pc.Pending++
		case "created":
			pc.Created++
		case "waiting_for_resource":
			pc.WaitingForResource++
		case "preparing":
			pc.Preparing++
		case "running":
			pc.Running++
		case "cancelled":
			pc.Cancelled++
		case "skipped":
			pc.Skipped++
		case "scheduled":
			pc.Scheduled++
		case "manual":
			pc.Manual++
		}
	}

	return &Metrics{
		Count:          pc,
		LatestDuration: duration,
	}, nil
}
