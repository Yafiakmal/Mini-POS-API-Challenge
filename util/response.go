package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string      `json:"status"`            // "success" / "error"
	Message string      `json:"message,omitempty"` // optional
	Data    interface{} `json:"data,omitempty"`    // payload
	Error   interface{} `json:"error,omitempty"`   // error detail
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status: "success",
		Data:   data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Status: "success",
		Data:   data,
	})
}

func Error(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, APIResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
