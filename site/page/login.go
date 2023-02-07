package page

import (
	"net/http"

	"github.com/JamesTiberiusKirk/go_web_template/models"
	"github.com/JamesTiberiusKirk/go_web_template/session"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	loginPageURI = "/login"
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
		MenuID:      "loginPage",
		Title:       "Login",
		Frame:       true,
		Path:        loginPageURI,
		Template:    "login.gohtml",
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
		return c.Redirect(http.StatusSeeOther, loginPageURI+"?error=need username and password")
	}

	dbUser := &models.User{}
	result := p.db.Where(&models.User{Email: submitUser.Email}).Find(dbUser)
	if result.Error != nil {
		logrus.
			WithError(result.Error).
			Error("error getting user from the database")
		return c.Redirect(http.StatusSeeOther, loginPageURI+"?error=internal server problem")
	}
	if dbUser.Password == "" {
		return c.Redirect(http.StatusSeeOther, loginPageURI+"?error=wrong user name or password")
	}

	comparison, err := dbUser.ComparePassword(submitUser.Password)
	if err != nil {
		logrus.
			WithError(err).
			Error("error comparing passwords")
		return c.Redirect(http.StatusSeeOther, loginPageURI+"?error=internal server problem")
	}

	if !comparison {
		return c.Redirect(http.StatusSeeOther, loginPageURI+"?error=wrong user name or password")
	}

	p.sessionManager.InitSession(*dbUser, c)
	return c.Redirect(http.StatusSeeOther, userSsrPageURI)
}
