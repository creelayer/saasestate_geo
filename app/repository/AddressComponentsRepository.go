package repository

import (
	"app/core"
	"app/entity"
	"strings"
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

func (c *AddressComponentsRepository) Exist(component *entity.AddressComponents) bool {

	types := make([]string, len(component.Types.Elements))
	component.Types.AssignTo(&types)

	if err := c.db.Conn.Where("short_name = ? AND types @> ?", component.ShortName, "{"+strings.Join(types, ",")+"}").Take(component).Error; err != nil {
		return false
	}

	return true
}
