package mock

import (
	lg "api/logger"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetStore(ctx context.Context, file string) (data interface{}, e error) {
	path, e := filepath.Abs(file)
	if e != nil {
		log.Println(e)
		return
	}

	_, e = os.Stat(path)
	if e != nil {
		// if database file doesn't exist, create a new one.
		if os.IsNotExist(e) {
			lg.Error(ctx, e)
			dir, _ := filepath.Split(path)

			e = os.MkdirAll(dir, 0775)
			if e != nil {
				lg.Error(ctx, e)
				return
			}

			e = WriteReplace(ctx, file, data)
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

// WriteReplace re-writes the contents of the store struct (above)
// into the appropriate mock datastore file
func WriteReplace(ctx context.Context, file string, data interface{}) error {
	if data == nil {
		data = []interface{}{}
	}

	contents, e := json.MarshalIndent(data, "", "    ")
	if e != nil {
		lg.Error(ctx, e)
		return e
	}

	return ioutil.WriteFile(file, contents, 0755)
}

