package dto

import "grpc_dev/server/pb/address"

type Country struct {
	ID     int64   `gorm:"primaryKey, autoIncrement"`
	Phone  int64   `gorm:"not null"`
	Code   *string `gorm:"not null"`
	Name   *string `gorm:"not null"`
	Symbol *string `gorm:"not null"`
}

func (c *Country) ToCountry() *address.Country {
	return &address.Country{
		Code:   *c.Code,
		Name:   *c.Name,
		Phone:  int32(c.Phone),
		Symbol: *c.Symbol,
	}
}
