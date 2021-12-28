package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Config struct {
	RouterEngine          *gin.Engine
	ApiBaseUrl            string
	RunMigration          string
	DataSourceName        string
	TableManagementConfig TableManagementConfig
	OpoPaymentConfig      OpoPaymentConfig
}

type TableManagementConfig struct {
	ApiBaseUrl string
}

type OpoPaymentConfig struct {
	ApiBaseUrl      string
	ClientSecretKey string
}

func NewConfig() *Config {
	config := new(Config)
	runMigration := os.Getenv("DB_MIGRATION")
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	//baseURL = "http://localhost:8888/api/table"
	tableManagementBaseUrl := os.Getenv("TABLE_API")
	tableManagementConfig := TableManagementConfig{ApiBaseUrl: tableManagementBaseUrl}

	//opoBaseURL = "http://159.223.42.164:8899/opo/payment"
	//key = "E157934D-EA2E-49F6-9DCE-398B750BE4F0"
	opoPaymentBaseUrl := os.Getenv("OPO_API")
	opoSecretKey := os.Getenv("OPO_KEY")
	opoPaymentConfig := OpoPaymentConfig{
		ApiBaseUrl:      opoPaymentBaseUrl,
		ClientSecretKey: opoSecretKey,
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	config.DataSourceName = dsn
	config.TableManagementConfig = tableManagementConfig
	config.OpoPaymentConfig = opoPaymentConfig

	r := gin.Default()
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	config.RunMigration = runMigration

	return config
}
