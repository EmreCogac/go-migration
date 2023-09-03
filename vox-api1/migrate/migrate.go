package main

import (
	"fmt"
	"log"
	"vox-api1/vox-api1/initializers"
	"vox-api1/vox-api1/models"
)

func init() {
	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
