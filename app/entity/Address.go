package entity

import (
	"github.com/jackc/pgtype"
	"github.com/paulsmith/gogeos/geos"
	"time"
)

type Address struct {
	Id int64
	PlaceId string `gorm:"type:varchar(150);index:idx_place_id,unique;"`
	Name string `gorm:"type:varchar(255);"`
	Lat float64 `gorm:"type:double precision;"`
	Lng float64 `gorm:"type:double precision;"`
	NLat float64 `gorm:"type:double precision;"`
	NLng float64 `gorm:"type:double precision;"`
	SLat float64 `gorm:"type:double precision;"`
	SLng float64 `gorm:"type:double precision;"`
	Pp float64 `gorm:"type:geography(POINT,4269);"`
	Components pgtype.Int4Array `gorm:"type:int[];"`
	Data pgtype.JSONB `gorm:"type:jsonb;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


type Geometry4326 *geos.Geometry
type Polygon4326 *geos.Coord