package route

import (
	"go_ssr_template/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	userApiRoute  = "/:email"
	usersApiRoute = "/user"
)

// UserRoute user route dependency struct
type UserRoute struct {
	db *gorm.DB
}

// NewUserRoute struct instance
func NewUserRoute(db *gorm.DB) *Route {
	depts := &UserRoute{
		db: db,
	}

	return &Route{
		SubRoute: &Route{
			Depts:      depts,
			Path:       userApiRoute,
			GetHandler: depts.GetUser,
		},
		Path:       usersApiRoute,
		Depts:      depts,
		GetHandler: depts.GetUsers,
	}
}

type GetHandlerResponse struct {
	Users []models.User `json:"users"`
}

func (r *UserRoute) GetUser(c echo.Context) error {
	userEmail := c.Param("email")
	if userEmail == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	user := models.User{}
	result := r.db.Where(models.User{Email: userEmail}).Find(&user)
	if result.Error != nil {
		c.NoContent(http.StatusInternalServerError)
		return result.Error
	}

	return c.JSON(http.StatusOK, user)
}

func (r *UserRoute) GetUsers(c echo.Context) error {
	users := []models.User{}
	result := r.db.Find(&users)
	if result.Error != nil {
		c.NoContent(http.StatusInternalServerError)
		return result.Error
	}

	return c.JSON(http.StatusOK, users)
}
