package main

import (
	"idea-generator-api/idea-generator/database"
	"idea-generator-api/idea-generator/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	database.Init()
	routes.IdeaRoute(router)
	router.Run(":3172")
}
