package cmd

import (
	"fmt"

	"github.com/mcucu/mns-ecommerce/config"
	"github.com/mcucu/mns-ecommerce/internal/appcontexts"
	"github.com/mcucu/mns-ecommerce/internal/repositories"
	"github.com/mcucu/mns-ecommerce/internal/services"
	"github.com/mcucu/mns-ecommerce/internal/utils"
)

func Execute() {
	var err error
	cfg := config.NewConfig()
	app := appcontexts.NewAppContext(cfg)

	// var sqliteDB *sql.DB
	sqliteDB, err := app.GetSQLiteConn()
	if err != nil {
		fmt.Printf("failed to start, error connect to PostgreSQL | %v\n", err)
		return
	}

	opt := utils.Options{
		Config: cfg,
		DB:     sqliteDB,
	}

	repo := wiringRepository(repositories.Options{
		Options: opt,
	})

	service := wiringService(services.Options{
		Options:      opt,
		Repositories: repo,
	})

	server := NewServer(opt, service)

	// run app
	server.StartApp()
}

// wiring up all repositories
func wiringRepository(repoOption repositories.Options) *repositories.Repositories {
	var (
		productRepo = repositories.NewProduct(repoOption)
	)

	repo := repositories.Repositories{
		Product: productRepo,
	}

	return &repo
}

// wiring up all services
func wiringService(serviceOption services.Options) *services.Services {
	var (
		orderService   = services.NewOrderService(serviceOption)
		productService = services.NewProductService(serviceOption)
	)

	svc := services.Services{
		OrderService:   orderService,
		ProductService: productService,
	}

	return &svc
}
