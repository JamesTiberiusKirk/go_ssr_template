package page

import (
	"fmt"

	"github.com/JamesTiberiusKirk/go_web_template/models"
	"github.com/JamesTiberiusKirk/go_web_template/session"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	signupPageURI = "/signup"
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
		MenuID:      "signupPage",
		Title:       "Signup",
		Frame:       true,
		Path:        "/signup",
		Template:    "signup.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
		PostHandler: deps.PostHandler,
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

func (p *SignupPage) PostHandler(c echo.Context) error {
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

	query := map[string]string{}

	notPassed, err := user.Validate()
	if err != nil {
		logrus.
			WithError(err).
			Error("Error validating signup form")
		query["error"] = internalServerError
		return redirect(c, signupPageURI, query)
	}

	if user.Password != user.RepeatPassword {
		notPassed = append(notPassed, "repeat_password")
	}

	if len(notPassed) > 0 {
		query["error"] = invalidData
		for _, fields := range notPassed {
			query[fields] = "not valid"
		}
		return redirect(c, signupPageURI, query)
	}

	err = user.SetPassword(user.Password)
	if err != nil {
		query["error"] = internalServerError
		return redirect(c, signupPageURI, query)
	}

	result := p.db.WithContext(c.Request().Context()).Create(&user.User)
	if result.Error != nil {
		msg := "failed to insert user into db"
		logrus.
			WithError(result.Error).
			Error(msg)

		query["error"] = internalServerError
		return redirect(c, signupPageURI, query)
	}

	p.sessionManager.InitSession(user.User, c)
	return redirect(c, userPageURI, nil)
}
