package wshandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-martini/martini"
)

type CateReqInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (ws *WebService) CreateCate(
	params martini.Params,
	res http.ResponseWriter,
	req *http.Request,
) {
	// Read body
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshal
	var cateReqInfo CateReqInfo
	err = json.Unmarshal(body, &cateReqInfo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	createdCate, err := ws.BookUsecase.CreateCate(
		cateReqInfo.Id,
		cateReqInfo.Name,
	)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	marshalAndRespond(createdCate, res)
}
