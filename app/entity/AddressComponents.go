package entity

import (
	"time"
	"github.com/jackc/pgtype"
)

type AddressComponents struct {
	Id int32 `json:"id"`
	Types pgtype.VarcharArray `gorm:"type:character varying[];" json:"types"`
	LongName string `json:"long_name"`
	ShortName string `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

