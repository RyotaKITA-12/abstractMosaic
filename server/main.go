package main

import (
	"github.com/RyotaKITA-12/abstractMosaic.git/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

    r.GET("/images", handler.List)
    r.POST("/images", handler.Upload)
    r.DELETE("/images/:uuid", handler.Delete)
	r.Run(":8888")
}
