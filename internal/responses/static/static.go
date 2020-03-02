package static

import (
	"io/ioutil"
	"net/http"

	"github.com/haozibi/httpbee/internal/config"
	"github.com/haozibi/httpbee/internal/responses"
	"github.com/pkg/errors"
)

type staticResponse struct{}

func New() responses.Responser {

	return &staticResponse{}
}

func (b *staticResponse) Respond(w http.ResponseWriter, r *http.Request, resp *config.WebResp) error {
	for k, v := range resp.Headers {
		w.Header().Add(k, v)
	}

	if resp.Status != 0 {
		w.WriteHeader(resp.Status)
	}

	fileBody, err := ioutil.ReadFile(resp.File)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = w.Write(fileBody)
	return errors.WithStack(err)
}
