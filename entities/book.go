package entities

type Book struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
	CateId    int     `json:"categoryId"`
	CateName  string  `json:"categoryName"`
}
