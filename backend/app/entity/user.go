package entity

type User struct {
	BaseEntity
	Name     string
	Email    string
	Password string
	Birthday *int64
	Phone    *string
}
