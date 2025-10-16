package handle

import (
	"net/http"

	"github.com/nanvenomous/ssrStarter/ui"
)

func init() {
	setupFuncs = append(setupFuncs, func(mux *http.ServeMux) {

		mux.HandleFunc("/modal", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				ModalHandler(w, r)
				return
			case http.MethodPut:
				stts, err := render(w, r,
					ui.ExampleModal(),
				)
				if err != nil {
					errorHTTP(w, stts, err)
				}
				return
			case http.MethodDelete:
				stts, err := render(w, r,
					ui.EmptyModalPopover(ui.PropsEmptyModalPopover{}),
				)
				if err != nil {
					errorHTTP(w, stts, err)
				}
				return
			}
		})

	})
}

func ModalHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := render(w, r,
		ui.PageModal(),
	)
	if err != nil {
		errorHTTP(w, stts, err)
	}
}
