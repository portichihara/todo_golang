package usecase

import (
    "errors"
    "todo-api/internal/domain"
    "todo-api/internal/repository"
    "todo-api/pkg/auth"
    "golang.org/x/crypto/bcrypt"
    "os"
)

type UserUseCase interface {
    Register(email, password string) (*domain.UserResponse, error)
    Login(email, password string) (*domain.UserResponse, error)
}

type userUseCase struct {
    repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
    return &userUseCase{
        repo: repo,
    }
}

func (u *userUseCase) Register(email, password string) (*domain.UserResponse, error) {
    // Check if user already exists
    existingUser, _ := u.repo.FindByEmail(email)
    if existingUser != nil {
        return nil, errors.New("email already registered")
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // Create user
    user := &domain.User{
        Email:    email,
        Password: string(hashedPassword),
    }

    if err := u.repo.Create(user); err != nil {
        return nil, err
    }

    // Generate JWT token
    token, err := auth.GenerateToken(user.ID, os.Getenv("JWT_SECRET"))
    if err != nil {
        return nil, err
    }

    return &domain.UserResponse{
        ID:    user.ID,
        Email: user.Email,
        Token: token,
    }, nil
}

func (u *userUseCase) Login(email, password string) (*domain.UserResponse, error) {
    user, err := u.repo.FindByEmail(email)
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    // Check password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    // Generate JWT token
    token, err := auth.GenerateToken(user.ID, os.Getenv("JWT_SECRET"))
    if err != nil {
        return nil, err
    }

    return &domain.UserResponse{
        ID:    user.ID,
        Email: user.Email,
        Token: token,
    }, nil
}