package page

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	userPageUri = "/user"
)

type UserPage struct {
	db *gorm.DB
}

type UserPageData struct {
}

func NewUserPage(db *gorm.DB) *Page {
	deps := &UserPage{
		db: db,
	}

	return &Page{
		MenuID:      "user-page",
		Title:       "User Page",
		Path:        userPageUri,
		Template:    "user",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *UserPage) GetPageData(c echo.Context) any {
	return UserPageData{}
}
