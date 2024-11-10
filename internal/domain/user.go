package domain

import (
    "time"
)

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Email     string    `json:"email" gorm:"unique"`
    Password  string    `json:"-"` // ハッシュ化されたパスワード
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type UserResponse struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
    Token string `json:"token"`
}