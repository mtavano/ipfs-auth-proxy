package app

import (
	"context"
	"io"

	"github.com/ipfs/interface-go-ipfs-core/path"

	"github.com/mtavano/ipfs-auth-proxy/internal/database"
	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

type Context struct {
	Storage     *database.Storage
	IDGenerator IDGenerator
	BlockGetter BlockGetter
}

type IDGenerator interface {
	Generate() (string, error)
}

type BlockGetter interface {
	Get(context.Context, path.Path) (io.Reader, error)
}

func NewContext(storage *database.Storage, blockGetter BlockGetter) (*Context, error) {
	idGenerator, err := shortid.New(1, "-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_", 2342)
	if err != nil {
		return nil, errors.Wrap(err, "app: NewContext shortid.New error")

	}

	return &Context{
		Storage:     storage,
		IDGenerator: idGenerator,
		BlockGetter: blockGetter,
	}, nil
}

type MockBlockGetter struct{}

func (m *MockBlockGetter) Get() (io.Reader, error) {
	return nil, nil
}
