package main

import (
	"log"
	"net/http"

	"github.com/alecthomas/kingpin/v2"

	"github.com/AleksanderWWW/prometheus-gitlabci-exporter/exporter"
)

var (
	configFile = kingpin.Flag("config.file", "Exporter configuration file.").Default("example.yaml").String()
	listenAddr = kingpin.Flag("web.port", "Port to listen on.").Default(":9115").String()
)

func main() {
	kingpin.Parse()

	c, err := exporter.ParseConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	manager := exporter.ProbeManager{Config: *c}

	http.HandleFunc("/probe", manager.ProbeHandler)

	log.Printf("🚀 Listening on %s (GitLab host: %s)", *listenAddr, *c.Gitlab.Host)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
