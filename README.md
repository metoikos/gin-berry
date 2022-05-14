# gin-berry
Gin-berry is an effort to create an opinionated boilerplate for microservice development.
It adds an extra layer on top of the [gin](https://github.com/gin-gonic/gin) framework as a config proxy to manage middlewares and payload validation in 
a more advanced way.

### Why?
We wanted to rewrite an existing API/microservice with golang, but in go frameworks, we couldn't find some key elements that we
needed. After digging through, we realized that with golang, it's easy to put a wrapper around an existing package and
add additional functionality on top of it.
This attempt is heavily inspired by [fastify framework](https://www.fastify.io/).

We wanted to have fastify's `Route options` within the framework and implemented as an additional layer on top of gin.
In addition, we wanted to modify the behavior of the middleware from the route handler when it's defined. To do that,
we created a `Route` method, and from that, we mapped the handlers to the actual gin routes.

To add some spice, we also integrated [gorm](https://gorm.io/) as a database wrapper.

### Features
- Ability to add global or route group-specific middlewares.
- Ability to add additional context value from route definition to a group middleware.
- Ability to automatically validate query strings, params, and payloads.
- Mimic [fastify's lifecycle](https://www.fastify.io/docs/latest/Reference/Lifecycle/) hooks and ability to
  add `preValidation`, `preHandler` like functionality in a more clear way.
- Ability to add custom error messages to the [go-playground/validator](https://github.com/go-playground/validator).

### Usage

We are wrapping the `gin.Engine` and `gin.RouterGroup` and creating a `berry.Service`.

```go
// setup service
service := berry.New()
```

Or with a global middleware:

```go
// setup service
service := berry.New(func(context *gin.Context) {
    log.Println("Initial service middleware")
    context.Next()
})
```

The `Service` wraps `gin.Default()` so `Logger` and `Recovery` middlewares are automatically added.

You can also add as many middlewares as you wish to the `berry.Service`

```go
// setup service
service := berry.New(func(context *gin.Context) {
    log.Println("Initial service middleware")
    context.Next()
}, CustomMiddleware(), AnotherMiddleware())
```

Now you can handle a request with the `Route` method.

```go
service.Route("GET", "/", controllers.ServiceIndex())
```

`controllers.ServiceIndex()` is an instance of `berry.RouterConfig`.

From the controller you'll have the following for minimal handler:

```go
import (
	"gin-berry/berry"
	"gin-berry/models"
	"github.com/gin-gonic/gin"
)

func ServiceIndex() berry.RouterConfig {
	return berry.RouterConfig{
		Handler: func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"hello":  "world",
			})
		},
	}
}
```

A full example:

```go

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

```

`berry.RouterOptions` can take the following to validate the incoming request.
```go

type RouterOptions struct {
	QueryString interface{}
	Params      interface{}
	Body        interface{}
}
```


### Acknowledgements

This is a work in progress. We are working on a lot of things, but we are not done yet.

Since I am still learning golang and the gin framework, this kind of a side project to learn more about the details
of the language and the framework.
That said, I believe it could be a good starting point for a microservice with a couple of more tweaks.

Still, it needs to be tested, and the additions that we made need to be measured in terms of
performance impact. Especially the custom error message handling part is not well tested.

### Contributors

<table style="border: 0">
<tr style="text-align: center; border: 0">
<td style="text-align: center; border: 0">
<img alt="Yılmaz Uğurlu" src="https://avatars.githubusercontent.com/u/107426?s=32&v=4" />
<br />
<a href="https://github.com/metoikos" target="_blank">metoikos</a>
</td>
<td style="text-align: center; border: 0">
<img alt="Kaya Kapağan" src="https://avatars.githubusercontent.com/u/34680852?s=32&v=4" />
<br />
<a href="https://github.com/kayakapagan" target="_blank">kayakapagan</a>
</td>
</tr>
</table>   

