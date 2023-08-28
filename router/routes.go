package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{

				"msg": "Get opening",
			})
		})
	}

	{
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{

				"msg": "POST opening",
			})
		})
	}
	{
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{

				"msg": "DELETE opening",
			})
		})
	}

	{
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{

				"msg": "PUT opening",
			})
		})
	}
	{
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{

				"msg": "GET ALL openings",
			})
		})
	}

}
