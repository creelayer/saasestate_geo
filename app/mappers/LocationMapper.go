package mappers

import (
	"app/entity"
	"app/helpers"
	"app/pkg/geo"
	"encoding/json"
)

type LocationMapper struct {
	*helpers.TransliteratorHelper
}

func NewLocationMapper() *LocationMapper {
	c := &LocationMapper{}
	c.TransliteratorHelper = helpers.NewTransliteratorHelper("uk")
	return c
}

func (c *LocationMapper) ToComponent(cmp *geo.Component) entity.AddressComponents {

	component := entity.AddressComponents{
		LongName:  cmp.LongName,
		ShortName: c.TransliteratorHelper.Transliterate(cmp.ShortName),
	}

	_ = component.Types.Set(cmp.Types)

	return component

}

func (c *LocationMapper) ToLocations(addresses *[]entity.Address) []geo.Response {

	locations := make([]geo.Response, len(*addresses))

	for i, address := range *addresses {
		locations[i] = c.ToLocation(&address)
	}

	return locations
}

func (c *LocationMapper) ToLocation(a *entity.Address) geo.Response {

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

func (c *LocationMapper) ToAddress(l *geo.Response) entity.Address {

	address := entity.Address{
		PlaceId: l.PlaceId,
		Name:    l.Name,
		Lat:     l.Lat,
		Lng:     l.Lng,
	}
	_ = address.Data.UnmarshalJSON(l.Data)

	return address

}
