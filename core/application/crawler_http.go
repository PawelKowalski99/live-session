package app_user

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"live-session-task/core/infrastructure/storage"
	sqlc "live-session-task/core/infrastructure/storage/user"
	"live-session-task/core/infrastructure/user"
	"net/http"
	"strconv"

	//	"real-user/core/infrastructure/storage"
	"live-session-task/internal/cache"
	//"real-user/internal/crawler"
)

// The HTTP Handler for User
type UserHttpService struct {
	gtw user.Gateway
}

func (t *UserHttpService) Get(c echo.Context) (err error) {

	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err, i := t.gtw.GetUser(id)
	if err != nil {
		return c.JSON(i, err.Error())
	}

	return c.JSON(i, user)
}

// Constructor
func NewCrawlerHttpService(ctx context.Context,
	db *sql.DB,
	cache cache.Cache,

) *UserHttpService {
	return &UserHttpService{
		gtw: user.NewLogic(
			storage.NewUserQueries(sqlc.New(db)),
			cache,
		),
	}
}
