package welcome

import (
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
)

// App data
type App struct {
	Name string
}

// Index function
func Index(req *request.HTTP) response.HTTP {
	return response.View("welcome.gohtml", App{Name: "Pulsar"})
}
