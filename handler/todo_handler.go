package handler

import (
	"net/http"
	"strconv"
	"todo-list/entities"
	"todo-list/usecases"

	"github.com/gin-gonic/gin"
)

// TodoHandler เป็น struct ที่จะรับคำขอจากผู้ใช้และส่งต่อไปยัง Use Case
type TodoHandler struct {
	usecase usecases.TodoUseCase
}

// NewTodoHandler สร้าง instance ใหม่ของ TodoHandler
func NewTodoHandler(usecase usecases.TodoUseCase) *TodoHandler {
	return &TodoHandler{usecase}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.usecase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateTodo สร้างรายการ Todo ใหม่
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo entities.Todolist
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.usecase.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully"})
}

// ToggleStatus อัปเดตสถานะของ Todo ตาม ID
func (h *TodoHandler) ToggleStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.usecase.ToggleTodoStatus(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo status updated"})
}

// DeleteTodo ลบรายการ Todo ตาม ID
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.usecase.DeleteTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
