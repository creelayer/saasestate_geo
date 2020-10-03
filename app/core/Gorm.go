package core

import (
	"app/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Gorm struct {
	Conn              *gorm.DB
	dsn string
	AutoSchema        bool
	connectRetryCount uint8
}

func NewGorm(dsn string) *Gorm {

	c := &Gorm{}
	c.dsn = dsn
	c.connectRetryCount = 30
	c.AutoSchema = true

	c.connect()

	return c
}

func (c *Gorm) connect()  {

	for {
		conn, err := gorm.Open(postgres.Open(c.dsn), &gorm.Config{})
		if err != nil {
			if c.connectRetryCount == 0 {
				log.Fatalf("Not able to establish connection to database %s", c.dsn)
			}

			log.Printf(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", c.connectRetryCount))
			c.connectRetryCount--
			time.Sleep(2 * time.Second)
		}else{
			c.Conn = conn
			break
		}
	}

	if c.AutoSchema {
		c.generateSchema()
	}
}

func (c *Gorm) generateSchema() {
	_ = c.Conn.Debug().AutoMigrate(&entity.Address{})
	_ = c.Conn.Debug().AutoMigrate(&entity.AddressComponents{})
}
