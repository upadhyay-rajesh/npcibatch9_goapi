package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(" Secure Go API over HTTPS!")
	})

	certFile := "./cert.pem"
	keyFile := "./key.pem"

	log.Println("Server running at https://localhost:8443")
	if err := app.ListenTLS(":8443", certFile, keyFile); err != nil {
		log.Fatal(err)
	}
}
