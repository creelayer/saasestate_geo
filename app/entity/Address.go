package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

type Address struct {
	Id int64 `json:"id"`
	PlaceId string `gorm:"type:varchar(150);index:idx_place_id,unique;" json:"-"`
	Name string `gorm:"type:varchar(255);" json:"name"`
	Lat float64 `gorm:"type:double precision;" json:"lat"`
	Lng float64 `gorm:"type:double precision;" json:"lng"`
	Data pgtype.JSONB `gorm:"type:jsonb;" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Components []AddressComponents `gorm:"many2many:address_to_components;" json:"components"`
}

//
//type Geometry4326 *geos.Geometry
//type Polygon4326 *geos.Coord