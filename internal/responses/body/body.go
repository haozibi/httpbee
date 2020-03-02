package body

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/haozibi/httpbee/internal/config"
	"github.com/haozibi/httpbee/internal/responses"
	"github.com/pkg/errors"
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
		fmt.Fprintf(w, "%v", body)
		return nil
	}

	if resp.Body != nil {
		body, err := prettyPrint(resp.Body)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%v", string(body))
	}

	return nil
}

func prettyPrint(body json.RawMessage) ([]byte, error) {

	var m map[string]interface{}

	err := json.Unmarshal([]byte(body), &m)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	b, err := json.Marshal(m)
	return b, errors.WithStack(err)
}
