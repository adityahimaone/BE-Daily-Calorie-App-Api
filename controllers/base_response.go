package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Meta struct {
		Code     int      `json:"code"`
		Status   string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Code = http.StatusOK
	response.Meta.Status = "success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

func NewSuccessCreatedResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Code = http.StatusCreated
	response.Meta.Status = "success"
	response.Data = data

	return c.JSON(http.StatusCreated, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = "Something not right"
	response.Meta.Code = status
	response.Meta.Messages = []string{err.Error()}
	return c.JSON(status, response)
}
