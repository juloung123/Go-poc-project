package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var users = []User{
	{ID: 1, FirstName: "John", LastName: "Doe"},
	{ID: 2, FirstName: "Jane", LastName: "Doe"},
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User Not Found"})
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Info"})
	}
	user.ID = len(users) + 1
	users = append(users, *user)
	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			updateUser := new(User)
			if err := c.BodyParser(updateUser); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Info"})
			}
			updateUser.ID = user.ID
			users[i] = *updateUser
			return c.JSON(updateUser)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User Not Found"})
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User Not Found"})
}

func main() {
	app := fiber.New()

	//Define routes
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	//Listen on port 8080
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}

}
