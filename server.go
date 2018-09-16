// https://echo.labstack.com/cookbook/crud

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"clean-arch/controller/user"
	"clean-arch/controller/post"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Route => handler
	e.GET("/", UserController.HelloWorld)
	e.GET("/users", UserController.Index)
	e.POST("/users", UserController.Create)
	e.GET("/users/:id", UserController.Show)
	e.PUT("/users/:id", UserController.Update)
	e.DELETE("/users/:id", UserController.Delete)
	
	// Posts Route
	e.GET("/posts", PostController.Index)
	e.POST("/posts", PostController.Create)
	e.GET("/posts/:id", PostController.Show)
	e.PUT("/posts/:id", PostController.Update)
	e.DELETE("/posts/:id", PostController.Delete)


	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
