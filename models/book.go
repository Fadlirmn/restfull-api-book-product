package models

type Book struct{
	IdBook int `json:"id_book"`
	NameBook string `json:"name_book"`
	Genre string `json:"genre"`
}