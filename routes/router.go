package routes

import (
	"go-web/controllers"
	"go-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		response := models.DefaultResponse{Status: "Failed", Code: http.StatusNotFound, Message: "Not Found"}
		c.JSON(http.StatusNotFound, response)
	})

	r.LoadHTMLGlob("views/*.tmpl")
	index := r.Group("/")
	{
		index.GET("", controllers.GetIndex)
		index.GET("create", controllers.Create)
		index.GET("submit", controllers.Submit)
		index.GET("edit", controllers.Edit)
		index.GET("detail", controllers.Detail)
	}

	return r
}
