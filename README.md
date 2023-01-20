# Go-Prom-HTTP

Simple GO lang application that expose metrics to prometheus and visualise data on Grafana.

## Requirements

- Golang
- Docker
- kind
- MetalLB load balancer. Used this: https://kind.sigs.k8s.io/docs/user/loadbalancer/

## Quick Intro:

Deployment of the application can be done in various ways using helm charts, writing a simple make file and deploying prom operator directly on the namespace. But, i have used bare metal ways(deploying in a hardway) like provisioning every single service,deployment,roles,RBAC,rolebindings,clusterrolebindings writing all the configuration files in their respective loacations for example all the promethues configuration can be found under prometheus directory.

## Getting started

I have used custom kind configuration for the cluster to bring the kubernetes cluster components targets up.

### Starting the application 

Application can be deployed using both docker-compose and kubernetes native way. 

1. start all services by running the docker-compose file:

```
docker-compose up -d
```
2. Deploy using the kubernetes deploy configuration. This can be found here: https://github.com/saitejaogirala/Go-Prom-HTTP/tree/main/deployment


### Generating load on HTTP Server

Used hey tool for load testing which helps to visualise data in a better way.

```
hey -z 5m -q 5 -m GET -H "Accept: text/html" http://localhost:9090

```

Application metrics are exported on http://localhost:9090/metrics.
End point for codes can be seen on http://localhost:9090/codes

### Importing Grafana dashboard

Grafana is running on http://localhost:3000/


1. Set Prometheus as your data source. http://host.docker.internal:9090 If this doesn't work please set the prometheus service IP as data source.
2. Import the grafana-dashboard.json as needed.
