package main

import "github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/core"

func main() {
	app := core.NewApp()
	app.Initialize().ConfigEndpoints().Run(":3201")
}
