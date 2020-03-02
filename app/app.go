package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/haozibi/httpbee/internal/config"
	"github.com/haozibi/httpbee/internal/responses"
	"github.com/haozibi/httpbee/internal/responses/body"
	"github.com/haozibi/httpbee/internal/responses/static"
)

// FakeServer fake server
type FakeServer struct {
	cfg    *Config
	config config.Configer
}

func NewFakeServer(cfg Config) (*FakeServer, error) {
	r, err := config.NewConfiger(cfg.RespPath)
	if err != nil {
		return nil, err
	}
	return &FakeServer{
		cfg:    &cfg,
		config: r,
	}, nil
}

// RunHTTP run http
func (f *FakeServer) RunHTTP() error {

	addr := fmt.Sprintf("0.0.0.0:%d", f.cfg.Port)

	log.Println("[bee] listen and serve:", addr)
	if err := http.ListenAndServe(addr, f); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (f *FakeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	c, err := f.config.Get(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Path: %s, Error: %v", path, err)
		return
	}

	var resp responses.Responser

	switch c.Type {
	case "", "body":
		resp = body.New()
	case "static":
		resp = static.New()
	default:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Path: %s, Error: type error", path)
		return
	}

	resp.Respond(w, r, c.Resp)
}
