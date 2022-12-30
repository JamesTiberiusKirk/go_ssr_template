package api

import (
	"go_ssr_template/api/route"
	"go_ssr_template/session"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Api api struct
type Api struct {
	rootApiPath    string
	publicRoutes   []*route.Route
	authedRoutes   []*route.Route
	echoGroup      *echo.Group
	db             *gorm.DB
	sessionManager *session.Manager
}

// NewApi new api instance
func NewApi(group *echo.Group, rootApiPath string, db *gorm.DB, sesessionManager *session.Manager) *Api {
	return &Api{
		rootApiPath:  rootApiPath,
		publicRoutes: []*route.Route{},
		authedRoutes: []*route.Route{
			route.NewUserRoute(db),
		},
		echoGroup:      group,
		db:             db,
		sessionManager: sesessionManager,
	}
}

// Serve api
func (a *Api) Serve() {
	a.mapRoutes(&a.publicRoutes)
	a.mapRoutes(&a.authedRoutes, session.SessionAuthMiddleware(a.sessionManager))
}

func (a *Api) mapRoutes(routes *[]*route.Route, middlewares ...echo.MiddlewareFunc) {
	for _, r := range *routes {
		r.Init("", a.echoGroup, middlewares...)
	}
}
