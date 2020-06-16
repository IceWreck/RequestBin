package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/IceWreck/RequestBin/db"
)

// var database *sql.DB

//RequestView - Endpoint to make your requests here
func RequestView(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	h.Set("Host", r.Host)

	err := r.ParseForm()
	if err != nil {
		return
	}

	details := db.RequestDetails{
		Parameters: r.Form,
		Headers:    getRequestHeaders(r),
		Origin:     getOrigin(r),
		URL:        getURL(r).String(),
		Method:     r.Method,
	}

	response, err := json.MarshalIndent(details, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Store them

	// save := db.SavedJSON{
	// 	SavedJSONField: string(response),
	// }
	// db.CreateTable(database)
	db.AddRequests(string(response))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}

// The following functions have been shamelessly stolen (:P)
// from https://github.com/mccutchen/go-httpbin/ (MIT Licensed)

func getRequestHeaders(r *http.Request) http.Header {
	h := r.Header
	h.Set("Host", r.Host)
	return h
}

func getOrigin(r *http.Request) string {
	origin := r.Header.Get("X-Forwarded-For")
	if origin == "" {
		origin = r.RemoteAddr
	}
	return origin
}

func getURL(r *http.Request) *url.URL {
	scheme := r.Header.Get("X-Forwarded-Proto")
	if scheme == "" {
		scheme = r.Header.Get("X-Forwarded-Protocol")
	}
	if scheme == "" && r.Header.Get("X-Forwarded-Ssl") == "on" {
		scheme = "https"
	}
	if scheme == "" {
		scheme = "http"
	}

	host := r.URL.Host
	if host == "" {
		host = r.Host
	}

	return &url.URL{
		Scheme:     scheme,
		Opaque:     r.URL.Opaque,
		User:       r.URL.User,
		Host:       host,
		Path:       r.URL.Path,
		RawPath:    r.URL.RawPath,
		ForceQuery: r.URL.ForceQuery,
		RawQuery:   r.URL.RawQuery,
		Fragment:   r.URL.Fragment,
	}
}
