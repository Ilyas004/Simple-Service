package app

import (
	"fmt"
	"log"
	"main/internal/app/endpoint"
	"main/internal/app/mv"
	"main/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mv.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server running")

	err := a.echo.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
