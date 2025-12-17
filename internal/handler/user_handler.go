package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"

	"user-api/internal/service"
)

type UserHandler struct {
	svc      *service.UserService
	validate *validator.Validate
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		svc:      s,
		validate: validator.New(),
	}
}

// POST /users
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req struct {
		Name string `json:"name" validate:"required"`
		DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.validate.Struct(req); err != nil {
		return fiber.ErrUnprocessableEntity
	}

	user, err := h.svc.Create(c.Context(), req.Name, req.DOB)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GET /users/:id
func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.svc.Get(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

// GET /users
func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.svc.List(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(users)
}

// PUT /users/:id
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req struct {
		Name string `json:"name"`
		DOB  string `json:"dob"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.svc.Update(c.Context(), int32(id), req.Name, req.DOB)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

// DELETE /users/:id
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.svc.Delete(c.Context(), int32(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
