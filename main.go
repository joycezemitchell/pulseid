package main

import (
	"mitchvillanueva.com/pulseid/app"

)

// @title Pulseid API
// @version 1.0
// @description This is an Pulseid API Documentation.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main(){
	app.StartApplication()
}