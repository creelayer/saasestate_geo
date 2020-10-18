package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

type Address struct {
	Id int64
	PlaceId string `gorm:"type:varchar(150);index:idx_place_id,unique;"`
	Name string `gorm:"type:varchar(255);"`
	Lat float64 `gorm:"type:double precision;"`
	Lng float64 `gorm:"type:double precision;"`
	Components pgtype.Int4Array `gorm:"type:int[];"`
	Data pgtype.JSONB `gorm:"type:jsonb;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//
//type Geometry4326 *geos.Geometry
//type Polygon4326 *geos.Coord