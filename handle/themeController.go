package handle

import (
	"net/http"

	"github.com/nanvenomous/ssrStarter/ui"
)

func init() {
	setupFuncs = append(setupFuncs, func(mux *http.ServeMux) {

		mux.HandleFunc("/theme", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				ThemeHandler(w, r)
				return
			}
		})

	})
}

func ThemeHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := render(w, r,
		ui.PageTheme(),
	)
	if err != nil {
		errorHTTP(w, stts, err)
	}
}
