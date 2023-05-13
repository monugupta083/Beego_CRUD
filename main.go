package main

import (
	"myproject/config"
	routers "myproject/router"
)

func main() {
	// Connect to the database
	config.Connect()

	// Start the router
	r := routers.SetupRouter()
	r.Run()
}
