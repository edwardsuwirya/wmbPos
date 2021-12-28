package api

import (
	"database/sql"
	"github.com/edwardsuwirya/wmbPos/config"
	"github.com/edwardsuwirya/wmbPos/delivery"
	"github.com/edwardsuwirya/wmbPos/entity"
	"github.com/edwardsuwirya/wmbPos/manager"
	"log"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
}

func NewApiServer() Server {
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManger(repo)
	return &server{
		config:  appConfig,
		infra:   infra,
		usecase: usecase,
	}
}

func (s *server) Run() {
	if !(s.config.RunMigration == "Y" || s.config.RunMigration == "y") {
		db, _ := s.infra.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(db)
		delivery.NewServer(s.config.RouterEngine, s.usecase)
		err := s.config.RouterEngine.Run(s.config.ApiBaseUrl)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		db := s.infra.SqlDb()
		err := db.AutoMigrate(&entity.CustomerOrder{}, &entity.CustomerOrderDetail{}, &entity.Payment{}, &entity.OrderPayment{})
		db.Unscoped().Where("id like ?", "%%").Delete(entity.Payment{})
		db.Model(&entity.Payment{}).Save([]entity.Payment{
			{
				ID:                "P01",
				PaymentMethodName: "Tunai",
			},
			{
				ID:                "P02",
				PaymentMethodName: "OPO",
			},
		})

		if err != nil {
			log.Fatalln(err)
		}
	}
}
