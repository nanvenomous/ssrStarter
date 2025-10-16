package handle

import (
	"net/http"

	"github.com/nanvenomous/ssrStarter/ui"
)

func init() {
	setupFuncs = append(setupFuncs, func(mux *http.ServeMux) {

		mux.HandleFunc("/alert", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				AlertHandler(w, r)
				return
			case http.MethodPost:
				// Get the alert type from query parameter, default to info
				alertType := r.URL.Query().Get("type")
				if alertType == "" {
					alertType = "info"
				}

				var alert ui.AlertType
				var message string

				switch alertType {
				case "success":
					alert = ui.AlertTypeSuccess
					message = "Success! Your action was completed successfully."
				case "warning":
					alert = ui.AlertTypeWarning
					message = "Warning: Please review your action carefully."
				case "error":
					alert = ui.AlertTypeError
					message = "Error: Something went wrong with your request."
				default: // info
					alert = ui.AlertTypeInfo
					message = "Info: This is an informational message."
				}

				// Create a button with "clicked" state and an alert
				stts, err := render(w, r,
					ui.AlertTestButton(ui.PropsAlertTestButton{
						Type:    alertType,
						Clicked: true,
					}),
					ui.Alert(ui.PropsAlert{
						Label: message,
						Type:  alert,
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

func AlertHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := render(w, r,
		ui.PageAlert(),
	)
	if err != nil {
		errorHTTP(w, stts, err)
	}
}
