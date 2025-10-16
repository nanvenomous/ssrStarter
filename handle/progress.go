package handle

import "net/http"

func init() {
	setupFuncs = append(setupFuncs, func(mux *http.ServeMux) {
	})
}
