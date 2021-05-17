package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "project-twit/docs"
<<<<<<< HEAD
	"project-twit/grpc"
=======
>>>>>>> 7d473d640dfa86d23e6f3c9c142c323adc887628
)

func main() {
	r := gin.New()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
<<<<<<< HEAD
	go r.Run()
	grpc.Server()
=======
	r.Run()
>>>>>>> 7d473d640dfa86d23e6f3c9c142c323adc887628
}
