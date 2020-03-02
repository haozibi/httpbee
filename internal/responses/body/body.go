package body

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/haozibi/httpbee/internal/config"
	"github.com/haozibi/httpbee/internal/responses"
)

type bodyResponse struct{}

// New new body response
func New() responses.Responser {
	return &bodyResponse{}
}

func (b *bodyResponse) Respond(w http.ResponseWriter, r *http.Request, resp *config.WebResp) error {

	for k, v := range resp.Headers {
		w.Header().Add(k, v)
	}

	if resp.Status != 0 {
		w.WriteHeader(resp.Status)
	}

	body := string(resp.Body)
	if strings.HasPrefix(body, `"`) {
		body = strings.TrimSuffix(body, `"`)
		body = strings.TrimPrefix(body, `"`)
	}

	if resp.Body != nil {
		fmt.Fprintf(w, "%v", body)
	}

	return nil
}
