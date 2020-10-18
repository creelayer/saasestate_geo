package service

import (
	"app/entity"
	"app/repository"
)

type AddressComponentService struct {
	*repository.AddressComponentsRepository
}

func NewAddressComponentService() *AddressComponentService {
	c := &AddressComponentService{}
	c.AddressComponentsRepository = repository.NewAddressComponentsRepository()
	return c
}

func (c *AddressComponentService) upsert(component *entity.AddressComponents) {

	if !c.AddressComponentsRepository.Exist(component) {
		c.AddressComponentsRepository.Save(component)
	}
}
