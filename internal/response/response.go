package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"go.uber.org/atomic"
)

// WebConfig Config
type WebConfig struct {
	Router string   `json:"router"`
	Resp   *WebResp `json:"resp"`
}

// WebResp WebResp
type WebResp struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
}

type Responser interface {
	Get(router string) (*WebConfig, error)
}

type fileResponse struct {
	path     string
	mustSync *atomic.Bool
	mu       sync.Mutex
	watcher  *fsnotify.Watcher
	data     []WebConfig
}

func NewFileResponse(path string) (Responser, error) {

	filePath, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	f := &fileResponse{
		path:     filePath,
		mustSync: atomic.NewBool(false),
	}

	if err := f.loadFile(); err != nil {
		return nil, err
	}
	f.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = f.watcher.Add(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "path: %s", filePath)
	}
	log.Printf("path: %s add file watcher\n", filePath)
	go f.watch()

	return f, nil
}

func (f *fileResponse) loadFile() error {

	body, err := ioutil.ReadFile(f.path)
	if err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(json.Unmarshal(body, &f.data))
}

func (f *fileResponse) Get(router string) (*WebConfig, error) {
	if f.mustSync.Load() {
		if err := f.loadFile(); err != nil {
			return nil, err
		}
		f.mustSync.Toggle()
	}
	for _, v := range f.data {
		if v.Router == router {
			return &v, nil
		}
	}

	return nil, errors.WithStack(fmt.Errorf("miss router"))
}

func (f *fileResponse) watch() {
	for {
		select {
		case e := <-f.watcher.Events:
			f.mustSync.CAS(false, true)
			log.Println("file", e)
		case err := <-f.watcher.Errors:
			log.Fatalln("failed to watch response file", err)
		}
	}
}
