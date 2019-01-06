package wshandlers

import (
	"encoding/json"
	"net/http"
)

func marshalAndRespond(result interface{}, res http.ResponseWriter) {
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
