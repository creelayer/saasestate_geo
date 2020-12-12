package service

import (
	"app/core"
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

func (c *AddressService) Reverse(lat float64, lng float64) []geo.Response {

	addresses := c.FindOneByLanLng(lat, lng)

	if len(addresses) > 0 {
		return c.LocationMapper.ToLocations(&addresses)
	}

	locations := c.Coder.Reverse(lat, lng)

	c.ReverseCallStatisticService.Add(lat, lng, locations)

	for _, location := range locations {

		componentsIds := make([]int32, len(location.Components))

		for j, component := range location.Components {
			entity := c.LocationMapper.ToComponent(&component)
			c.AddressComponentService.upsert(&entity)
			componentsIds[j] = entity.Id
		}

		address := c.LocationMapper.ToAddress(&location)
		_ = address.Components.Set(componentsIds)
		c.AddressRepository.Save(&address)
	}

	addresses = c.FindOneByLanLng(lat, lng)
	return c.LocationMapper.ToLocations(&addresses)
}
