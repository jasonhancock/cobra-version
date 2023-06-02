package version

import (
	"encoding/json"
	"net/http"
)

// Info is information about the version of the application.
type Info struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"build_date"`
	Go      string `json:"go"`
}

// HTTPHandlerFunc returns a HandlerFunc that will output the version information
// in JSON format.
func (i Info) HTTPHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(i, "", "\t")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
