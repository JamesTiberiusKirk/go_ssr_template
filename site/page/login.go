package page

import (
	"go_ssr_template/models"
	"go_ssr_template/session"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	loginPageUri = "/login"
)

type LoginPage struct {
	db             *gorm.DB
	sessionManager *session.Manager
}

type LoginPageData struct {
}

func NewLoginPage(db *gorm.DB, sessionManager *session.Manager) *Page {
	deps := &LoginPage{
		db:             db,
		sessionManager: sessionManager,
	}

	return &Page{
		MenuID:      "login-page",
		Title:       "Login",
		Path:        loginPageUri,
		Template:    "login",
		Deps:        deps,
		GetPageData: deps.GetPageData,
		PostHandler: deps.PostHandler,
	}
}

func (p *LoginPage) GetPageData(c echo.Context) any {
	return LoginPageData{}
}

func (p *LoginPage) PostHandler(c echo.Context) error {
	submitUser := models.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	if submitUser.Password == "" || submitUser.Email == "" {
		return c.Redirect(http.StatusSeeOther, loginPageUri+"?error=need username and password")
	}

	dbUser := &models.User{}
	result := p.db.Where(&models.User{Email: submitUser.Email}).Find(dbUser)
	if result.Error != nil {
		logrus.
			WithError(result.Error).
			Error("error getting user from the database")
		return c.Redirect(http.StatusSeeOther, loginPageUri+"?error=internal server problem")
	}

	if dbUser.Password == "" {
		return c.Redirect(http.StatusSeeOther, loginPageUri+"?error=wrong user name or password")
	}

	comparison, err := dbUser.ComparePassword(submitUser.Password)
	if err != nil {
		logrus.
			WithError(err).
			Error("error comparing passwords")
		return c.Redirect(http.StatusSeeOther, loginPageUri+"?error=internal server problem")
	}

	if !comparison {
		return c.Redirect(http.StatusSeeOther, loginPageUri+"?error=wrong user name or password")
	}

	p.sessionManager.InitSession(dbUser.Email, dbUser.ID, c)
	return c.Redirect(http.StatusSeeOther, homePageUri)
}
