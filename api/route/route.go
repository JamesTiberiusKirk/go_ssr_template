package route

import "github.com/labstack/echo/v4"

type Route struct {
	SubRoute      *Route
	Path          string
	Depts         any
	GetHandler    echo.HandlerFunc
	PostHandler   echo.HandlerFunc
	DeleteHandler echo.HandlerFunc
	PutHandler    echo.HandlerFunc
}

func (r *Route) Init(rootPath string, e *echo.Group, middlewares ...echo.MiddlewareFunc) {
	if r.GetHandler != nil {
		e.GET(rootPath+r.Path, r.GetHandler, middlewares...)
	}

	if r.PostHandler != nil {
		e.POST(rootPath+r.Path, r.PostHandler, middlewares...)
	}

	if r.DeleteHandler != nil {
		e.DELETE(rootPath+r.Path, r.DeleteHandler, middlewares...)
	}

	if r.PutHandler != nil {
		e.PUT(rootPath+r.Path, r.PutHandler, middlewares...)
	}

	if r.SubRoute != nil {
		r.SubRoute.Init(r.Path, e, middlewares...)
	}
}
