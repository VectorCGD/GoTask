package main

import (
	_ "BockWeb/Model"
	router "BockWeb/Router"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Create("ginlog")
	if err != nil {
		fmt.Println("Create file error! err:", err)
	}
	gin.DefaultWriter = io.MultiWriter(file)

	ser := gin.Default()
	ser.LoadHTMLGlob("static/*")
	ser.Static("/static", "./static")
	router.SetUpRouter(ser)

	ser.Run(":8080")
	file.Close()
}
