package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	helloWorldApiRoute = "/"
)

// HelloWorld user route dependency struct
type HelloWorld struct {
}

// NewHelloWorld struct instance
func NewHelloWorld() *Route {
	depts := &HelloWorld{}

	return &Route{
		RouteID:    "users",
		Path:       helloWorldApiRoute,
		Depts:      depts,
		GetHandler: depts.Hello,
	}
}

func (r *HelloWorld) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}
