package page

import (
	"go_web_template/models"
	"go_web_template/session"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	userPageUri = "/user"
)

type UserPage struct {
	db      *gorm.DB
	session *session.Manager
}

// TsType
type UserPageData struct {
	User models.User `json:"user"`
}

func NewUserPage(db *gorm.DB, session *session.Manager) *Page {
	deps := &UserPage{
		db:      db,
		session: session,
	}

	return &Page{
		MenuID:      "userPage",
		Title:       "User Page",
		Frame:       true,
		Path:        userPageUri,
		Template:    "user/user.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *UserPage) GetPageData(c echo.Context) any {
	user, err := p.session.GetUser(c)
	if err != nil {
		query := map[string]string{
			"error": internalServerError,
		}
		redirect(c, loginPageUri, query)
		return err
	}

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
