package main

import (
	"fmt"
	"log"

	"github.com/AleksanderWWW/prometheus-gitlabci-exporter/exporter"
)


func main() {
	c, err := exporter.ParseConfig("example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.Gitlab.Host)
}
