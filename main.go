package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	os.Exit(exec())
}

func exec() int {
	r, err := router()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	_ = r.Run(":8081")

	return 0
}

func router() (*gin.Engine, error) {
	r := gin.Default()

	health := r.Group("api")
	health.GET("/healthcheck", HealthCheck)

	group := r.Group("api")
	group.Use(VerifyToken())
	group.GET("/hello", Hello)

	return r, nil
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// トークン検証処理を書く
		return
	}
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "hello"})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
