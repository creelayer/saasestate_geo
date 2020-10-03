package service

import (
	"app/core"
	"app/entity"
	"app/pkg/geo"
	"app/pkg/geo/google"
	"app/repository"
	"encoding/json"
)

type AddressService struct {
	*repository.AddressRepository
	*repository.AddressComponentsRepository
	geo.Coder
}

func NewAddressService() *AddressService {
	c := &AddressService{}
	c.AddressRepository = repository.NewAddressRepository()
	c.AddressComponentsRepository = repository.NewAddressComponentsRepository()
	c.Coder = google.NewGoogleCoder(core.App.Config["google_geo_key"])
	return c
}

func (c *AddressService) Reverse(lat float64, lng float64) []geo.Response {

	addresses := c.FindOneByLanLng(lat, lng)

	if len(addresses) > 0 {
		return c.mapToLocations(&addresses)
	}

	locations := c.Coder.Reverse(lat, lng)

	for _, location := range locations {


		componentsIds := make([]int32, len(location.Components))

		for j, component:= range location.Components{
			e := c.mapToComponent(&component)
			c.AddressComponentsRepository.Save(&e)
			componentsIds[j] = e.Id
		}

		address := c.mapToAddress(&location)
		_ = address.Components.Set(componentsIds)
		c.AddressRepository.Save(&address)

	}

	return locations
}


func (c *AddressService) mapToComponent(cmp *geo.Component) entity.AddressComponents {

	component := entity.AddressComponents{
		LongName: cmp.LongName,
		ShortName:    cmp.ShortName,
	}

	_ = component.Types.Set(cmp.Types)

	return component

}


func (c *AddressService) mapToLocations(addresses *[]entity.Address) []geo.Response {

	locations := make([]geo.Response, len(*addresses))

	for i, address := range *addresses {
		locations[i] = c.mapToLocation(&address)
	}

	return locations
}

func (c *AddressService) mapToLocation(a *entity.Address) geo.Response {

	location := geo.Response{
		Id:      a.Id,
		PlaceId: a.PlaceId,
		Name:    a.Name,
		Lat:     a.Lat,
		Lng:     a.Lng,
	}

	location.Data, _ = json.Marshal(a.Data)

	return location

}

func (c *AddressService) mapToAddress(l *geo.Response) entity.Address {

	address := entity.Address{
		PlaceId: l.PlaceId,
		Name:    l.Name,
		Lat:     l.Lat,
		Lng:     l.Lng,
	}
	_ = address.Data.UnmarshalJSON(l.Data)

	return address

}
