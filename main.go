package main

import (
	"fmt"

	"example.com/config"
	"example.com/db"
	"example.com/router"
	"example.com/validation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := config.New()
	db := db.New(config.Database)

	if db == nil {
		panic("Failed to connect to database")
	}

	validation.UseJsonKeyTagName()
	router.SetupRoutes(r, config, db)

	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}
