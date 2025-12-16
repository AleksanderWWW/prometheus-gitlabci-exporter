package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var client *http.Client = &http.Client{}

func TestAPI(t *testing.T) {
	baseUrl := os.Getenv("TEST_BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:9115/probe"
	}

	type testCase struct {
		group         string
		project       string
		shouldSucceed bool
	}

	tests := []testCase{
		{
			group:         "alwojnarowicz",
			project:       "some-project1",
			shouldSucceed: true,
		},
		{
			group:         "alwojnarowicz",
			project:       "non-existent",
			shouldSucceed: false,
		},
	}

	for _, tCase := range tests {
		t.Run(fmt.Sprintf("%s/%s", tCase.group, tCase.project), func(t *testing.T) {
			url := fmt.Sprintf("%s?group=%s&project=%s", baseUrl, tCase.group, tCase.project)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Fatalf("creating request: %v", err)
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("making request: %v", err)
			}
			defer saveCloseRespBody(resp)

			bodyText, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("reading response body: %v", err)
			}

			succeeded := strings.Contains(string(bodyText), "gitlab_probe_success 1")

			if succeeded != tCase.shouldSucceed {
				t.Errorf("test failed for %s/%s", tCase.group, tCase.project)
			}
		})
	}
}

func TestResponseContent(t *testing.T) {
	baseUrl := os.Getenv("TEST_BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:9115/probe"
	}

	url := fmt.Sprintf("%s?group=%s&project=%s", baseUrl, "alwojnarowicz", "some-project1")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("creating request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("making request: %v", err)
	}
	defer saveCloseRespBody(resp)

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response body: %v", err)
	}

	assert.Contains(t, string(bodyText), "gitlab_probe_success")
	assert.Contains(t, string(bodyText), "gitlab_host")
	assert.Contains(t, string(bodyText), "gitlab_pipeline_total")
	assert.Contains(t, string(bodyText), "gitlab_pipeline_total")
	assert.Contains(t, string(bodyText), "exporter_probe_duration_seconds")
}

func saveCloseRespBody(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		panic(err)
	}
}
