package controllers

import (
	"gin-berry/berry"
	"gin-berry/models"
	"github.com/gin-gonic/gin"
	"log"
)

type RouteConfig struct {
	ForceAuth   bool
	ResolveUser bool
}

type QueryParams struct {
	Username string `validate:"required" json:"username" msg_required:"User name is required!"`
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
			var user models.User
			state, paging := user.GetUsers(1, 20)
			ctx.JSON(200, gin.H{
				"results": state,
				"paging":  paging,
			})
		},
		Options: berry.RouterOptions{
			// we will require that a `Username` value must exist in the request query string.
			QueryString: QueryParams{},
		},
		Config: RouteConfig{
			ForceAuth:   true,
			ResolveUser: false,
		},
	}
}
