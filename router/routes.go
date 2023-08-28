package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliocesar014/go-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"api_version": "1.0.0.0",
				"status":      "Ok, running...",
			})
		})
	}

	{
		v1.GET("/opening", handler.ShowOpeningHandler)
	}

	{
		v1.POST("/opening", handler.CreateOpeningHandler)
	}

	{
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

	}

	{
		v1.PUT("/opening", handler.UpdateOpeningHandler)

	}

	{
		v1.GET("/openings", handler.ListOpeningsHandler)

	}

}
