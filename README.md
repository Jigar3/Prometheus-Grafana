
# Prometheus Monitoring Server With Grafana Dashboard 

**Prometheus** is metrics-based monitoring & alerting stack made for dynamic cloud environments. 

**Grafana** is open-source visualization and analytics software. It allows you to query, visualize, alert on, and explore your metrics no matter where they are stored.

You need to install the below-mentioned tools to be able to run this program.
1. [Grafana](https://grafana.com/docs/grafana/latest/installation/)
2. [Prometheus](https://prometheus.io/docs/prometheus/latest/installation/)
3. [Vegeta](https://github.com/tsenart/vegeta)

To install dependencies, do
```
go mod tidy
```

To run the server, do
```
go run main.go	# start the API server
prometheus --config.file=./prometheus.yml	# start the prometheus server
grafana-server	# start the grafana server

# To create artificial traffic
echo "GET http://localhost:9000/api/v2" | vegeta attack -duration=600s | tee results.bin | vegeta report


```

More information can be accessed via a blog written for the same.
https://medium.com/@chavada.jigar.3/configure-a-prometheus-monitoring-server-with-a-grafana-dashboard-59faa4a42ef6