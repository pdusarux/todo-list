package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodolistBody struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      bool   `json:"status" default:"false"`
}

var Db *gorm.DB

func ConnectingDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
		Db.AutoMigrate(&TodolistBody{})
	}
}

func main() {
	ConnectingDatabase()
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todo-list", Todolist)
	r.POST("/create", CreateTodo)
	r.PATCH("/check/:id", CheckTodolist)
	r.DELETE("/delete/:id", DeleteTodo)

	r.Run(":8080")
}

func Todolist(c *gin.Context) {
	var todos []TodolistBody
	result := Db.Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todolist TodolistBody
	if err := c.ShouldBindJSON(&todolist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title and description are required",
		})
		return
	}

	if result := Db.Create(&todolist); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create todo successful",
	})

}

func CheckTodolist(c *gin.Context) {
	id := c.Param("id")
	var todo TodolistBody

	if err := Db.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo not found",
		})
		return
	}

	oldStatus := todo.Status

	result := Db.Model(&TodolistBody{}).Where("id = ?", id).Update("status", !oldStatus)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Todo updated successfully",
		"todo":      todo,
		"newStatus": !oldStatus,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if result := Db.Delete(&TodolistBody{}, id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})

}
