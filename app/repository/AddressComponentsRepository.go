package repository

import (
	"app/core"
	"app/entity"
)

type AddressComponentsRepository struct {
	db *core.Gorm
}

func NewAddressComponentsRepository() *AddressComponentsRepository {
	c := &AddressComponentsRepository{}
	c.db = core.App.Gorm
	return c
}

func (c *AddressComponentsRepository) Save(entity *entity.AddressComponents) bool {
	c.db.Conn.Save(entity)
	return true
}
