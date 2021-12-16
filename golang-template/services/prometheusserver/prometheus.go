package prometheusserver

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(port int) {
	http.HandleFunc("/health", (func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK!"))
	}))
	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("Prometheus server listening to port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
