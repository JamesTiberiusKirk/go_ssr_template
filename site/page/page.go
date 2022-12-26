package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	Title    string
	UrlError string
	Success  string
}

// Page is used by every page in a site
// Deps being each page's own struct for dependencies, might not even be needed
type Page struct {
	Path             string
	Template         string
	Deps             interface{}
	GetPageData      func(c echo.Context) any
	GetPostHandler   echo.HandlerFunc
	GetDeleteHandler echo.HandlerFunc
	GetPutHandler    echo.HandlerFunc
}

// GetPageHandler is a get handler which uses the echo Render function
func (p *Page) GetPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, p.Template, echo.Map{
			"data": p.GetPageData(c),
		})
	}
}

// type PageInterface interface {
// 	GetPageData(c echo.Context) any
// 	GetPagePath() string
// 	GetPageHandler() echo.HandlerFunc
// 	GetPostHandler() echo.HandlerFunc
// }

// type FramePage struct {
// 	MenuId  string
// 	Content PageInterface
// }

// type FramePageData struct {
// 	MenuGroupId string
// 	MenuId      string
// 	Template    string
// 	PageData    any
// }
