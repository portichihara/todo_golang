package main

import (
    "log"
    "os"
    "todo-api/internal/domain"
    "todo-api/internal/handler"
    "todo-api/internal/repository"
    "todo-api/internal/usecase"
    _ "todo-api/pkg/auth"
    "todo-api/pkg/database"
    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }

    // Initialize database
    if err := database.InitDB(); err != nil {
        log.Fatalf("failed to initialize database: %v", err)
    }
    db := database.DB

    // Drop existing tables
    if err := db.Migrator().DropTable(&domain.User{}, &domain.Todo{}, &domain.Tag{}); err != nil {
        log.Fatalf("failed to drop tables: %v", err)
    }

    // Auto migrate the schema
    if err := db.AutoMigrate(&domain.User{}, &domain.Todo{}, &domain.Tag{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    // Initialize dependencies
    userRepo := repository.NewUserRepository(db)
    todoRepo := repository.NewTodoRepository(db)
    tagRepo := repository.NewTagRepository(db)

    userUseCase := usecase.NewUserUseCase(userRepo)
    todoUseCase := usecase.NewTodoUseCase(todoRepo, tagRepo)

    userHandler := handler.NewUserHandler(userUseCase)
    todoHandler := handler.NewTodoHandler(todoUseCase)

    // Setup Echo
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Public routes
    e.POST("/register", userHandler.Register)
    e.POST("/login", userHandler.Login)

    // Protected routes
    api := e.Group("/api")
    api.Use(handler.JWTMiddleware(os.Getenv("JWT_SECRET")))

    api.POST("/todos", todoHandler.Create)
    api.GET("/todos", todoHandler.GetAll)
    api.PUT("/todos/:id", todoHandler.Update)
    api.DELETE("/todos/:id", todoHandler.Delete)

    // Start server
    if err := e.Start(":8081"); err != nil {
        log.Fatal(err)
    }
}
