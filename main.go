package main

import (
	"github.com/gofiber/fiber/v2"
	"shellrean.id/belajar-golang-rest-api/internal/config"
	"shellrean.id/belajar-golang-rest-api/internal/connection"
)

func main() {

	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	app.Get("/developers", developers)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
