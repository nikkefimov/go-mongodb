package main

import (
	"os"

	"github.com/gofiber/fiber"
)

func main() {

	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {

	// create app
	app := fiber.New()

	// start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8081"
	}
	app.Listen(":" + port)

	return nil
}
