package martini

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
)

type WebService interface {
	CreateCate(martini.Params, http.ResponseWriter, *http.Request)
	CreateBook(martini.Params, http.ResponseWriter, *http.Request)
	ShowBooksByCate(martini.Params, http.ResponseWriter, *http.Request)
}

type MartiniWF struct {
	WebService WebService
}

func NewMartiniWF(ws WebService) *MartiniWF {
	return &MartiniWF{
		WebService: ws,
	}
}

func (mwf *MartiniWF) Run() {
	m := martini.Classic()

	m.Post("/categories", mwf.WebService.CreateCate)
	m.Post("/categories/:cateId/books", mwf.WebService.CreateBook)
	m.Get("/categories/:cateId/books", mwf.WebService.ShowBooksByCate)

	m.Run()
	fmt.Println("Listening on port 3000")
}
