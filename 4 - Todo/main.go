package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	db "todo/data/db"
	"todo/routers/api"
)

var (
	router = gin.Default()
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")

		return
	}

	db.InitializeDB()

	if err != nil {
		panic(err)
	}

	api.SetRoute(db.GetDB(), router)

	router.Use(gin.Recovery())

	router.Run(":8002")
}
