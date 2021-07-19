// main.go

package main

import (
	"os"

	"github.com/neel-singhania/iitk-coin/controllers"

	"log"

	"github.com/neel-singhania/iitk-coin/models"

	"github.com/neel-singhania/iitk-coin/middlewares"

	"github.com/neel-singhania/iitk-coin/database"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	route := gin.Default()

	route.GET("/check", func(context *gin.Context) {
		context.String(200, "good to go")
	})

	route.POST("/login", controllers.Login)
	route.POST("/signup", controllers.Signup)
	route.POST("/init", middlewares.Authz_Admin(), controllers.Account_init)
	route.GET("/balance", middlewares.Authz(), controllers.GetBalance)
	route.POST("/transfer", middlewares.Authz(), controllers.Transfer)
	route.GET("/secretpage", middlewares.Authz(), controllers.Profile)

	// api_file := route.Group("/secretpage")
	// {
	// 	protected_route := api_file.Group("/").Use(middlewares.Authz())
	// 	{
	// 		protected_route.GET("/", controllers.Profile)
	// 	}
	// }

	return route
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	err := database.InitDatabase(dsn)
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{})

	dsn = os.Getenv("ACCOUNT_DATABASE_URL")
	err = database.InitDatabaseAcc(dsn)
	if err != nil {
		log.Fatalln("could not create Acc ", err)
	}
	database.GlobalDBAcc.AutoMigrate(&models.Account{})

	dsn = os.Getenv("TRANSACTION_DATABASE_URL")
	err = database.InitDatabaseTrans(dsn)
	if err != nil {
		log.Fatalln("could not create Transfer database.... ", err)
	}
	database.GlobalDBTrans.AutoMigrate(&models.Transaction{})

	route := setupRouter()
	route.Run(":8080")
}
