package main

import (
	_"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app:= fiber.New() 

    app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string {
			"munashe": "madhuna",
			"admin": "munashe2001",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "munashe" && pass == "madhuna" {
				return true
			}
		if user == "admin" && pass == "munashe2001" {
			return true
		}
		return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
	         return c.Status(fiber.StatusBadRequest).SendString("error")
		},
		ContextUsername: "user",
		ContextPassword: "pass",
	}))

	

	log.Fatal(app.Listen(":3001"))
}
