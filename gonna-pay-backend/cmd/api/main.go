// @title           Gonna Pay API
// @version         1.0
// @description     API de gestão financeira e divisão de custos
// @host            localhost:3333
// @BasePath        /
// @schemes         http

// @securityDefinitions.apikey  CookieAuth
// @in                          cookie
// @name                        access_token

package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/database"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/controller"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/service"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/rafabcanedo/gonna-pay/gonna-pay-backend/docs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	db, err := database.InitPostgres()
	if err != nil {
		log.Fatalf("error initializing postgres: %v", err)
	}
	defer db.Close()

	if err := database.RunMigrations(); err != nil {
		log.Fatalf("error running migrations: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db)
	contactRepo := repository.NewContactRepository(db)
	groupRepo := repository.NewGroupRepository(db)
	costRepo := repository.NewCostRepository(db)

	userSvc := service.NewUserService(userRepo)
	contactSvc := service.NewContactService(contactRepo)
	groupSvc := service.NewGroupService(groupRepo)
	costSvc := service.NewCostService(costRepo, groupRepo)

	userCtrl := controller.NewUserController(userSvc)
	authCtrl := controller.NewAuthController(userSvc, authRepo)
	contactCtrl := controller.NewContactController(contactSvc)
	groupCtrl := controller.NewGroupController(groupSvc)
	costCtrl := controller.NewCostController(costSvc)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	router.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "db_error", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	routes.AuthRoutes(router, authCtrl)
	routes.UserRoutes(router, userCtrl)
	routes.ContactRoutes(router, contactCtrl)
	routes.GroupRoutes(router, groupCtrl)
	routes.CostRoutes(router, costCtrl)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := ":3333"
	log.Printf("Server running on %s", port)

	if err := router.Run(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
