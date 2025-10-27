package exporter

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Metrics struct {
	PipelineSuccessCount int32
	PipelineFailuedCount int32
}

func GetMetrics(config exporterConfig, group, project string) (*Metrics, error) {
	projectId := url.PathEscape(fmt.Sprintf("%s/%s", group, project))

	apiURL := fmt.Sprintf("%s/api/v4/projects/%s/pipelines", *config.Gitlab.Host, projectId)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Add Authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *config.Gitlab.ApiToken))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK response: %s\n%s", resp.Status, body)
	}


	// here a real call to Gitlab CI/CD API

	// TODO: gitlab.com/gitlab-org/api/client-go

	return &Metrics{
		PipelineSuccessCount: 42,
		PipelineFailuedCount: 1,
	}, nil
}
