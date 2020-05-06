package main

import (
	"os"

	_ "beego-hello/routers"

	"github.com/astaxie/beego"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	beego.Run(":" + port)
}
