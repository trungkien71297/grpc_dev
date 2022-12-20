package dao

import (
	"grpc_dev/server/db"
	"grpc_dev/server/db/dto"
)

func GetAllCountries() ([]*dto.Country, error) {
	var countries []*dto.Country
	result := db.Client.Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}

func GetByCode(code string) (*dto.Country, error) {
	var country *dto.Country
	result := db.Client.Where("code = ?", code).First(&country)
	if result.Error != nil {
		return nil, result.Error
	}
	return country, nil
}
