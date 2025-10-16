package handle

import (
	"net/http"
	"time"

	"github.com/nanvenomous/ssrStarter/ui"
)

func init() {
	setupFuncs = append(setupFuncs, func(mux *http.ServeMux) {

		mux.HandleFunc("/progress", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				ProgressHandler(w, r)
				return
			case http.MethodPost:
				// Simulate some work by sleeping for 4 seconds
				time.Sleep(4 * time.Second)

				// Return the button back to original state and send alert out-of-band
				stts, err := render(w, r,
					ui.ProgressButton(),
					ui.Alert(ui.PropsAlert{
						Label: "Task completed successfully after 4 seconds!",
						Type:  ui.AlertTypeSuccess,
					}),
				)
				if err != nil {
					errorHTTP(w, stts, err)
				}
				return
			}
		})

	})
}

func ProgressHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := render(w, r,
		ui.PageProgress(),
	)
	if err != nil {
		errorHTTP(w, stts, err)
	}
}
