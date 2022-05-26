package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"mitchvillanueva.com/pulseid/datasources/mysql/schema"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Create schema
	schema.CreateSchema()
	
	// Map url
	mapUrls()

	log.Println("about to start the application...")
	router.Run(":7070")
}

