package config

import (
	"database/sql"
	"fmt"
	"github.com/edwardsuwirya/wmbPos/delivery"
	"github.com/edwardsuwirya/wmbPos/entity"
	"github.com/edwardsuwirya/wmbPos/manager"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	RouterEngine   *gin.Engine
	ApiBaseUrl     string
	runMigration   string
}

func NewConfig() *Config {
	runMigration := os.Getenv("DB_MIGRATION")
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	infraManager := manager.NewInfra(dsn)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManger(repoManager)

	config := new(Config)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager

	r := gin.Default()
	delivery.NewServer(r, useCaseManager)
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	config.runMigration = runMigration

	return config
}

func (c *Config) RunMigration() {
	if c.runMigration == "Y" || c.runMigration == "y" {
		db := c.InfraManager.SqlDb()
		err := db.AutoMigrate(&entity.CustomerOrder{}, &entity.CustomerOrderDetail{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (c *Config) StartEngine() {
	if !(c.runMigration == "Y" || c.runMigration == "y") {
		db, _ := c.InfraManager.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(db)
		err := c.RouterEngine.Run(c.ApiBaseUrl)
		if err != nil {
			log.Fatal(err)
		}
	}
}
