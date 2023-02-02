package page

import (
	"github.com/JamesTiberiusKirk/go_web_template/server"
	"github.com/JamesTiberiusKirk/go_web_template/session"

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
func (p *Page) GetPageHandler(httpStatus int, session session.Manager, routesMap map[string]server.RoutesMap) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(UseFrameName, p.Frame)
		auth := echo.Map{}
		user, err := session.GetUser(c)
		if err != nil {
			if err.Error() == "securecookie: the value is not valid" {
				auth = echo.Map{}
			} else {
				return err
			}
		} else {
			auth = echo.Map{
				"email":    user.Email,
				"username": user.Username,
			}
		}

		echoData := echo.Map{
			"meta":   p.buildBasePageMetaData(c),
			"auth":   auth,
			"routes": routesMap,
		}

		pageData := p.GetPageData(c)
		if pageData != nil {
			echoData["data"] = pageData
		}

		err = c.Render(httpStatus, p.Template, echoData)
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
