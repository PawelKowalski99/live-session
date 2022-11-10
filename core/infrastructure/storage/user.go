package storage

import (
	"context"
	"live-session-task/core/entities"
	"live-session-task/core/infrastructure/storage/user"
)

type User interface {
	CreateUser(ctx context.Context, name string) (entities.User, error)
	GetUser(ctx context.Context, id int32) (entities.User, error)
}

type UserQueries struct {
	queries *user.Queries
}

func NewUserQueries(queries *user.Queries) *UserQueries {
	return &UserQueries{queries: queries}
}

func (u *UserQueries) CreateUser(ctx context.Context, name string) (entities.User, error) {
	usr, err := u.queries.CreateUser(ctx, name)
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:   int(usr.ID),
		Name: usr.Name,
	}, err
}

func (u *UserQueries) GetUser(ctx context.Context, id int32) (entities.User, error) {
	usr, err := u.queries.GetUser(ctx, id)
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:   int(usr.ID),
		Name: usr.Name,
	}, err
}
