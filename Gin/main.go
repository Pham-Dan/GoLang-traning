package main

import (
	"log"
	"main/api/route"
	bootstrap "main/bootrap"
	"os"
)
var App bootstrap.Application
func main() {
	App = bootstrap.App()
	env := App.Env
	// db := app.DB
	os.Setenv("jwtSecret", "123")
	router := route.SetUp()

	log.Printf("Serve at: http://0.0.0.0%s", env.ServerAddress)
	router.Run(env.ServerAddress)

}
