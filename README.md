# Go-Prom-HTTP

Simple GO lang application that expose metrics to prometheus and visualise data on Grafana.

## Requirements

- Golang
- Docker

## Getting started

### Starting the application

start all services by running the docker-compose file:

```
docker-compose up -d
```

### Generating load on HTTP Server

Used hey tool for load testing which helps to visualise data in a better way.

```
hey -z 5m -q 5 -m GET -H "Accept: text/html" http://localhost:9090

```

Application metrics are exported on http://localhost:9090/metrics.
End point for codes can be seen on http://localhost:9090/codes

### Importing Grafana dashboard

Grafana is running on http://localhost:3000/


1. Set Prometheus as your data source. http://host.docker.internal:9090
2. Import the grafana-dashboard.json as needed.
