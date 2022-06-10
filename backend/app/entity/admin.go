package entity

type Admin struct {
	BaseEntity
	Name     string
	Email    *string
	Password *string
}
