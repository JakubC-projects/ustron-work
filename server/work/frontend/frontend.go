package frontend

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(location string) http.Handler {
	parsed, err := url.Parse(location)
	if err != nil {
		panic(err)
	}
	if parsed.Scheme == "" {
		return getFrontendFileHandler(location)
	}
	return getFrontendProxyHandler(parsed)
}

func getFrontendProxyHandler(remote *url.URL) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(r *http.Request) {
			r.Header = req.Header
			r.Host = remote.Host
			r.URL.Scheme = remote.Scheme
			r.URL.Host = remote.Host
		}
		proxy.ServeHTTP(w, req)
	})
}

func getFrontendFileHandler(path string) http.Handler {
	mux := http.NewServeMux()

	indexPath, err := url.JoinPath(path, "index.html")
	if err != nil {
		panic(fmt.Errorf("cannot get index path: %w", err))
	}

	mux.Handle("/", NoCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, indexPath)
	})))

	assetsPath, err := url.JoinPath(path, "assets")
	if err != nil {
		panic(fmt.Errorf("cannot get assets path: %w", err))
	}
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir(assetsPath))))

	return mux
}
