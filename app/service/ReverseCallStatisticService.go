package service

import (
	"app/entity"
	"app/pkg/geo"
	"app/repository"
)

type ReverseCallStatisticService struct {
	*repository.ReverseCallStatisticRepository
}

func NewReverseStatisticService() *ReverseCallStatisticService {
	c := &ReverseCallStatisticService{}
	c.ReverseCallStatisticRepository = repository.NewReverseStatisticRepository()
	return c
}

func (c *ReverseCallStatisticService) Add(lat float64, lng float64, response []geo.Response) {
	statistic := entity.ReverseCallStatistic{}
	statistic.Lat = lat
	statistic.Lng = lng
	c.ReverseCallStatisticRepository.Save(&statistic)

}
