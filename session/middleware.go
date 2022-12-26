package session

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	echoSession "github.com/labstack/echo-contrib/session"
)

// SessionAuthMiddleware middleware for authenticating sessions
func SessionAuthMiddleware(sessionManager *Manager) echo.MiddlewareFunc {
	return echoSession.MiddlewareWithConfig(echoSession.Config{
		Skipper: func(c echo.Context) bool {
			skip := sessionManager.IsAuthenticated(c)
			log.Printf("Middleware skipper func: %t", skip)

			if !skip {
				c.Redirect(http.StatusSeeOther, "/login")
			}

			return skip
		},
		Store: sessionManager.Jar,
	})
}
