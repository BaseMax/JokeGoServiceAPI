package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"user"`
	Password string `gorm:"not null" json:"password"`
}

func RegisterUser(u *User) error {
	return nil
}

func LoginUser(u *User) error {
	return nil
}
