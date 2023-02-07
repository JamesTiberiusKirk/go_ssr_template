package route

import (
	"net/http"

	"github.com/JamesTiberiusKirk/go_web_template/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	userAPIRoute = "/:email"
)

// UserRoute user route dependency struct.
type UserRoute struct {
	db *gorm.DB
}

// NewUserRoute struct instance.
func NewUserRoute(db *gorm.DB) *Route {
	depts := &UserRoute{
		db: db,
	}

	return &Route{
		RouteID:    "user",
		Path:       userAPIRoute,
		Depts:      depts,
		GetHandler: depts.GetUser,
	}
}

func (r *UserRoute) GetUser(c echo.Context) error {
	userEmail := c.Param("email")
	if userEmail == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user := models.User{}
	result := r.db.Where(models.User{Email: userEmail}).Find(&user)
	if result.Error != nil {
		_ = c.NoContent(http.StatusInternalServerError)
		return result.Error
	}

	return c.JSON(http.StatusOK, user)
}
