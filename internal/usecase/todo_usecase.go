package usecase

import (
    "todo-api/internal/domain"
    "todo-api/internal/repository"
    "time"
)

type TodoUseCase interface {
    Create(userID uint, title string, status domain.Status, priority domain.Priority, tagNames []string) error
    GetAll(userID uint, search string, status string, priority *domain.Priority, tag string) (*domain.TodoResponse, error)
    Update(id uint, userID uint, title string, status domain.Status, priority domain.Priority, tagNames []string) error
    Delete(id uint, userID uint) error
}

type todoUseCase struct {
    todoRepo repository.TodoRepository
    tagRepo  repository.TagRepository
}

func NewTodoUseCase(todoRepo repository.TodoRepository, tagRepo repository.TagRepository) TodoUseCase {
    return &todoUseCase{
        todoRepo: todoRepo,
        tagRepo:  tagRepo,
    }
}

func (u *todoUseCase) Create(userID uint, title string, status domain.Status, priority domain.Priority, tagNames []string) error {
    // Create or get existing tags
    tags, err := u.getOrCreateTags(tagNames)
    if err != nil {
        return err
    }

    todo := &domain.Todo{
        Title:     title,
        UserID:    userID,
        Status:    status,
        Priority:  priority,
        Tags:      tags,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    return u.todoRepo.Create(todo)
}

func (u *todoUseCase) GetAll(userID uint, search string, status string, priority *domain.Priority, tag string) (*domain.TodoResponse, error) {
    todos, err := u.todoRepo.FindAll(userID, search, status, priority, tag)
    if err != nil {
        return nil, err
    }
    return &domain.TodoResponse{Todos: todos}, nil
}

func (u *todoUseCase) Update(id uint, userID uint, title string, status domain.Status, priority domain.Priority, tagNames []string) error {
    todo, err := u.todoRepo.FindByID(id, userID)
    if err != nil {
        return err
    }

    // Update tags
    tags, err := u.getOrCreateTags(tagNames)
    if err != nil {
        return err
    }

    todo.Title = title
    todo.Status = status
    todo.Priority = priority
    todo.Tags = tags
    todo.UpdatedAt = time.Now()

    return u.todoRepo.Update(todo)
}

func (u *todoUseCase) Delete(id uint, userID uint) error {
    return u.todoRepo.Delete(id, userID)
}

func (u *todoUseCase) getOrCreateTags(tagNames []string) ([]domain.Tag, error) {
    var tags []domain.Tag
    for _, name := range tagNames {
        tag, err := u.tagRepo.FindOrCreate(name)
        if err != nil {
            return nil, err
        }
        tags = append(tags, *tag)
    }
    return tags, nil
}