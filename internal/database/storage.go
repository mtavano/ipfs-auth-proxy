package database

import (
	"sync"

	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("database: Storage.GetAdmin admin not found")

type Storage struct {
	adminUsers  *sync.Map
	adminTokens *sync.Map
	apiKeys     *sync.Map
}

func NewStorage(user, pass string) *Storage {
	var adminUsers sync.Map
	var adminTokens sync.Map
	var apiKeys sync.Map

	adminUsers.Store(user, &Admin{
		User:     user,
		password: pass,
		ApiKeys:  1,
	})

	apiKeys.Store("test", &ApiKey{
		Key:     "test",
		Enabled: true,
	})

	return &Storage{
		adminUsers:  &adminUsers,
		adminTokens: &adminTokens,
		apiKeys:     &apiKeys,
	}
}
