package responses

import (
	"net/http"

	"github.com/haozibi/httpbee/internal/config"
)

type Responser interface {
	Respond(w http.ResponseWriter, r *http.Request, resp *config.WebResp) error
}
