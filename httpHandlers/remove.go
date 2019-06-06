package httpHandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/BenjaminMcC/test_go_service/httpHandlers/httpUtils"
	"github.com/BenjaminMcC/test_go_service/storage"
	"github.com/BenjaminMcC/test_go_service/structs"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	var id structs.ID

	err = json.Unmarshal(requestBody, &id)

	if err != nil {
		httpUtils.HandleError(&w, 500, "Bad Request", "ID not provided", nil)
		return
	}

	if id.ID == 0 {
		httpUtils.HandleError(&w, 500, "Bad Request", "ID not provided", nil)
		return
	}

	if storage.Remove(id.ID) {
		httpUtils.HandleSuccess(&w, structs.ID{ID: id.ID})
	} else {
		httpUtils.HandleError(&w, 400, "Bad Request", "ID not found", nil)
	}
}
