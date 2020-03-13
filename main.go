package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/synw/fwr/ws"
)

var v = flag.Bool("v", false, "Verbosity of the output: -v to enable")

func main() {
	flag.Parse()
	var verbose bool
	flag.BoolVar(&verbose, "verbose", true, "")
	if verbose {
		fmt.Println("Running websockets server")
	}
	go ws.RunWs()
	if verbose {
		fmt.Println("Running changes watcher")
	}
	go watch(verbose)
	if verbose {
		fmt.Println("Running http server")
	}
	runServer()
}

func runServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "build/web")

	e.Logger.Fatal(e.Start(":8085"))
}
