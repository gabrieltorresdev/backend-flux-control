package main

import (
	"fmt"
	"log"

	"github.com/gabrieltorresdev/backend-flux-control/internal/application/service"
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/config"
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/http/v1/rest/gin/controller"
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/http/v1/rest/gin/routes"
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/persistence/gorm/db"
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/persistence/gorm/repository"
	"github.com/gabrieltorresdev/backend-flux-control/pkg/hateoas"
	"github.com/gin-gonic/gin"
)

func init() {
	hateoas.GlobalInstance.Setup("/v1")

	hateoas.GlobalInstance.RegisterResource("transaction", hateoas.ResourceConfig{
		ResourceName:     "transactions",
		DefaultLinkTypes: []string{"self", "collection", "create", "show", "update", "delete"},
		CustomLinks:      map[string]string{},
		PaginationLinks:  []string{"self", "collection"},
	})
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// db := db.NewGormDB(config)
	db, err := db.NewGormDBWithAutoMigrate(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	router := gin.Default()

	routes.SetupRoutes(router, transactionController)

	router.Run(fmt.Sprintf(":%d", config.GetInt("server.port")))
}
