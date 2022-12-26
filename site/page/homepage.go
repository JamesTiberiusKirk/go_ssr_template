package page

import (
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HomePage struct {
	path     string
	template string
	db       *gorm.DB
}

type HomePageData struct {
	Title string
}

func NewHomePage(db *gorm.DB) *Page {
	deps := &HomePage{
		db: db,
	}

	return &Page{
		Path:        "/",
		Template:    "homepage",
		Deps:        deps,
		GetPageData: deps.GetPageData,
	}
}

func (p *HomePage) GetPageData(c echo.Context) any {
	log.Print("Home page")
	return HomePageData{
		Title: "Hello world",
	}
}
