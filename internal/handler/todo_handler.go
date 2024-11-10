package handler

import (
    "net/http"
    "strconv"
    "todo-api/internal/domain"
    "todo-api/internal/usecase"
    "github.com/labstack/echo/v4"
)

type TodoHandler struct {
    useCase usecase.TodoUseCase
}

type createTodoRequest struct {
    Title    string         `json:"title"`
    Status   domain.Status  `json:"status"`
    Priority domain.Priority`json:"priority"`
    Tags     []string      `json:"tags"`
}

type updateTodoRequest struct {
    Title    string         `json:"title"`
    Status   domain.Status  `json:"status"`
    Priority domain.Priority`json:"priority"`
    Tags     []string      `json:"tags"`
}

func (h *TodoHandler) Create(c echo.Context) error {
    userID := c.Get("userID").(uint)
    var req createTodoRequest
    if err := c.Bind(&req); err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    err := h.useCase.Create(userID, req.Title, req.Status, req.Priority, req.Tags)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable)
    }

    return c.NoContent(http.StatusNoContent)
}

func (h *TodoHandler) GetAll(c echo.Context) error {
    userID := c.Get("userID").(uint)
    search := c.QueryParam("search")
    status := c.QueryParam("status")
    priorityStr := c.QueryParam("priority")
    tag := c.QueryParam("tag")

    var priority *domain.Priority
    if priorityStr != "" {
        p, err := strconv.Atoi(priorityStr)
        if err == nil {
            prio := domain.Priority(p)
            priority = &prio
        }
    }

    response, err := h.useCase.GetAll(userID, search, status, priority, tag)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable)
    }

    return c.JSON(http.StatusOK, response)
}

func (h *TodoHandler) Update(c echo.Context) error {
    userID := c.Get("userID").(uint)
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    var req updateTodoRequest
    if err := c.Bind(&req); err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    err = h.useCase.Update(uint(id), userID, req.Title, req.Status, req.Priority, req.Tags)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable)
    }

    return c.NoContent(http.StatusNoContent)
}

func (h *TodoHandler) Delete(c echo.Context) error {
    userID := c.Get("userID").(uint)
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        return c.NoContent(http.StatusBadRequest)
    }

    err = h.useCase.Delete(uint(id), userID)
    if err != nil {
        return c.NoContent(http.StatusServiceUnavailable)
    }

    return c.NoContent(http.StatusNoContent)
}

func NewTodoHandler(useCase usecase.TodoUseCase) *TodoHandler {
    return &TodoHandler{
        useCase: useCase,
    }
}
