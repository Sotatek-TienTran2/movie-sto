package model

type Movie struct {
	IdMovie      int    `json:"_idMovie,omitempty"`
	Name         string `json:"Name,omitempty"`
	NameOfAuthor string `json:"NameOfAuthor,omitempty"`
	Category     string `json:"Category,omitempty"`
}
