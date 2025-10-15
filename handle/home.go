package handle

import (
	"net/http"
	"strconv"
	"sync"

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

var (
	counterMutex sync.Mutex
	counterValue int = 0
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	operation := r.URL.Query().Get("op")
	if operation != "inc" && operation != "dec" {
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	counterMutex.Lock()
	if operation == "inc" {
		counterValue++
	} else {
		counterValue--
	}
	currentValue := counterValue
	counterMutex.Unlock()

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(strconv.Itoa(currentValue)))
}

func SetupHome(mux *http.ServeMux) {
	mux.HandleFunc("/counter", CounterHandler)
}
