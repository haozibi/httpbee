package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/haozibi/httpbee/internal/response"
)

// FakeServer fake server
type FakeServer struct {
	cfg       *Config
	responser response.Responser
}

func NewFakeServer(cfg Config) (*FakeServer, error) {
	r, err := response.NewFileResponse(cfg.RespPath)
	if err != nil {
		return nil, err
	}
	return &FakeServer{
		cfg:       &cfg,
		responser: r,
	}, nil
}

// RunHTTP run http
func (f *FakeServer) RunHTTP() error {
	fmt.Printf(logo, BuildVersion)

	addr := fmt.Sprintf(":%d", f.cfg.Port)

	log.Println("listen and serve:", addr)
	if err := http.ListenAndServe(addr, f); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (f *FakeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	c, err := f.responser.Get(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Path: %s, Error: %v", path, err)
		return
	}
	f.serve(w, r, c.Resp)
}

func (f *FakeServer) serve(w http.ResponseWriter, r *http.Request, c *response.WebResp) {
	if c == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Path: %s, Error: %v", r.URL.Path, fmt.Errorf("miss response"))
		return
	}

	for k, v := range c.Headers {
		w.Header().Add(k, v)
	}

	ww := w.Header().Clone()
	spew.Dump(ww)

	w.WriteHeader(c.Status)
	if c.Body != nil {
		fmt.Fprintf(w, "%v", c.Body)
	}
}
