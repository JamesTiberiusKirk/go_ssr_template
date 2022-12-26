package page

import (
	"github.com/labstack/echo/v4"
)

type Page struct {
	QueryError string
}

type PageInterface interface {
	GetPageData(c echo.Context) any
	GetPagePath() string
	GetPageHandler() echo.HandlerFunc
	GetPostHandler() echo.HandlerFunc
}

type FramePage struct {
	MenuId  string
	Content PageInterface
}

type FramePageData struct {
	MenuGroupId string
	MenuId      string
	Template    string
	PageData    any
}
