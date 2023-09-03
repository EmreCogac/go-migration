package main

import (
	"log"
	"net/http"
	"vox-api1/vox-api1/initializers"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {

	db := initializers.DB
	result := map[string]interface{}{}
	db.Table("users").Take(&result)

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.GET("/sql", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"param": result,
		})
	})

	log.Fatal(server.Run(":8080"))
}
