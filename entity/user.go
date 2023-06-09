package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"not null;unique;type:varchar(191)"`
	Products  []Product `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (u *User) TableName() string {
	return "tb_user"
}
