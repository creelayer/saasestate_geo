package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

type AddressComponents struct {
	Id int32
	Types pgtype.VarcharArray `gorm:"type:character varying[];"`
	LongName string
	ShortName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

