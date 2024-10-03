package usecases

import (
	"todo-list/entities"
	"todo-list/repository"
)

type TodoUseCase interface {
	GetAllTodos() ([]entities.Todolist, error)
	CreateTodo(todo *entities.Todolist) error
	ToggleTodoStatus(id uint) error
	DeleteTodo(id uint) error
}

// todoUseCase เป็น struct ที่ทำงานตาม use case ที่ระบุ
type todoUseCase struct {
	repo repository.TodoRepository
}

// NewTodoUseCase สร้าง instance ใหม่ของ todoUseCase
func NewTodoUseCase(repo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{repo}
}

// GetAllTodos ดึงรายการ Todo ทั้งหมดจาก repository
func (u *todoUseCase) GetAllTodos() ([]entities.Todolist, error) {
	return u.repo.FindAll()
}

// CreateTodo สร้าง Todo ใหม่ผ่าน repository
func (u *todoUseCase) CreateTodo(todo *entities.Todolist) error {
	return u.repo.Create(todo)
}

// ToggleTodoStatus อัปเดตสถานะของ Todo ผ่าน repository
func (u *todoUseCase) ToggleTodoStatus(id uint) error {
	todo, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	todo.Status = !todo.Status
	return u.repo.UpdateStatus(id, todo.Status)
}

// DeleteTodo ลบ Todo ผ่าน repository
func (u *todoUseCase) DeleteTodo(id uint) error {
	return u.repo.Delete(id)
}
