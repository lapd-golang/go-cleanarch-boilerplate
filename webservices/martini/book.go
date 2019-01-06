package wshandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

type BookReqInfo struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (ws *WebService) ShowBooksByCate(
	params martini.Params,
	res http.ResponseWriter,
	req *http.Request,
) {
	cateId, err := strconv.Atoi(params["cateId"])
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	books, err := ws.BookUsecase.GetBooksByCate(cateId)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	marshalAndRespond(books, res)
}

func (ws *WebService) CreateBook(
	params martini.Params,
	res http.ResponseWriter,
	req *http.Request,
) {
	cateId, err := strconv.Atoi(params["cateId"])
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read body
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshal
	var bookReqInfo BookReqInfo
	err = json.Unmarshal(body, &bookReqInfo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	createdBook, err := ws.BookUsecase.CreateBook(
		bookReqInfo.Id,
		bookReqInfo.Name,
		bookReqInfo.Price,
		cateId,
	)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	marshalAndRespond(createdBook, res)
}
