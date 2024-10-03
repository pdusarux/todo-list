package repository

import (
	"todo-list/entities"

	"gorm.io/gorm"
)

// TodoRepository เป็น interface ที่กำหนดว่า repository จะต้องทำงานอะไรบ้าง
type TodoRepository interface {
	FindAll() ([]entities.Todolist, error)        // ดึงรายการ Todo ทั้งหมด
	FindByID(id uint) (*entities.Todolist, error) // ดึงข้อมูล Todo ตาม ID
	Create(todo *entities.Todolist) error         // สร้างรายการ Todo ใหม่
	UpdateStatus(id uint, status bool) error      // อัปเดตสถานะของ Todo
	Delete(id uint) error                         // ลบรายการ Todo
}

type todoRepositoryImpl struct {
	db *gorm.DB
}

// NewTodoRepository สร้าง instance ใหม่ของ todoRepositoryImpl
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepositoryImpl{db}
}

// FindAll ดึงรายการ Todo ทั้งหมด
func (r *todoRepositoryImpl) FindAll() ([]entities.Todolist, error) {
	var todos []entities.Todolist
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// FindByID ดึงรายการ Todo โดยใช้ ID
func (r *todoRepositoryImpl) FindByID(id uint) (*entities.Todolist, error) {
	var todo entities.Todolist
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// Create สร้างรายการ Todo ใหม่
func (r *todoRepositoryImpl) Create(todo *entities.Todolist) error {
	return r.db.Create(todo).Error
}

// UpdateStatus อัปเดตสถานะของ Todo
func (r *todoRepositoryImpl) UpdateStatus(id uint, status bool) error {
	return r.db.Model(&entities.Todolist{}).Where("id = ?", id).Update("status", status).Error
}

// Delete ลบรายการ Todo ตาม ID
func (r *todoRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Todolist{}, id).Error
}
