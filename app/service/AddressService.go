package service

import (
	"app/core"
	"app/entity"
	"app/mappers"
	"app/pkg/geo"
	"app/pkg/geo/google"
	"app/repository"
)

type AddressService struct {
	*repository.AddressRepository
	geo.Coder
	*mappers.LocationMapper
	*AddressComponentService
	*ReverseCallStatisticService
}

func NewAddressService() *AddressService {
	c := &AddressService{}
	c.AddressRepository = repository.NewAddressRepository()
	c.Coder = google.NewGoogleCoder(core.App.Config["google_geo_key"])
	c.LocationMapper = mappers.NewLocationMapper()
	c.AddressComponentService = NewAddressComponentService()
	c.ReverseCallStatisticService = NewReverseStatisticService()
	return c
}

func (c *AddressService) Reverse(lat float64, lng float64) []entity.Address {

	addresses := c.FindOneByLanLng(lat, lng)

	if len(addresses) > 0 {
		return addresses
	}

	locations := c.Coder.Reverse(lat, lng)

	c.ReverseCallStatisticService.Add(lat, lng, locations)

	for _, location := range locations {

		address := c.LocationMapper.ToAddress(&location)

		for _, component := range location.Components {
			entity := c.LocationMapper.ToComponent(&component)
			c.AddressComponentService.upsert(&entity)
			address.Components = append(address.Components, entity)
		}

		c.AddressRepository.Save(&address)
	}

	return c.FindOneByLanLng(lat, lng)
}
