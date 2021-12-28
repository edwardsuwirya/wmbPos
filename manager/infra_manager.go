package manager

import (
	"github.com/edwardsuwirya/wmbPos/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Infra interface {
	SqlDb() *gorm.DB
	HttpClient() *http.Client
	Config() *config.Config
}

type infra struct {
	db     *gorm.DB
	config *config.Config
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Panicln(err)
	}
	return &infra{
		db:     resource,
		config: config,
	}
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func (i *infra) Config() *config.Config {
	return i.config
}
func (i *infra) HttpClient() *http.Client {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	return netClient
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
