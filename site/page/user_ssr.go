package page

import (
	"go_web_template/models"
	"go_web_template/session"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	userSsrPageUri = "/user-ssr"
)

type UserSSRPage struct {
	db      *gorm.DB
	session *session.Manager
}

type UserSSRPageData struct {
	User models.User
}

func NewUserSSRPage(db *gorm.DB, session *session.Manager) *Page {
	deps := &UserSSRPage{
		db:      db,
		session: session,
	}

	return &Page{
		MenuID:      "user-page",
		Title:       "User Page",
		Path:        userSsrPageUri,
		Template:    "user_ssr_example",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *UserSSRPage) GetPageData(c echo.Context) any {
	user := p.session.GetUser(c)

	dbUser := &models.User{}
	result := p.db.Where(&models.User{Email: user.Email}).Find(dbUser)
	if result.Error != nil {
		logrus.
			WithError(result.Error).
			Error("error getting user from the database")
		query := map[string]string{
			"error": internalServerError,
		}
		return redirect(c, loginPageUri, query)
	}
	return UserPageData{*dbUser}
}
