package controllers

import (
	"gin-berry/core"
	"gin-berry/models"
	"github.com/gin-gonic/gin"
)

func ServiceIndex() core.ServiceRouterConfig {
	return core.ServiceRouterConfig{
		Handler: func(ctx *gin.Context) {
			var user models.User
			state, paging := user.GetUsers(1, 20)
			ctx.JSON(200, gin.H{
				"results": state,
				"paging":  paging,
			})
		},
	}
}
