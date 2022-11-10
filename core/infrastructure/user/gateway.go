package user

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"live-session-task/core/entities"
	"live-session-task/core/infrastructure/storage"
	"net/http"
	"strconv"

	"live-session-task/internal/cache"
)

// Gateway for access to EstateStorage, crawler and cache
type Gateway interface {
	GetUser(id int) (entities.User, error, int)
}

// Logic Domain
type Logic struct {
	queries storage.User
	cache   cache.Cache
}

// GetUser ...
func (t *Logic) GetUser(id int) (entities.User, error, int) {
	ctx := context.Background()
	var usr entities.User

	usrJson, err := t.cache.Get(ctx, "user-"+strconv.Itoa(id))
	if err != nil && err != cache.NotExistError {

		return entities.User{}, err, http.StatusBadRequest
	} else if err == nil {

		err = json.Unmarshal([]byte(usrJson), &usr)
		if err != nil {

			logrus.Fatalf("Failed to JSON Unmarshal")
			return usr, err, http.StatusBadRequest
		}
		return usr, nil, http.StatusOK

	}

	usrDb, err := t.queries.GetUser(context.Background(), int32(id))
	if err != nil {
		return entities.User{}, err, http.StatusBadRequest
	}

	usr = entities.User{
		ID:   usrDb.ID,
		Name: usrDb.Name,
	}

	usrStr, err := json.Marshal(usr)
	if err != nil {

		return entities.User{}, err, http.StatusBadRequest
	}
	err = t.cache.Set(ctx, "user-"+strconv.Itoa(id), string(usrStr))
	if err != nil {

		return entities.User{}, err, http.StatusForbidden
	}

	return usr, nil, http.StatusOK

}

// Constructor
func NewLogic(
	queries storage.User,
	cache cache.Cache,

) *Logic {
	return &Logic{
		queries: queries,
		cache:   cache,
	}
}
