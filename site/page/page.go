package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
	MenuID           string
	Title            string
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
		pageData := p.GetPageData(c)
		err := c.Render(http.StatusOK, p.Template, echo.Map{
			"data": pageData,
			"meta": p.buildBasePageMetaData(c),
		})

		if err != nil {
			logrus.WithError(err).Error("Template render error")
		}

		return err
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
