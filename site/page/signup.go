package page

import (
	"fmt"
	"go_ssr_template/models"
	"go_ssr_template/session"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	signupPageUri = "/signup"
)

type SignupPage struct {
	db             *gorm.DB
	sessionManager *session.Manager
}

type SignupPageData struct {
	Message    string
	Validation struct {
		models.User
		RepeatPassword string
	}
	Misc string
}

func NewSignupPage(db *gorm.DB, sessionManager *session.Manager) *Page {
	deps := &SignupPage{
		db:             db,
		sessionManager: sessionManager,
	}

	return &Page{
		MenuID:         "signup-page",
		Title:          "Signup",
		Path:           "/signup",
		Template:       "signup",
		Deps:           deps,
		GetPageData:    deps.GetPageData,
		GetPostHandler: deps.GetPostHandler(),
	}
}

func (p *SignupPage) GetPageData(c echo.Context) any {
	data := SignupPageData{
		Message: c.QueryParam("message"),
	}

	data.Validation.Email = c.QueryParam("email")
	data.Validation.Username = c.QueryParam("username")
	data.Validation.Password = c.QueryParam("password")
	data.Validation.RepeatPassword = c.QueryParam("repeat_password")

	data.Misc = fmt.Sprintf("%+v", data.Validation)

	return data
}

func (p *SignupPage) GetPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := struct {
			models.User
			RepeatPassword string
		}{
			User: models.User{
				Email:    c.FormValue("email"),
				Username: c.FormValue("username"),
				Password: c.FormValue("password"),
			},
			RepeatPassword: c.FormValue("repeat_password"),
		}

		// TODO: use the redirect() function
		// query := map[string]string{}

		notPassed, err := user.Validate()
		if err != nil {
			logrus.
				WithError(err).
				Error("Error validating signup form")
			return c.Redirect(http.StatusSeeOther, signupPageUri+"?error=internal server error")
		}

		if user.Password != user.RepeatPassword {
			notPassed = append(notPassed, "repeat_password")
		}

		log.Printf("%+v", notPassed)

		if len(notPassed) > 0 {
			query := "?error=Unable to validate data"
			for _, fields := range notPassed {
				query = fmt.Sprintf("%s&%s=not valid", query, fields)
			}
			return c.Redirect(http.StatusSeeOther, signupPageUri+query)
		}

		user.SetPassword(user.Password)

		result := p.db.WithContext(c.Request().Context()).Create(&user)
		if result.Error != nil {
			msg := "failed to insert user into db"
			logrus.
				WithError(result.Error).
				Error(msg)

			redirectURL := fmt.Sprintf("%s?error=Internal server error", signupPageUri)
			c.Redirect(http.StatusSeeOther, redirectURL)
		}

		p.sessionManager.InitSession(user.Email, user.ID, c)
		return c.Redirect(http.StatusSeeOther, homePageUri)
	}
}
