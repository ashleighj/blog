package mock

import (
	"api/config"
	lg "api/logger"
	"api/pkg/mock"
	err "api/service_error"
	"context"
	"errors"
)

var (
	errCast = errors.New("store contents do not match store struct")
)

type MockDatastore struct {
	filePath string
	data     store
}

type store struct {
}

// New creates, initialises and returns a new MockDatastore
func New(ctx context.Context,
	conf config.DataSource) (db *MockDatastore, e error) {

	db = &MockDatastore{filePath: conf.Host}
	e = db.Connect(ctx)
	return
}

// Connect attempts to initialise and establish a
// connection to the mock datastore
func (db *MockDatastore) Connect(ctx context.Context) (e error) {
	data, e := mock.GetStore(ctx, db.filePath)
	if e != nil {
		lg.Error(ctx, e)
		return
	}

	if data != nil {
		var ok bool
		if db.data, ok = data.(store); !ok {
			e = err.DatastoreInit(errCast)
			lg.Error(ctx, e)
			return
		}
	}

	return
}

// WriteReplace re-writes the contents of the store struct (above)
// into the appropriate mock datastore file
func (db *MockDatastore) WriteReplace(ctx context.Context) error {
	return mock.WriteReplace(ctx, db.filePath, db.data)
}

func (db *MockDatastore) Close(ctx context.Context) error {
	return nil
}
