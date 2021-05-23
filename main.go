package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "project-twit/docs"
	"project-twit/grpc"
)

func main() {
	r := gin.New()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	go func() {
		err := r.Run()
		if err != nil {
			return
		}
	}()
	grpc.Server()
}
