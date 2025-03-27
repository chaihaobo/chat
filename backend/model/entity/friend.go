package entity

type Friend struct {
	BaseEntity
	UserID   uint64 `gorm:"column:user_id"`
	ToUserID uint64 `gorm:"column:to_user_id"`
}

func (Friend) TableName() string {
	return "friend"
}
