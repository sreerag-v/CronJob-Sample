package main

import (
	"GoCron/CronJob-Api/Cron"
	"GoCron/CronJob-Api/api"

	"github.com/gin-gonic/gin"
)

func main() {
	Cron.RunCron()
	
	router := gin.Default()

	router.GET("/message", api.SendMessage)

	router.Run(":8080")
}
