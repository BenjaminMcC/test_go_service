package httpHandlers

import (
	"net/http"

	"github.com/BenjaminMcC/test_go_service/httpHandlers/httpUtils"
	"github.com/BenjaminMcC/test_go_service/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	httpUtils.HandleSuccess(&w, storage.Get())
}
