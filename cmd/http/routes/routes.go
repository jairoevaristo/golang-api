package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jairoevaristo/cmd/http/handler"
)

func InitializeRoutes(r *gin.Engine) {
	{
		r.POST("/short-link", handler.CreateShortLink)
		r.GET("/:hash", handler.GetShortLink)
	}
}
