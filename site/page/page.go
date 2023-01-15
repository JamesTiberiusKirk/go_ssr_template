package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PageMetaData is used to give certain page meta data and basic params to each template
type PageMetaData struct {
	MenuID   string
	Title    string
	UrlError string
	Success  string
}

// Page is used by every page in a site
// Deps being each page's own struct for dependencies, might not even be needed
type Page struct {
	MenuID        string
	Title         string
	Frame         bool
	Path          string
	Template      string
	Deps          interface{}
	GetPageData   func(c echo.Context) any
	PostHandler   echo.HandlerFunc
	DeleteHandler echo.HandlerFunc
	PutHandler    echo.HandlerFunc
}

// ContextParams is for the context passed to the controllers
type ContextParams struct {
	UseFrame bool
}

const (
	UseFrameName = "frame"
)

func createContext(frame bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(UseFrameName, ContextParams{UseFrame: frame})
			return next(c)
		}
	}
}

// GetPageHandler is a get handler which uses the echo Render function
func (p *Page) GetPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(UseFrameName, p.Frame)

		err := c.Render(http.StatusOK, p.Template, echo.Map{
			"data": p.GetPageData(c),
			"meta": p.buildBasePageMetaData(c),
			// "auth":
		})

		if err != nil {
			return err
		}

		return nil
	}
}

func (p *Page) buildBasePageMetaData(c echo.Context) PageMetaData {
	return PageMetaData{
		MenuID:   p.MenuID,
		Title:    p.Title,
		UrlError: c.QueryParam("error"),
		Success:  c.QueryParam("success"),
	}
}
