package dto

import "time"

type UserAccount struct {
	ID          int64     `gorm:"primaryKey, autoIncrement"`
	Username    string    `gorm:"unique"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Email       string    `gorm:"unique"`
	DateCreated time.Time `gorm:"autoCreateTime"`
}
