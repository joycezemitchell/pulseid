package app

import (
	"mitchvillanueva.com/pulseid/controllers/client"
	"mitchvillanueva.com/pulseid/controllers/ping"
	"mitchvillanueva.com/pulseid/controllers/admin"
	"mitchvillanueva.com/pulseid/controllers/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "mitchvillanueva.com/pulseid/docs"
)


func mapUrls() {

	url := ginSwagger.URL("http://localhost:7070/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))


	router.GET("/ping", ping.Ping)
	router.POST("/login", admin.Login)
	router.GET("/validateToken", client.ValidateToken)

	router.Use(middleware.Middleware)
	router.GET("/requesttoken", admin.RequestToken)
	router.GET("/getalltokens", client.GetAllTokens)
}

