package echow

import (
	"go_web_template/site/page"
	"go_web_template/site/renderer"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const templateEngineKey = "foolin-goview-echoview"

// ViewEngine view engine for echo
type ViewEngine struct {
	*renderer.ViewEngine
}

// New new view engine
func New(config renderer.Config) *ViewEngine {
	return Wrap(renderer.New(config))
}

// Wrap wrap view engine for goview.ViewEngine
func Wrap(engine *renderer.ViewEngine) *ViewEngine {
	return &ViewEngine{
		ViewEngine: engine,
	}
}

// Default new default config view engine
func Default() *ViewEngine {
	return New(renderer.DefaultConfig)
}

// Render render template for echo interface
func (e *ViewEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	frame := c.Get(page.UseFrameName).(bool)
	logrus.Print("FRAME:", frame)
	return e.RenderWriter(w, name, data, frame)
}

// Render html render for template
// You should use helper func `Middleware()` to set the supplied
// TemplateEngine and make `Render()` work validly.
// func Render(ctx echo.Context, code int, name string, data interface{}) error {
// 	if val := ctx.Get(templateEngineKey); val != nil {
// 		if e, ok := val.(*ViewEngine); ok {
// 			return e.Render(ctx.Response().Writer, name, data, ctx)
// 		}
// 	}
// 	return ctx.Render(code, name, data)
// }

// NewMiddleware echo middleware for func `echoview.Render()`
// func NewMiddleware(config renderer.Config) echo.MiddlewareFunc {
// 	return Middleware(New(config))
// }

// Middleware echo middleware wrapper
// func Middleware(e *ViewEngine) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			c.Set(templateEngineKey, e)
// 			return next(c)
// 		}
// 	}
// }
