package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/synw/fwr/ws"
)

func main() {
	go ws.RunWs()
	go watch()
	runServer()
}

func runServer() {
	e := echo.New()

	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "build/web")

	e.Logger.Fatal(e.Start(":8085"))
}
