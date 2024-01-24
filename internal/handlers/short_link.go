package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ved2pj/Duanlink/internal/models"
	"github.com/ved2pj/Duanlink/internal/services"
)

type ShortLinkHandler struct {
	shortLinkService services.ShortLinkService
}

func NewShortLinkHandler(shortLinkService services.ShortLinkService) *ShortLinkHandler {
	return &ShortLinkHandler{
		shortLinkService: shortLinkService,
	}
}

func (handler *ShortLinkHandler) Create(c *gin.Context) {
	var req models.CreateShortLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if shortLink, err := handler.shortLinkService.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, shortLink)
	}
}

func (handler *ShortLinkHandler) Lookup(c *gin.Context) {
	shortCode := c.Param("short_code")
	if shortLink, err := handler.shortLinkService.LookupByShortCode(shortCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, shortLink)
	}
}

func (handler *ShortLinkHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("short_code")
	if shortLink, err := handler.shortLinkService.LookupByShortCode(shortCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.Redirect(http.StatusFound, shortLink.OriginURL)
	}
}
