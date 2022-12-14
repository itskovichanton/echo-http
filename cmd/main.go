package main

import (
	"fmt"
	"github.com/itskovichanton/echo-http"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		println(c.QueryParam("a"))
		return nil
	})
	err := e.Start(fmt.Sprintf(":%v", 3009))
	println(err.Error())
}
