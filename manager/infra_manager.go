package manager

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func NewInfra(dataSourceName string) Infra {
	resource, err := initDbResource(dataSourceName)
	if err != nil {
		log.Panicln(err)
	}
	return &infra{
		db: resource,
	}
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
