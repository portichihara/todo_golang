package handler

import (
    "net/http"
    "todo-api/internal/usecase"
    "github.com/labstack/echo/v4"
)

type UserHandler struct {
    useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
    return &UserHandler{useCase: useCase}
}

type registerRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *UserHandler) Register(c echo.Context) error {
    var req registerRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
    }

    user, err := h.useCase.Register(req.Email, req.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c echo.Context) error {
    var req loginRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
    }

    user, err := h.useCase.Login(req.Email, req.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
    }

    return c.JSON(http.StatusOK, user)
}
