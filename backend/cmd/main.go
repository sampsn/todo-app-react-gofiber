package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var todoList = []Todo{
	{Id: 1, Item: "go to store"},
	{Id: 2, Item: "buy!"},
	{Id: 3, Item: "Sleep"},
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001",
		AllowCredentials: true,
		AllowMethods:     "*",
		AllowHeaders:     "*",
	}))

	app.Get("/todo", getTodos)

	app.Post("/todo", createTodo)

	app.Put("/todo/:id", updateTodo)

	app.Delete("/todo/:id", deleteTodo)

	app.Listen(":8001")
}

func getTodos(c *fiber.Ctx) error {
	return c.JSON(ListTodosResponse{Data: todoList})
}

func createTodo(c *fiber.Ctx) error {
	newTodo := new(Todo)

	if err := c.BodyParser(newTodo); err != nil {
		return err
	}
	todoList = append(todoList, *newTodo)

	return nil
}

func deleteTodo(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	for i, todo := range todoList {
		if todo.Id == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
		}
	}
	return nil
}

func updateTodo(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	tempTodo := new(Todo)
	if err := c.BodyParser(tempTodo); err != nil {
		return err
	}

	for i, todo := range todoList {
		if todo.Id == id {
			todoList[i] = *tempTodo
		}
	}
	return nil
}
