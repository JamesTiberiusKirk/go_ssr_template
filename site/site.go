package site

import (
	"go_web_template/session"
	"go_web_template/site/page"
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
			page.NewHomePage(),
		},
		authedPages: []*page.Page{
			page.NewUserPage(db, sessionManager),
			page.NewUserSSRPage(db, sessionManager),
		},
		db:             db,
		sessionManager: sessionManager,
		echo:           e,
		frameTmpl:      "frame",
		tmplFuncs: template.FuncMap{
			"stringify": stringyfyJson,
		},
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

		if p.PostHandler != nil {
			s.echo.POST(s.rootSitePath+p.Path, p.PostHandler, middlewares...)
		}

		if p.DeleteHandler != nil {
			s.echo.DELETE(s.rootSitePath+p.Path, p.DeleteHandler, middlewares...)
		}

		if p.PutHandler != nil {
			s.echo.PUT(s.rootSitePath+p.Path, p.PutHandler, middlewares...)
		}
	}
}
