package page

import (
	"github.com/labstack/echo/v4"
)

const (
	homePageUri = "/"
)

type HomePage struct {
}

type HomePageData struct {
}

func NewHomePage() *Page {
	deps := &HomePage{}

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
