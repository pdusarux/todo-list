package main

import (
	"todo-list/config"
	"todo-list/handler"
	"todo-list/repository"
	"todo-list/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectingDatabase()

	// สร้าง repository และ use case
	todoRepo := repository.NewTodoRepository(db)
	todoUsecase := usecases.NewTodoUseCase(todoRepo)

	// สร้าง handler
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todo-list", todoHandler.GetTodos)
	r.POST("/create", todoHandler.CreateTodo)
	r.PATCH("/check/:id", todoHandler.ToggleStatus)
	r.DELETE("/delete/:id", todoHandler.DeleteTodo)

	r.Run(":8080")
}
