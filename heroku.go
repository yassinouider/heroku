package heroku

import "net/http"

func NewHeaders(r *http.Request) Headers {
	return Headers{
		For:       r.Header.Get("X-Forwarded-For"),
		Proto:     r.Header.Get("X-Forwarded-Proto"),
		Port:      r.Header.Get("X-Forwarded-Port"),
		Start:     r.Header.Get("X-Request-Start"),
		RequestID: r.Header.Get("X-Request-Id"),
		Via:       r.Header.Get("Via"),
	}
}

type Headers struct {
	For       string //the originating IP address of the client connecting to the Heroku router
	Proto     string //the originating protocol of the HTTP request (example: https)
	Port      string //the originating port of the HTTP request (example: 443)
	Start     string //unix timestamp (milliseconds) when the request was received by the router
	RequestID string //the Heroku HTTP Request ID
	Via       string //a code name for the Heroku router
}

func RedirectToHTTPS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := NewHeaders(r)
		if headers.Proto != "https" {
			http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}
