package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"krama/configuration"
	"krama/controller"
)

func main() {
	router := echo.New()
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "IP=${remote_ip}, PATH=${path}, METHOD=${method}, LATENCY=${latency_human}\n",
	}))
	router.POST("/authentication", controller.Auth)

	mobile := router.Group("/mobile")
	config := middleware.JWTConfig{
		Claims:     &controller.MyClaim{},
		SigningKey: []byte(configuration.JWT_SECRET),
	}

	mobile.Use(middleware.JWTWithConfig(config))
	mobile.GET("/verify", controller.Verify)
	mobile.GET("/users", controller.GetAllAccount)
	mobile.GET("/user", controller.GetAccountById)
	mobile.POST("/user", controller.InsertAccount)
	mobile.PUT("/user", controller.UpdateAccountById)
	mobile.DELETE("/user", controller.DeleteAccountById)

	router.Start(":3000")

/*	var username string = "APP-Mobile"
	var password string = "$2y$12$aXdvFMVniT1NLGH3X0JTr.pcBy/8yuXI0ykIT87Ixq6M357gZep0S"
*/

}
