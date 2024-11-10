package repository

import (
    "todo-api/internal/domain"
    "gorm.io/gorm"
)

type TodoRepository interface {
    Create(todo *domain.Todo) error
    FindAll(userID uint, search string, status string, priority *domain.Priority, tag string) ([]domain.Todo, error)
    Update(todo *domain.Todo) error
    Delete(id uint, userID uint) error
    FindByID(id uint, userID uint) (*domain.Todo, error)
}

type todoRepository struct {
    db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
    return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
    return r.db.Create(todo).Error
}

func (r *todoRepository) FindAll(userID uint, search string, status string, priority *domain.Priority, tag string) ([]domain.Todo, error) {
    var todos []domain.Todo
    query := r.db.Where("user_id = ?", userID).Preload("Tags")

    if search != "" {
        query = query.Where("title LIKE ?", "%"+search+"%")
    }

    if status != "" {
        query = query.Where("status = ?", status)
    }

    if priority != nil {
        query = query.Where("priority = ?", *priority)
    }

    if tag != "" {
        query = query.Joins("JOIN todo_tags ON todos.id = todo_tags.todo_id").
               Joins("JOIN tags ON todo_tags.tag_id = tags.id").
               Where("tags.name = ?", tag)
    }

    err := query.Find(&todos).Error
    return todos, err
}

func (r *todoRepository) Update(todo *domain.Todo) error {
    return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint, userID uint) error {
    return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Todo{}).Error
}

func (r *todoRepository) FindByID(id uint, userID uint) (*domain.Todo, error) {
    var todo domain.Todo
    err := r.db.Where("id = ? AND user_id = ?", id, userID).Preload("Tags").First(&todo).Error
    if err != nil {
        return nil, err
    }
    return &todo, nil
}