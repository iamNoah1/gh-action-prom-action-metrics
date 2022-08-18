# Prometheus GitHub Action Metrics Writer

## Run locally 
* `export INPUT_PROMETHEUS_REMOTE_WRITE_URL=https://prometheus-prod-01-eu-west-0.grafana.net/api/prom/push`
* `export INPUT_PROMETHEUS_USERNAME=<username>`
* `export INPUT_PROMETHEUS_PASSWORD=<pwd>`
* `go run main.go`