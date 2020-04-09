/*
@Time :  2020-04-09 15:30
@Author : Caden
@File : mian
@Software: GoLand
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Age  int    `json:"age" xml:"age" form:"age" query:"age"`
}

func main() {
	app := echo.New()
	app.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello world!")
	})
	app.GET("/users/:id", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, ctx.Param("id"))
	})
	app.GET("/show", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		age := ctx.QueryParam("age")
		return ctx.String(http.StatusOK, fmt.Sprintf("name=%v,age=%v", name, age))
	})
	app.POST("/form", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		age := ctx.FormValue("age")
		return ctx.String(http.StatusOK, fmt.Sprintf("name=%v,age=%v", name, age))
	})
	app.POST("/form", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		age := ctx.FormValue("age")
		return ctx.String(http.StatusOK, fmt.Sprintf("name=%v,age=%v", name, age))
	})
	app.POST("/formFile", func(ctx echo.Context) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}
		open, err := file.Open()
		defer open.Close()
		reader := bufio.NewReader(open)
		var buffer bytes.Buffer
		for {
			red, err := reader.ReadString('\n')
			buffer.WriteString(red)
			//fmt.Print(readString)
			if err != nil {
				if err != io.EOF {
					return err
				} else {
					break
				}
			}
		}
		value := ctx.FormValue("name")
		return ctx.HTML(http.StatusOK, fmt.Sprintf("<b>Thank you! %v</b></br><b>%s</b>", value, buffer.String()))
	})
	app.POST("/json", func(ctx echo.Context) error {
		u := new(User)
		if err := ctx.Bind(u); err != nil { //根据tag标签和head Content-Type自动转换json,xml,form,query
			return err
		}
		return ctx.JSON(http.StatusOK, u)

	})
	app.Logger.Fatal(app.Start(":8080"))
}
