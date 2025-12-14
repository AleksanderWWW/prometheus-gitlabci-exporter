package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"gitlab.com/gitlab-org/api/client-go"

	"github.com/AleksanderWWW/prometheus-gitlabci-exporter/exporter"
)

var (
	apiToken   = kingpin.Flag("api.token", "Gitlab API token.").String()
	listenAddr = kingpin.Flag("web.port", "Port to listen on.").Default(":9115").String()
)

func main() {
	kingpin.Parse()

	if apiToken == nil {
		token := os.Getenv("GITLAB_API_TOKEN")
		if token == "" {
			log.Fatal("no Gitlab API token provided")
		}
		apiToken = &token
	}

	git, err := gitlab.NewClient(*apiToken)
	if err != nil {
		log.Fatalf("Could not create client: %s", err)
	}

	manager := exporter.ProbeManager{Client: git, Sender: &exporter.DefaultMetricsSender{}}

	http.HandleFunc("/probe", manager.ProbeHandler)

	log.Printf("🚀 Listening on %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
