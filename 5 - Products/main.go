package main

import (
	"fmt"
	"log"
	"products/data/database"
	"products/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-siris/siris/sessions/postgres"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file.")

		panic(err)
	}

	router.Use(gin.Recovery())

	database.InitializeDB()
	db := database.GetDB()

	routers.SetAPIRoute(db, router)

	defer db.Close()

	log.Fatal(router.Run(":8002"))
}
