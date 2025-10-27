package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kingpin/v2"

	"github.com/AleksanderWWW/prometheus-gitlabci-exporter/exporter"
)

var (
	configFile = kingpin.Flag("config.file", "Exporter configuration file.").Default("example.yaml").String()
)

func main() {
	kingpin.Parse()

	c, err := exporter.ParseConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s:%d", *c.Gitlab.Host, *c.Gitlab.Port)
}
