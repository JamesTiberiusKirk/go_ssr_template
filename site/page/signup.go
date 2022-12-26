package page

import (
	"fmt"
	"go_ssr_template/models"
	"go_ssr_template/session"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SignupPage struct {
	path           string
	template       string
	db             *gorm.DB
	sessionManager *session.Manager
}

type SignupPageData struct {
	Title      string
	Message    string
	UrlError   string
	Validation models.User
}

func NewSignupPage(db *gorm.DB, sessionManager *session.Manager) *SignupPage {
	return &SignupPage{
		path:           "/signup",
		template:       "signup",
		db:             db,
		sessionManager: sessionManager,
	}
}

func (p *SignupPage) GetPageData(c echo.Context) any {
	return SignupPageData{
		Title:    "Signup page",
		Message:  c.QueryParam("message"),
		UrlError: c.QueryParam("error"),
		Validation: models.User{
			Email:    c.QueryParam("email"),
			Username: c.QueryParam("username"),
			Password: c.QueryParam("password"),
		},
	}
}

func (p *SignupPage) GetPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{
			Email:    c.FormValue("email"),
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}

		notPassed, err := user.Validate()
		if err != nil {
			logrus.
				WithError(err).
				Error("Error validating signup form")
			return c.Redirect(http.StatusSeeOther, p.path+"?error=internal server error")
		}

		if len(notPassed) != 0 {
			query := "?error=Unable to validate data"
			for _, fields := range notPassed {
				query = fmt.Sprintf("%s&%s=not valid", query, fields)
			}
			return c.Redirect(http.StatusSeeOther, p.path+query)
		}

		user.SetPassword(user.Password)

		result := p.db.WithContext(c.Request().Context()).Create(&user)
		if result.Error != nil {
			msg := "failed to insert user into db"
			logrus.
				WithError(result.Error).
				Error(msg)

			redirectURL := fmt.Sprintf("%s?error=Internal server error", p.path)
			c.Redirect(http.StatusSeeOther, redirectURL)
		}

		return c.Redirect(http.StatusSeeOther, p.path+"?message=Successful")
	}
}

func (p *SignupPage) GetPagePath() string {
	return p.path
}

func (p *SignupPage) GetPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, p.template, echo.Map{
			"data": p.GetPageData(c),
		})
	}
}
