package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "project-twit/docs"
	"project-twit/methods"
)

func main() {
	r := gin.New()
	url := ginSwagger.URL(methods.GetEnvVariable("SWAGGER_URL",".")) // The url pointing to API definition
	r.GET(methods.GetEnvVariable("SWAGGER_PATH","."), ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	go func() {
		err := r.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
	methods.Configure()
}
