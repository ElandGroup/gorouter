package service

import (
	"net/http"
"github.com/labstack/echo"

)

type (
	APIResult struct {
		Result  interface{} `json:"result"`
		Success bool        `json:"success"`
		Error   APIError    `json:"error"`
	}

	APIError struct {
		Code    int         `json:"code"`
		Details interface{} `json:"details"`
		Message string      `json:"message"`
	}
)

type User struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func GetUser(c echo.Context) error {
	user := User{}
	user.Name = "xiaomiao"
	user.Sex = "Male"
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: user})

}
