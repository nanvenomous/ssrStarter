package handle

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/amalfra/etag/v3"
	"github.com/nanvenomous/ssrStarter/ui"
)

var (
	embeddedResources fs.FS
)

func Setup(mux *http.ServeMux, buildFS embed.FS) (http.Handler, error) {
	var err error
	embeddedResources, err = fs.Sub(buildFS, "build")
	if err != nil {
		return mux, err
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains([]string{"", "/"}, r.URL.Path) {
			HomeHandler(w, r)
			return
		}
		serveResourceCachedETag(w, r, getBundledFile)
	})

	SetupAlert(mux)
	SetupHome(mux)
	SetupModal(mux)
	SetupThemeController(mux)

	return loggingMiddleware(mux), nil
}

func getBundledFile(r *http.Request) ([]byte, error) {
	if len(r.URL.Path) > 500 {
		return []byte{}, fmt.Errorf(
			"The filepath you entered: %s was suspiciously long",
			r.URL.Path,
		)
	}

	fl, err := embeddedResources.Open(strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil {
		return []byte{}, err
	}
	defer fl.Close()

	return io.ReadAll(fl)
}

func serveResourceCachedETag(w http.ResponseWriter, r *http.Request,
	fileCheck func(r *http.Request) ([]byte, error),
) {
	w.Header().Set("Cache-Control", "max-age=0")
	content, err := fileCheck(r)
	if err != nil {
		http.Error(w, "Could not locate the file to serve: "+err.Error(), http.StatusNotFound)
		return
	}

	etg := etag.Generate(string(content), false)
	if r.Header.Get("If-None-Match") == etg {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("ETag", etg)
	http.ServeContent(w, r, r.URL.Path, time.Unix(0, 0), bytes.NewReader(content))
}

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode    int
	headerWritten bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.headerWritten {
		return
	}
	rw.statusCode = code
	rw.headerWritten = true
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.headerWritten {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default to 200 if WriteHeader is never called
		}

		next.ServeHTTP(wrappedWriter, r)

		var color string
		switch {
		case wrappedWriter.statusCode >= 500:
			color = "\033[31m" // Red for 500s
		case wrappedWriter.statusCode >= 400:
			color = "\033[33m" // Yellow for 400s
		case wrappedWriter.statusCode >= 300:
			color = "\033[34m" // Blue for 300s
		case wrappedWriter.statusCode >= 200:
			color = "\033[32m" // Green for 200s
		default:
			color = "\033[0m" // Default color for other codes
		}

		if !slices.Contains([]int{http.StatusNotModified}, wrappedWriter.statusCode) {
			log.Printf(
				"[%s] %s %d %s %s",
				r.Method,
				color,
				wrappedWriter.statusCode,
				"\033[0m", // Reset color
				r.URL.String(),
			)
		}
	})
}

// render takes multiple templ components and writes their html to the handler output
// it returns an error if something fails as well as a net/http status code
func render(w http.ResponseWriter, r *http.Request, cmp ...templ.Component) (int, error) {
	err := ui.Multi(cmp...).Render(r.Context(), w)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("renderMulti: %v", err)
	}
	return http.StatusOK, nil
}

// errorHTTP calls http.Error
func errorHTTP(w http.ResponseWriter, status int, err error) {
	if err != nil {
		log.Println(err)
	}
	http.Error(w, err.Error(), status)
}
