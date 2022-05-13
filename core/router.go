package core

import (
	"errors"
	"gin-berry/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Service defines a high level of application service with gin framework embedded.
type Service struct {
	*gin.RouterGroup
	Engine        *gin.Engine
	HandlersChain []gin.HandlerFunc
}

// ServiceRouterOptions defines payload config.
type ServiceRouterOptions struct {
	QueryString interface{}
	Params      interface{}
	Body        interface{}
}

// ServiceRouterConfig defines the routing handler and the various configuration that will work
// within the route handler.
type ServiceRouterConfig struct {
	Middlewares []gin.HandlerFunc
	Handler     gin.HandlerFunc
	Options     ServiceRouterOptions
	Config      interface{}
}

// ErrorResponse is a generic error response utility fn.
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// New creates a new Service instance.
func New(middleware ...gin.HandlerFunc) *Service {
	app := gin.Default()

	if middleware != nil {
		app.Use(middleware...)
	}

	return &Service{RouterGroup: app.Group("/"), Engine: app}
}

// handleRouterOptions handles the payload validation given in the `ServiceRouterOptions` object.
func (s *Service) handleRouterOptions(config interface{}, opts ServiceRouterOptions) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.Set("routeConfig", config)
		if opts.QueryString != nil {
			if err := c.ShouldBindQuery(opts.QueryString); err != nil {
				var errorResponse any
				apiErrors, _ := utils.BuildAPiError(opts.QueryString, err)
				if apiErrors != nil {
					errorResponse = apiErrors
				} else {
					errorResponse = err.Error()
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorResponse})
				return
			}
			c.Set("queryString", opts.QueryString)
		}
		c.Next()
	}
}

type HttpMethodType string

// Route registers a new route handler to the service.
func (s *Service) Route(method HttpMethodType, path string, conf ServiceRouterConfig) {
	// initial payload validation
	preHandler := s.handleRouterOptions(conf.Config, conf.Options)
	// first execute and mutate the context
	handlers := []gin.HandlerFunc{preHandler}
	// initial middlewares from higher group call
	handlers = append(handlers, s.HandlersChain...)
	// middlewares registered to this route with the `ServiceRouterConfig`
	handlers = append(handlers, conf.Middlewares...)
	// the real request handler
	handlers = append(handlers, conf.Handler)

	switch method {
	case "GET":
		s.GET(path, handlers...)
	case "HEAD":
		s.HEAD(path, handlers...)
	case "POST":
		s.POST(path, handlers...)
	case "PUT":
		s.PUT(path, handlers...)
	case "DELETE":
		s.DELETE(path, handlers...)
	case "PATCH":
		s.PATCH(path, handlers...)
	case "OPTIONS":
		s.OPTIONS(path, handlers...)

	default:
		panic(errors.New("invalid http method"))
	}
}

// Group creates a new router group.
func (s *Service) Group(relativePath string, handlers ...gin.HandlerFunc) *Service {
	return &Service{
		RouterGroup: s.Engine.Group(relativePath, handlers...),
	}
}

// Use adds middleware to the group, see example code in GitHub.
func (s *Service) Use(middleware ...gin.HandlerFunc) *Service {
	s.HandlersChain = append(s.HandlersChain, middleware...)
	return s
}

// Run attaches router to the http.Server and start listening.
func (s *Service) Run(addr ...string) (err error) {
	return s.Engine.Run(addr...)
}
