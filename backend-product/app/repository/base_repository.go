package repository

type BaseRepository[T any] interface {
	Insert(entity T) T

	FindAll() []T

	FindById(id uint) T
}
