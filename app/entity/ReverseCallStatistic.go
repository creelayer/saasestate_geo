package entity

import (
	"time"
)

type ReverseCallStatistic struct {
	Id        int64
	Lat       float64 `gorm:"type:double precision;"`
	Lng       float64 `gorm:"type:double precision;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//
//type Geometry4326 *geos.Geometry
//type Polygon4326 *geos.Coord
