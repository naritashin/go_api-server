package main

import (
	_ "github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"

	router "github.com/naritashin/go_api-server/controller"
)

func main() {
	router.Run()
}
