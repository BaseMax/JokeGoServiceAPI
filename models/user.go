package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"user"`
	Password string `gorm:"not null" json:"password"`
}
