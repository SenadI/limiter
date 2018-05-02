package http

import (
	"net/http"
	"time"

	"github.com/senadi/limiter"
	json "github.com/senadi/limiter/pkg/json"
)

// greeting - contains basic information about the project and associcated HTTP server uptime.
type greeting struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Uptime      string `json:"uptime,omitempty"`
}

// greetingsHandler returns information about the service status and
// demonstrates the rate limiting capabilities. In the "real-world"
// there would be no interest in limiting the status handler.
func greetingsHandler(startTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := greeting{
			Name:        limiter.Name,
			Description: limiter.Description,
			Uptime:      time.Now().UTC().Sub(startTime).String(),
		}
		json.WriteJSON(status, w)
	}
}
