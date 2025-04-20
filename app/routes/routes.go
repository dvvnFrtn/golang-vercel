package routes

import (
	"golang-vercel/app/handler"
	"net/http"

	_ "golang-vercel/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.NoRoute(ErrRouter)

	app.GET("/ping", handler.Ping)

	route := app.Group("/api")
	{
		route.GET("/hello/:name", handler.Hello)
		route.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		route.GET("/users/me", handler.GetUser)
	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
