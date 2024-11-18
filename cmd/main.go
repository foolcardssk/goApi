package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("API Hit")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to laod .env file")
	}

	port := os.Getenv("PORT")

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", homeHandler)

	fmt.Println("API Server listening on ", port)
	app.Listen(port)
}
