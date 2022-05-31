package controllers

import (
	"gin-berry/berry"
	"gin-berry/models"
	"log"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	ForceAuth   bool
	ResolveUser bool
}

type QueryParams struct {
	Page  uint8 `form:"page" json:"page"  binding:"number,min=1" msg_min:"Invalid page!" msg_number:"Page must be a number!"`
	Limit uint8 `form:"limit" json:"limit" binding:"number,oneof=10 25 50" msg_number:"Limit must be a number!" msg_oneof:"Invalid limit value!"`
}

func ServiceIndex() berry.RouterConfig {
	return berry.RouterConfig{
		// these will be executed before the route handler
		// but after the group middleware
		Middlewares: []gin.HandlerFunc{func(ctx *gin.Context) {
			log.Println("Pre-route middleware")
		}},
		// this handles the actual route
		Handler: func(ctx *gin.Context) {
			log.Println("Route handler: query", ctx.MustGet("query"))
			var user models.User
			users, paging := user.GetUsers(1, 20)
			ctx.JSON(200, gin.H{
				"users":  users,
				"paging": paging,
			})
		},
		Validation: func() berry.RouterOptions {
			return berry.RouterOptions{
				QueryString: &QueryParams{
					Page:  1, // default values
					Limit: 10,
				},
			}
		},
		Config: RouteConfig{
			ForceAuth:   true,
			ResolveUser: false,
		},
	}
}
