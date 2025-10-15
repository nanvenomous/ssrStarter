package handle

import (
	"net/http"

	"github.com/nanvenomous/ssrStarter/ui"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := render(w, r,
		ui.Home(),
	)
	if err != nil {
		errorHTTP(w, stts, err)
	}
}
