package user_test

import (
	"context"
	"github.com/jellydator/ttlcache/v3"
	"github.com/stretchr/testify/assert"
	"live-session-task/core/entities"
	"live-session-task/core/infrastructure/user"
	"live-session-task/internal/cache"
	"net/http"
	"sync"
	"testing"
	"time"
)

type cacheMock struct {
	cache *ttlcache.Cache[string, string]
	mu    *sync.Mutex
}

type User struct {
	Name string `json:"name,omitempty"`
}

type dbMock struct {
	users map[int]User
	mu    *sync.Mutex
}

func (m *cacheMock) Get(ctx context.Context, id string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var err error

	item := m.cache.Get(id)

	if item == nil {
		return "", cache.NotExistError
	}

	return item.Value(), err
}

func (m *cacheMock) Set(ctx context.Context, id, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache.Set(id, name, ttlcache.DefaultTTL)

	return nil
}

func (m *dbMock) CreateUser(ctx context.Context, name string) (entities.User, error) {
	return entities.User{}, nil
}

func (m *dbMock) GetUser(ctx context.Context, id int32) (entities.User, error) {
	m.mu.Lock()
	// Dummy sleeper because db is slower than cache.
	time.Sleep(800 * time.Millisecond)

	defer m.mu.Unlock()

	usr := m.users[int(id)]

	return entities.User{
		ID:   int(id),
		Name: usr.Name,
	}, nil
}

func TestLogic_GetUser(t *testing.T) {
	var lc = user.NewLogic(&dbMock{
		users: map[int]User{
			0: {Name: "a"},
			1: {Name: "b"},
			2: {Name: "c"},
			3: {Name: "d"},
			4: {Name: "e"},
			5: {Name: "f"},
		},
		mu: &sync.Mutex{},
	}, &cacheMock{
		cache: ttlcache.New[string, string](
			ttlcache.WithTTL[string, string](2 * time.Second)),
		mu: &sync.Mutex{},
	})

	type Out struct {
		user   entities.User
		err    error
		status int
	}

	cases := []struct {
		name    string
		in      int
		out     Out
		wantErr bool
	}{
		{
			name: "0",
			in:   0,
			out: Out{
				user:   entities.User{ID: 0, Name: "a"},
				err:    nil,
				status: http.StatusOK,
			},
			wantErr: false,
		},
		{
			name: "0 from cache",
			in:   0,
			out: Out{
				user:   entities.User{ID: 0, Name: "a"},
				err:    nil,
				status: http.StatusOK,
			},
			wantErr: false,
		},
	}

	for _, tc := range cases {
		userRsp, err, status := lc.GetUser(tc.in)
		if err != nil && !tc.wantErr {

		}

		assert.Equal(t, tc.out.user, userRsp)
		assert.Equal(t, tc.out.status, status)

	}
}

func BenchmarkLogic_GetUser(b *testing.B) {
	var lc = user.NewLogic(&dbMock{
		users: map[int]User{
			0: {Name: "a"},
			1: {Name: "b"},
			2: {Name: "c"},
			3: {Name: "d"},
			4: {Name: "e"},
			5: {Name: "f"},
		},
		mu: &sync.Mutex{},
	}, &cacheMock{
		cache: ttlcache.New[string, string](
			ttlcache.WithTTL[string, string](1 * time.Second)),
		mu: &sync.Mutex{},
	})

	for i := 0; i < b.N; i++ {
		lc.GetUser(0)

	}
}
