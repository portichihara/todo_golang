package domain

import (
    "time"
)

type Priority int

const (
    PriorityLow Priority = iota + 1
    PriorityMedium
    PriorityHigh
)

type Status string

const (
    StatusNotStarted Status = "NOT_STARTED"
    StatusInProgress Status = "IN_PROGRESS"
    StatusCompleted  Status = "COMPLETED"
)

type Todo struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title"`
    UserID      uint      `json:"userId"`
    Status      Status    `json:"status"`
    Priority    Priority  `json:"priority"`
    Tags        []Tag     `json:"tags" gorm:"many2many:todo_tags;"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

type TodoResponse struct {
    Todos []Todo `json:"todos"`
}