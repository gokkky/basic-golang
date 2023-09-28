package main

import (
	"log"

	"project-prim-aor/route"

	"project-prim-aor/primdb"

	"github.com/gofiber/fiber/v2"
)

// clean architecture
// hexagonal architecture : port adaptor
func main() {
	// init db
	db := primdb.NewDatabase() // return session database

	app := fiber.New()

	// init router
	route.InitRoute(app, db)

	// run app with port :8000
	log.Fatal(app.Listen(":8000"))
}
