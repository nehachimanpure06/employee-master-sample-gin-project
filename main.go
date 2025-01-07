package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Welcome to the employee master data application...")

	log.Info("starting the server")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to the employee master data",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Error("error in starting the server :", err.Error())
	}
}
