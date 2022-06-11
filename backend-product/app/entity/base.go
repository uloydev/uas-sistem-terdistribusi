package entity

type BaseEntity struct {
	ID        uint   `gorm:"primaryKey"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
	DeletedAt *int64 `gorm:"index"`
}
