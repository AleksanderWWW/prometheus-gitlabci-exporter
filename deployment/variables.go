package main

import (
	"log"
	"os"
)

const EXPORTER_IMAGE string = "ghcr.io/aleksanderwww/prometheus-gitlabci-exporter"
const EXPORTER_TAG string = "0.2.0"
const EXPORTER_PORT int = 9115

var EXPORTER_API_TOKEN string

const PROMETHEUS_IMAGE string = "prom/prometheus"
const PROMETHEUS_TAG string = "v3.8.1"
const PROMETHEUS_PORT int = 9090

const DOCKER_NETWORK_NAME string = "prom-net"

func init() {
	apiToken, ok := os.LookupEnv("GITLAB_API_TOKEN")
	if !ok {
		log.Fatalf("GITLAB_API_TOKEN not set")
	}

	EXPORTER_API_TOKEN = apiToken
}
