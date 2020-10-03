package core

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

type Pgx struct {
	Conn       *pgx.Conn
	AutoSchema bool
}

func NewPgx(dsn string) *Pgx{

	c := &Pgx{}
	conn, err :=  pgx.Connect(context.Background(), dsn)

	if err != nil {
		log.Fatal(err)
	}

	c.Conn = conn
	c.AutoSchema = true

	return c
}

