package controllers

import "github.com/gin-gonic/gin"

func Greetings(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	ctx.JSON(200, gin.H{
		"message": "Hello World, " + name,
	})
}
