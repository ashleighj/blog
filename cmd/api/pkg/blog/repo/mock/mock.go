package mock

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"api/config"
	lg "api/logger"
	"api/pkg/blog/repo/entity"
)

var (
	errCast = errors.New("store contents do not match store struct")
)

type MockDatastore struct {
	filePath string
	data     store
}

type store struct {
	Articles []entity.Article
	// Contributors        []entity.Contributor
	// CustomContentTypes  []entity.CustomContentType
	// CustomContent       []entity.CustomContent
	// ContributorStatuses []entity.ContributorStatus
	// PublishStatuses     []entity.PublishStatus
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
	db.data, e = db.getStore(ctx)
	if e != nil {
		lg.Error(ctx, e)
	}

	return
}

// WriteReplace re-writes the contents of the store struct (above)
// into the appropriate mock datastore file
func (db *MockDatastore) WriteReplace(ctx context.Context) error {
	contents, e := json.MarshalIndent(db.data, "", "    ")
	if e != nil {
		lg.Error(ctx, e)
		return e
	}

	return ioutil.WriteFile(db.filePath, contents, 0755)
}

func (db *MockDatastore) getStore(ctx context.Context) (data store, e error) {
	path, e := filepath.Abs(db.filePath)
	if e != nil {
		log.Println(e)
		return
	}

	_, e = os.Stat(path)
	if e != nil {
		// if database file doesn't exist, create a new one.
		if os.IsNotExist(e) {
			lg.Warn(ctx, e)
			dir, _ := filepath.Split(path)

			e = os.MkdirAll(dir, 0775)
			if e != nil {
				lg.Error(ctx, e)
				return
			}

			db.data = store{}
			e = db.WriteReplace(ctx)
			if e != nil {
				lg.Error(ctx, e)
				return
			}
		}
	}

	storeFile, e := ioutil.ReadFile(path)
	lg.Info(ctx, path)
	if e != nil {
		lg.Error(ctx, e)
		return
	}

	e = json.Unmarshal(storeFile, &data)
	if e != nil {
		lg.Error(ctx, e)
		return
	}

	return
}

func (db *MockDatastore) Close(ctx context.Context) error {
	return nil
}
