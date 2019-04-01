package routes

import (
	"../app/Http/Controllers/welcome"
	middlewares "../app/Http/Middlewares"

	"github.com/pulsar-go/pulsar/router"
)

// Register registers the application routes.
func Register() {
	router.Routes.
		Group(&router.Options{Prefix: "/", Middleware: middlewares.Sample}, func(routes *router.Router) {
			routes.Get("/", welcome.Index)
		})
}
