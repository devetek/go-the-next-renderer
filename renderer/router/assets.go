package router

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func AssetsHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("r.URL.Path:", r.URL.Path)

	originServerURL, err := url.Parse("http://127.0.0.1:3000")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = originServerURL.Host
	r.URL.Host = originServerURL.Host
	r.URL.Scheme = originServerURL.Scheme

	proxy := httputil.NewSingleHostReverseProxy(originServerURL)

	proxy.ServeHTTP(w, r)
}
