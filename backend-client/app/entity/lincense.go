package entity

type License struct {
	BaseEntity
	Username string
	Key      string
	IsActive bool
}
