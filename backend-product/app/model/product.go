package model

type ProductResponse struct {
	BasicData
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	IsMaster    bool   `json:"is_master"`
}

type ProductRequest struct {
	Title       string `json:"title" example:"Macbook Air M1"`
	Price       int    `json:"price" example:"15000000"`
	Description string `json:"description" example:"Macbook Air with apple silicon M1 chip"`
	Stock       int    `json:"stock" example:"29"`
}
