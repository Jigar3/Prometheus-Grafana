package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listOfCodes = []int{http.StatusOK, http.StatusNotFound, http.StatusForbidden}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/api/v1", getRoute)
	http.HandleFunc("/api/v2", getRoute)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getRoute(w http.ResponseWriter, r *http.Request) {
	statusCode := listOfCodes[rand.Intn(cap(listOfCodes))]

	statusCodeProcess.With(prometheus.Labels{"endpoint": r.URL.Path, "status_code": strconv.Itoa(statusCode)}).Inc()

	if statusCode == http.StatusOK {
		w.WriteHeader(statusCode)
		w.Write([]byte("200 - Status OK!"))
	}
	if statusCode == http.StatusNotFound {
		w.WriteHeader(statusCode)
		w.Write([]byte("404 - Status Not Found!"))
	}
	if statusCode == http.StatusForbidden {
		w.WriteHeader(statusCode)
		w.Write([]byte("403 - Status Forbidden!"))
	}

}

var (
	statusCodeProcess = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_status_codes",
		Help: "Gives status codes of particular route",
	},
		[]string{"endpoint", "status_code"},
	)
)
