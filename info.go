package version

import (
	"encoding/json"
	"net/http"
	"runtime"
)

// Info is information about the version of the application.
type Info struct {
	Version  string         `json:"version"`
	Commit   string         `json:"commit"`
	Date     string         `json:"build_date"`
	Go       string         `json:"go"`
	Metadata map[string]any `json:"metadata,omitempty"`
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

// New initializes a new Info.
func New(version, commit, date string, opts ...Option) *Info {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	return &Info{
		Version:  version,
		Commit:   commit,
		Date:     date,
		Go:       runtime.Version(),
		Metadata: o.metadata,
	}
}

type options struct {
	metadata map[string]any
}

// Option is used to customize the version info.
type Option func(*options)

// WithMetadata sets the metadata.
func WithMetadata(data map[string]any) Option {
	return func(o *options) {
		o.metadata = data
	}
}
