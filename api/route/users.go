package route

import (
	"net/http"

	"github.com/JamesTiberiusKirk/go_web_template/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	usersAPIRoute = "/users"
)

// UserRoute user route dependency struct.
type UsersRoute struct {
	db *gorm.DB
}

// NewUsersRoute struct instance.
func NewUsersRoute(db *gorm.DB, sr *Route) *Route {
	depts := &UsersRoute{
		db: db,
	}

	return &Route{
		SubRoute:   sr,
		RouteID:    "users",
		Path:       usersAPIRoute,
		Depts:      depts,
		GetHandler: depts.GetUsers,
	}
}

type GetHandlerResponse struct {
	Users []models.User `json:"users"`
}

func (r *UsersRoute) GetUsers(c echo.Context) error {
	users := []models.User{}
	result := r.db.Find(&users)
	if result.Error != nil {
		_ = c.NoContent(http.StatusInternalServerError)
		return result.Error
	}

	return c.JSON(http.StatusOK, users)
}
