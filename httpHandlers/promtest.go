package httpHandlers

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	adds = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "added_data",
			Help: "Counts added data",
		})
	removes = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "removed_data",
			Help: "Counts removed data",
		})
)

func addsOne() {
	adds.Inc()
}
func removesOne() {
	removes.Inc()
}
func init() {
	prometheus.MustRegister(adds)
	prometheus.MustRegister(removes)
}
