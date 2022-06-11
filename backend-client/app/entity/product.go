package entity

type Product struct {
	BaseEntity
	Title       string
	Price       int
	Description string
	Stock       int
	IsMaster    bool
}
