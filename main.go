package main

import (
	"employee-master/routes"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	log.Info("Welcome to the employee master data application...")
	log.Info("starting the server")
	defer log.Warn("Exiting the server..")

	r := gin.Default()

	routes.InitRoute(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("error in starting the server :", err.Error())
	}
}
