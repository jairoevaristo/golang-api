package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jairoevaristo/cmd/http/routes"
)

func main() {
	r := gin.Default()
	routes.InitializeRoutes(r)

	r.Run(":3333")
}
