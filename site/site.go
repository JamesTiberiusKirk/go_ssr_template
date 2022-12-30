package site

import (
	"go_ssr_template/session"
	"go_ssr_template/site/page"
	"html/template"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
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
	frameTmpl      string
	tmplFuncs      template.FuncMap
}

// NewSite init Site
func NewSite(rootSitePath string, db *gorm.DB, sessionManager *session.Manager,
	e *echo.Echo) Site {
	return Site{
		rootSitePath: rootSitePath,
		publicPages: []*page.Page{
			page.NewLoginPage(db, sessionManager),
			page.NewSignupPage(db, sessionManager),
		},
		authedPages: []*page.Page{
			page.NewHomePage(db),
		},
		db:             db,
		sessionManager: sessionManager,
		echo:           e,
		frameTmpl:      "frame",
	}
}

// Serve to start the server
func (s *Site) Serve() {
	s.buildRenderer()

	s.mapPages(&s.publicPages)
	s.mapPages(&s.authedPages, session.SessionAuthMiddleware(s.sessionManager))
}

func (s *Site) buildRenderer() {
	s.echo.Renderer = echoview.New(goview.Config{
		Root:         "site/page/templates",
		Extension:    ".gohtml",
		Master:       s.frameTmpl,
		Funcs:        s.tmplFuncs,
		DisableCache: true,
	})
}

func (s *Site) mapPages(pages *[]*page.Page, middlewares ...echo.MiddlewareFunc) {
	for _, p := range *pages {
		s.echo.GET(s.rootSitePath+p.Path, p.GetPageHandler(), middlewares...)

		if p.GetPostHandler != nil {
			s.echo.POST(s.rootSitePath+p.Path, p.GetPostHandler, middlewares...)
		}

		if p.GetDeleteHandler != nil {
			s.echo.DELETE(s.rootSitePath+p.Path, p.GetDeleteHandler, middlewares...)
		}

		if p.GetPutHandler != nil {
			s.echo.PUT(s.rootSitePath+p.Path, p.GetPutHandler, middlewares...)
		}
	}
}
