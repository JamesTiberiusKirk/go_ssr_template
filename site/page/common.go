package page

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func redirect(c echo.Context, uri string, query map[string]string) error {
	if len(query) > 0 {
		withQuery := fmt.Sprintf("%s?", uri)

		for k, v := range query {
			withQuery = fmt.Sprintf("%s%s=%s&", withQuery, k, v)
		}

		return c.Redirect(http.StatusSeeOther, withQuery)
	}

	return c.Redirect(http.StatusSeeOther, uri)
}
