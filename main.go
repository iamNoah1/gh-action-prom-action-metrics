package main

import (
	"log"
	"os"

	"github.com/iamNoah1/prometheus-action-metrics/internal"
)

var (
	prometheusRemoteHost     string
	prometheusRemoteUsername string
	prometheusRemotePassword string
)

func main() {
	prometheusRemoteHost = os.Getenv("INPUT_PROMETHEUS_REMOTE_WRITE_URL")
	prometheusRemoteUsername = os.Getenv("INPUT_PROMETHEUS_USERNAME")
	prometheusRemotePassword = os.Getenv("INPUT_PROMETHEUS_PASSWORD")

	//This is only relevant when running locally. When running as GH Action, those params are mandatory.
	if prometheusRemoteHost == "" || prometheusRemotePassword == "" || prometheusRemoteUsername == "" {
		log.Fatal("Invalid options, please provide host, username and password for Prometheus")
	}

	err := internal.WriteToPrometheus(prometheusRemoteHost, prometheusRemoteUsername, prometheusRemotePassword)

	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
