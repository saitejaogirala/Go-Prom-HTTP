package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var codeCounts *prometheus.CounterVec

func handleCodes(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Code int `json:"response"`
	}
	json.NewDecoder(r.Body).Decode(&request)

	codeCounts.With(prometheus.Labels{"response": fmt.Sprintf("%d", request.Code)}).Inc()
	fmt.Fprintf(w, "Code %d count: %v", request.Code, codeCounts)
}

func main() {
	codeCounts = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "response_count",
		Help: "Counts of response codes received",
	}, []string{"response"})
	prometheus.MustRegister(codeCounts)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/codes", handleCodes)
	http.ListenAndServe(":9000", nil)
}
