package repository

import (
	"app/core"
	"app/entity"
)

type AddressRepository struct {
	db *core.Gorm
}

func NewAddressRepository() *AddressRepository {
	c := &AddressRepository{}
	c.db = core.App.Gorm
	return c
}

func (c *AddressRepository) Save(entity *entity.Address) bool {
	c.db.Conn.Save(entity)
	return true
}

func (c *AddressRepository) FindOneByLanLng(lat float64, lng float64) []entity.Address {

	var e []entity.Address

	result := c.db.Conn.
		Raw("SELECT * FROM addresses WHERE ST_SetSRID(ST_MakePoint(lat,lng), 4326) <-> ST_SetSRID(ST_MakePoint(?, ?), 4326) < 0.00020 "+
			"ORDER BY ST_SetSRID(ST_MakePoint(lat,lng), 4326) <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)", lat, lng, lat, lng).
		Find(&e)

	if result.RowsAffected == 0 {
		return nil
	}

	return e
}

func (c *AddressRepository) FindOneById(id int64) *entity.Address {

	e := &entity.Address{}

	result := c.db.Conn.First(e, id)

	if result.RowsAffected == 0 {
		return nil
	}

	return e
}
