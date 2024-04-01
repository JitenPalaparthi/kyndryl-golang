package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	MAXRETRY     = 10
	MAXRETRYTIME = 10
)

func GetConnection(dsn string) (*gorm.DB, error) {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
retry:
	r := 1
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})

	if err != nil {
		if r <= MAXRETRY {
			time.Sleep(time.Second * MAXRETRYTIME)
			r++
			goto retry
		} else {
			return db, err
		}
	}
	return db, err
}
