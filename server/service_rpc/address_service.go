package service_rpc

import (
	"context"
	"errors"
	"grpc_dev/server/db/dao"
	"grpc_dev/server/pb/address"
)

type AddressServer struct {
	address.UnimplementedAddressServer
}

func (s *AddressServer) GetCountries(c context.Context, in *address.GetCountriesRequest) (*address.GetCountriesReponse, error) {
	res, err := dao.GetAllCountries()
	if err != nil {
		return nil, err
	}
	var countries []*address.Country
	for _, c := range res {
		countries = append(countries, c.ToCountry())
	}
	return &address.GetCountriesReponse{
		Countries: countries,
	}, nil
}

func (s *AddressServer) GetCountryByCode(ctx context.Context, in *address.GetCountryRequest) (*address.GetCountryReponse, error) {
	if in.Code == nil {
		return nil, errors.New("code nil")
	}
	res, err := dao.GetByCode(*in.Code)
	if err != nil {
		return nil, err
	}
	return &address.GetCountryReponse{
		Country: res.ToCountry(),
	}, nil
}
