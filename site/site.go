package site

import (
	"html/template"
	"music_manager/session"
	"music_manager/site/page"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Site site struct with config and dependencies
type Site struct {
	rootSitePath   string
	publicPages    []page.PageInterface
	authedPages    []page.PageInterface
	db             *gorm.DB
	sessionManager *session.Manager
	echo           *echo.Echo
	frameTmpl      string
	tmplFuncs      template.FuncMap
	sessionSecret  string
}

// NewSite init Site
func NewSite(rootSitePath string, db *gorm.DB, sessionManager *session.Manager, e *echo.Echo,
	sessionSecret string) Site {
	return Site{
		rootSitePath: rootSitePath,
		publicPages: []page.PageInterface{
			page.NewLoginPage(db, sessionManager),
			page.NewSignupPage(db, sessionManager),
		},
		authedPages: []page.PageInterface{
			page.NewHomePage(db),
		},
		db:             db,
		sessionManager: sessionManager,
		echo:           e,
		frameTmpl:      "frame",
		sessionSecret:  sessionSecret,
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

func (s *Site) mapPages(pages *[]page.PageInterface, middlewares ...echo.MiddlewareFunc) {
	for _, p := range *pages {
		s.echo.GET(p.GetPagePath(), p.GetPageHandler(), middlewares...)

		post := p.GetPostHandler()
		if post != nil {
			s.echo.POST(p.GetPagePath(), post, middlewares...)
		}
	}

}
