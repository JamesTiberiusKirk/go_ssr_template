package page

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	homePageUri = "/"
)

type HomePage struct {
	db *gorm.DB
}

type HomePageData struct {
}

func NewHomePage(db *gorm.DB) *Page {
	deps := &HomePage{
		db: db,
	}

	return &Page{
		MenuID:      "home-page",
		Title:       "Home Page",
		Path:        homePageUri,
		Template:    "homepage",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *HomePage) GetPageData(c echo.Context) any {
	return HomePageData{}
}
