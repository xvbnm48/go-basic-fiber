package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api:= app.Group("/api")
	api.Get("/endo", func(c *fiber.Ctx) error {
		go func() {
			time.Sleep(time.Second)
			// c.Send("Hello from endo")
			log.Print(c.Path())
		}()
		return c.SendString("sakura endo")
	})
	api.Get("/nako", func(c *fiber.Ctx) error {
		go func() {
			time.Sleep(time.Second)
			// c.Send("Hello from endo")
			log.Print(c.Path())
		}()
		return c.SendString("Yabuki Nako")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		go func() {
			time.Sleep(time.Second)
			log.Print(c.Path())
		}()
		return c.Send([]byte ("Hello, World!!"))
	})

	app.Get("/sakura", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Sakura!")
	})

	app.Listen(":3000")
}