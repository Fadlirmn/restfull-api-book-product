package models

type Product struct{
	ProductId int `json:"id"`
	NameProduct string `json:"name_product"`
	Item int `json:"item"`
	Type string `json:"type"`
} 