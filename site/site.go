package site

import (
	"go_web_template/server"
	"go_web_template/session"
	"go_web_template/site/page"
	"go_web_template/site/renderer"
	"html/template"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Site site struct with config and dependencies
type Site struct {
	rootSitePath   string
	publicPages    []*page.Page
	authedPages    []*page.Page
	db             *gorm.DB
	sessionManager *session.Manager
	echo           *echo.Echo
	frameTmpls     map[string]string
	tmplFuncs      template.FuncMap
	routes         map[string]server.RoutesMap
}

// NewSite init Site
func NewSite(e *echo.Echo, rootSitePath string, db *gorm.DB,
	sessionManager *session.Manager) *Site {
	return &Site{
		rootSitePath: rootSitePath,
		publicPages: []*page.Page{
			page.NewLoginPage(db, sessionManager),
			page.NewSignupPage(db, sessionManager),
			page.NewHomePage(),
		},
		authedPages: []*page.Page{
			page.NewUserPage(db, sessionManager),
			page.NewUserSSRPage(db, sessionManager),
		},
		db:             db,
		sessionManager: sessionManager,
		echo:           e,
		frameTmpls: map[string]string{
			"frame":    "frame.gohtml",
			"no_frame": "no_frame.gohtml",
		},
		tmplFuncs: template.FuncMap{
			"stringify": stringyfyJson,
		},
		routes: map[string]server.RoutesMap{
			"site": {},
		},
	}
}

// Serve to start the server
func (s *Site) Serve() {
	s.buildRenderer()

	s.mapPages(&s.publicPages)
	s.mapPages(&s.authedPages, session.SessionAuthMiddleware(s.sessionManager))
}

// GetRoutes to get routes which have been made in the server
func (s *Site) GetRoutes() server.RoutesMap {
	return s.routes["site"]
}

// SetRoutes which would be used in the templating engine
func (s *Site) SetRoutes(t string, r server.RoutesMap) {
	s.routes[t] = r
}

func (s *Site) buildRenderer() {
	s.echo.Renderer = renderer.New(renderer.Config{
		Root:         "site/page/templates",
		Master:       s.frameTmpls["frame"],
		NoFrame:      s.frameTmpls["no_frame"],
		Funcs:        s.tmplFuncs,
		DisableCache: true,
	})
}

func (s *Site) mapPages(pages *[]*page.Page, middlewares ...echo.MiddlewareFunc) {
	for _, p := range *pages {
		route := s.rootSitePath + p.Path
		s.routes["site"][p.MenuID] = route
	}

	for _, p := range *pages {
		route := s.rootSitePath + p.Path
		s.echo.GET(route, p.GetPageHandler(*s.sessionManager, s.routes), middlewares...)

		if p.PostHandler != nil {
			s.echo.POST(route, p.PostHandler, middlewares...)
		}

		if p.DeleteHandler != nil {
			s.echo.DELETE(route, p.DeleteHandler, middlewares...)
		}

		if p.PutHandler != nil {
			s.echo.PUT(route, p.PutHandler, middlewares...)
		}
	}
}
