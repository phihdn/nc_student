package main

import (
	"fmt"
	"log"

	"github.com/phihdn/nc_student/config"
	"github.com/phihdn/nc_student/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())

	route.All(e)

	log.Println(e.Start(":9090"))
}
