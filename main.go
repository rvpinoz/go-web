package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-web/config"
	"go-web/routes"
	"go-web/tools"
	"os"
)

var err error

func main() {
	start := tools.Now()
	tools.Logger()
	err = godotenv.Load()
	if err != nil {
		go tools.Fatal(err.Error(), "File Environment Not Found", start)
	}
	gin.SetMode(os.Getenv("MODE"))
	config.Mysql()
	//config.DB.AutoMigrate(&models.TableToken{}) //For Automate Create Table When Service Run
	r := routes.SetupRouter()
	r.Run()
}
