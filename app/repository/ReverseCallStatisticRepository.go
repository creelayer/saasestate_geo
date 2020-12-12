package repository

import (
	"app/core"
	"app/entity"
)

type ReverseCallStatisticRepository struct {
	db *core.Gorm
}

func NewReverseStatisticRepository() *ReverseCallStatisticRepository {
	c := &ReverseCallStatisticRepository{}
	c.db = core.App.Gorm
	return c
}

func (c *ReverseCallStatisticRepository) Save(entity *entity.ReverseCallStatistic) bool {
	c.db.Conn.Save(entity)
	return true
}
