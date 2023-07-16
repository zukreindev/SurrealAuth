package main

import (
	"FiberAuthWithSurrealDb/Database"
	"FiberAuthWithSurrealDb/Routers"
	"FiberAuthWithSurrealDb/Util"
	"github.com/gofiber/fiber/v2"
	"log"
)


func main() {
	app := fiber.New()

	app.Post("/login", routers.Login)
	app.Post("/register", routers.Register)

	database.Connect()
	log.Fatal(app.Listen("127.0.0.1:" + util.GetConfig("server", "port")))

}
