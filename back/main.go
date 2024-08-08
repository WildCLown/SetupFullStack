package main

import (
	"go-back/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080") // Porta onde a API vai rodar
}
