package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jairoevaristo/cmd/http/internal/entities"
	"github.com/jairoevaristo/cmd/http/internal/repositories"
)

type CreateShortLinkRequest struct {
	Url string `json:"url" binding:"required"`
}

func CreateShortLink(ctx *gin.Context) {
	request := CreateShortLinkRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Field Url is required",
		})
		return
	}

	link, err := entities.NewLink(request.Url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error of parsed json object",
		})
		return
	}

	repositories.Insert(link.Hash, link.OriginalUrl)

	ctx.IndentedJSON(http.StatusOK, link)
}

func GetShortLink(ctx *gin.Context) {
	hash := ctx.Params.ByName("hash")
	url := repositories.FindLinkByHash(hash)

	ctx.Redirect(http.StatusMovedPermanently, url)
}
