package route

import (
	"log"
	"project-prim-aor/handler"
	"project-prim-aor/primdb"

	"github.com/gofiber/fiber/v2"
)

func init() {
	log.Println("init. routing....")
}

func InitRoute(app *fiber.App, db *primdb.DatabaseNaja) {
	// init handler
	h := handler.InitHandler(db)

	// init. route
	app.Get("/", h.HomeHandler)
	app.Get("/insert", h.InsertHandler)
	app.Get("/get", h.GetHandler)
}
