package entity

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
	"time"
)

type AddressComponents struct {
	Id int32 `json:"id"`
	Types pgtype.VarcharArray `gorm:"type:character varying[];" json:"-"`
	LongName string `json:"name"`
	ShortName string `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	TypeData []pgtype.Varchar `gorm:"-" sql:"-" json:"types"`
}

func (c *AddressComponents) AfterFind(tx *gorm.DB) (err error) {
	c.TypeData = c.Types.Elements
	return
}

