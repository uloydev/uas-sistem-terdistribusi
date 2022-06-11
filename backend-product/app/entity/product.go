package entity

type Product struct {
	BaseEntity
	Title       string
	Price       string
	Description string
	Stock       string
	IsMaster    bool
}
