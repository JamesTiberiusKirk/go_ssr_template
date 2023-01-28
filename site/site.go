package site

import (
	"go_web_template/server"
	"go_web_template/session"
	"go_web_template/site/page"
	"go_web_template/site/renderer"
	"go_web_template/site/spa"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// Site site struct with config and dependencies
type Site struct {
	rootSitePath   string
	publicPages    []*page.Page
	authedPages    []*page.Page
	notFoundPage   *page.Page
	staticFolders  map[string]string
	spaSites       []*spa.Site
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
		staticFolders: map[string]string{
			"/static": "site/static/",
			"/assets": "site/assets/",
		},
		spaSites: []*spa.Site{
			spa.NewReactPortal(),
		},
		notFoundPage:   page.NewNotFoundPage(),
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

	// Mapping 404 page
	s.echo.GET(s.rootSitePath+s.notFoundPage.Path,
		s.notFoundPage.GetPageHandler(http.StatusNotFound, *s.sessionManager, s.routes))

	s.mapStatic()
	s.mapSpaSites()
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

func (s *Site) mapSpaSites(middlewares ...echo.MiddlewareFunc) {
	for _, spa := range s.spaSites {
		route := s.rootSitePath + spa.Path

		group := s.echo.Group(route)
		group.Use(echoMw.StaticWithConfig(echoMw.StaticConfig{
			Root:   spa.Dist,
			Index:  spa.Index,
			Browse: false,
			HTML5:  true,
		}))
	}
}

func (s *Site) mapStatic() {
	for k, v := range s.staticFolders {
		s.echo.Static(k, v)
	}
}

func (s *Site) mapPages(pages *[]*page.Page, middlewares ...echo.MiddlewareFunc) {
	for _, p := range *pages {
		route := s.rootSitePath + p.Path
		s.routes["site"][p.MenuID] = route
	}

	for _, p := range *pages {
		route := s.rootSitePath + p.Path
		s.echo.GET(route, p.GetPageHandler(http.StatusOK, *s.sessionManager, s.routes), middlewares...)

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
