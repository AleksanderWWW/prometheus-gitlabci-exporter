package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AleksanderWWW/prometheus-gitlabci-exporter/cmd"
	"github.com/AleksanderWWW/prometheus-gitlabci-exporter/internal"
	"github.com/alecthomas/kingpin/v2"
	"gitlab.com/gitlab-org/api/client-go"
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

	manager := cmd.ProbeManager{Client: git, Sender: &internal.DefaultMetricsSender{}}

	http.HandleFunc("/probe", manager.ProbeHandler)

	log.Printf("🚀 Listening on %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
