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
		MenuID:      "homePage",
		Title:       "Home Page",
		Frame:       true,
		Path:        homePageUri,
		Template:    "homepage/homepage.gohtml",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *HomePage) GetPageData(c echo.Context) any {
	return HomePageData{}
}
