package userhdl

import (
	"net/http"
	"pi/internal/core/domains"
	"pi/internal/core/ports"
	"pi/internal/dto"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	s ports.UserService
}

func New(s ports.UserService) *Handler {
	return &Handler{s: s}
}

// @Summary  Create new user
// @Tags     User
// @Accept   json
// @Produce  json
// @Router   /user [post]
// @Param user body dto.CreateUserRequest true "user to create"
// @Response 200      {object} dto.BaseResponseWithData[dto.CreateUserResponse]
// @Response 400      {object} dto.BaseErrorResponse
// @Response 500      {object} dto.BaseErrorResponse
func (h *Handler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate request
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return err
	}

	// create user
	user, err := h.s.CreateUser(ctx, &domains.CreateUserRequest{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.BaseResponseWithData[dto.CreateUserResponse]{
		BaseResponse: dto.BaseResponse{
			Code: 0,
		},
		Data: dto.CreateUserResponse{
			ID: user.ID,
		},
	})
}

// @Summary  Get user by ID
// @Tags     User
// @Accept   json
// @Produce  json
// @Router   /user/{id} [get]
// @Param    id  path string true "id"
// @Response 200 {object} dto.BaseResponseWithData[dto.User]
// @Response 400 {object} dto.BaseErrorResponse
// @Response 500 {object} dto.BaseErrorResponse
func (h *Handler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.GetUserByIDRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate request
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return err
	}

	// login
	res, err := h.s.GetUserByID(ctx, req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.BaseResponseWithData[dto.User]{
		BaseResponse: dto.BaseResponse{
			Code: 0,
		},
		Data: dto.User{
			ID:       res.ID,
			Username: res.Username,
			Email:    res.Email,
		},
	})
}

// @Summary  Update user
// @Tags     User
// @Accept   json
// @Produce  json
// @Router   /user/{id} [put]
// @Param    id   path string true "id"
// @Param    user body dto.CreateUserRequest true "user to update"
// @Response 200        {object} dto.BaseResponse
// @Response 400        {object} dto.BaseErrorResponse
// @Response 500        {object} dto.BaseErrorResponse
func (h *Handler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate request
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return err
	}

	// update user
	err := h.s.Update(ctx, req.ID, &domains.UpdateUserRequest{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.BaseResponse{
		Code: 0,
	})
}

// @Summary  Delete user
// @Tags     User
// @Accept   json
// @Produce  json
// @Router   /user/{id} [delete]
// @Param    id  path string true "id"
// @Response 200 {object} dto.BaseResponse
// @Response 400 {object} dto.BaseErrorResponse
// @Response 500 {object} dto.BaseErrorResponse
func (h *Handler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	var req dto.DeleteUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate request
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return err
	}

	// delete user
	err := h.s.Delete(ctx, req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.BaseResponse{
		Code: 0,
	})
}
