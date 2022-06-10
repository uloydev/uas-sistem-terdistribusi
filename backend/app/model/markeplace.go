package model

type MarketplaceRequest struct {
	Name string `json:"name" example:"tokopedia"`
	Url  string `json:"url" example:"https://tokopedia.co.id"`
}

type MarketplaceResponse struct {
	BasicData
	Name string `json:"name"`
	Url  string `json:"url"`
}
