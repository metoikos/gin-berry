package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Service struct {
	*gin.RouterGroup
	Engine        *gin.Engine
	HandlersChain []gin.HandlerFunc
}

type ServiceRouterOptions struct {
	QueryString interface{}
	Params      interface{}
	Body        interface{}
}

type ServiceRouterConfig struct {
	Middlewares []gin.HandlerFunc
	Handler     gin.HandlerFunc
	Options     ServiceRouterOptions
	Config      interface{}
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func New(middleware ...gin.HandlerFunc) *Service {
	app := gin.Default()
	if middleware != nil {
		app.Use(middleware...)
	}
	return &Service{RouterGroup: app.Group("/"), Engine: app}
}

func (r *Service) HandleRouterOptions(config interface{}, opts ServiceRouterOptions) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		log.Println("HandleRouterOptions")
		c.Set("routeConfig", config)
		if opts.QueryString != nil {
			if err := c.ShouldBindQuery(opts.QueryString); err != nil {
				c.JSON(http.StatusBadRequest, ErrorResponse(err))
				c.Abort()
				return
			}
			c.Set("_queryString", opts.QueryString)
		}
		c.Next()
	}
}

type HttpMethodType string

func (r *Service) Route(method HttpMethodType, path string, conf ServiceRouterConfig) {
	// initial payload validation
	preHandler := r.HandleRouterOptions(conf.Config, conf.Options)
	// first execute and mutate the context
	handlers := []gin.HandlerFunc{preHandler}
	// initial middlewares from higher
	handlers = append(handlers, r.HandlersChain...)
	// middlewares registered to this route
	handlers = append(handlers, conf.Middlewares...)
	// the real request handler
	handlers = append(handlers, conf.Handler)

	switch method {
	case "GET":
		r.GET(path, handlers...)
	case "HEAD":
		r.HEAD(path, handlers...)
	case "POST":
		r.POST(path, handlers...)
	case "PUT":
		r.PUT(path, handlers...)
	case "DELETE":
		r.DELETE(path, handlers...)
	case "PATCH":
		r.PATCH(path, handlers...)
	case "OPTIONS":
		r.OPTIONS(path, handlers...)

	default:
		panic(errors.New("invalid http method"))
	}
}

func (r *Service) Group(relativePath string, handlers ...gin.HandlerFunc) *Service {
	return &Service{
		RouterGroup: r.Engine.Group(relativePath, handlers...),
	}
}

// Use adds middleware to the group, see example code in GitHub.
func (r *Service) Use(middleware ...gin.HandlerFunc) *Service {
	r.HandlersChain = append(r.HandlersChain, middleware...)
	return r
}

func (r *Service) Run(addr ...string) (err error) {
	return r.Engine.Run(addr...)
}
