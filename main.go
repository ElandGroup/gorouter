package main

import (
	"net/http"
	"gorouter/logs"
	//"gorouter/api"
	"github.com/labstack/echo"
)
var (
e = echo.New()
)
func main() {
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo api")
	})
	RegisterUserApi()
	logs.Warning.Println("RegisterUserApi end")
	
	e.Start(":1111")
}

