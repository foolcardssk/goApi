package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to laod .env file")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("GET / call made from : ", c.IP())
		return c.SendString("Hello, World!")
	})

	fmt.Println("API Server listening on ", port)
	app.Listen(port)
}
