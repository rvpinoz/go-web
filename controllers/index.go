package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"status":  "Success",
	//	"code":    http.StatusOK,
	//	"message": "Service Running Well!!",
	//})
	fmt.Println("go go go")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "Golang Website",
		"header": "Project Golang Website",
	})
}

func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "create.tmpl", gin.H{
		"title":  "Golang Website",
		"header": "Create",
	})
}

func Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "edit.tmpl", gin.H{
		"title":  "Golang Website",
		"header": "Edit",
	})
}

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.tmpl", gin.H{
		"title":  "Golang Website",
		"header": "Detail",
	})
}

func Submit(c *gin.Context) {
	fmt.Println(c.Request.Body)
	fmt.Println(c.Request.PostForm)
	fmt.Println(c.Params)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "Golang Website",
		"header": "Project Golang Website",
	})
}
