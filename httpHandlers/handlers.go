package httpHandlers

import (
	"log"
	"net/http"

	"github.com/BenjaminMcC/test_go_service/httpHandlers/httpUtils"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r.Method)
	switch r.Method {
	case http.MethodGet:
		List(w, r)
		break
	case http.MethodPost:
		Add(w, r)
		addsOne()
		break
	case http.MethodDelete:
		Remove(w, r)
		removesOne()
		break
	default:
		httpUtils.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

// routes pages
func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.Handle("/metrics", promhttp.Handler())
}
