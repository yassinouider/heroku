# Heroku [![GoDoc](http://godoc.org/github.com/yassinouider/heroku?status.svg)](http://godoc.org/github.com/yassinouider/heroku)

Package heroku read http headers add by heroku HTTP Routing

## Installation

```bash
$ go get github.com/yassinouider/heroku
```

## Usage

Show ip and more info from heroku HTTP Routing
```go
func ShowIP(w http.ResponseWriter, r *http.Request) {
  headers := heroku.NewHeaders(r)
  ip := headers.For

  fmt.Fprintf(w, "ip: %s", ip)
}
```

Middleware force redirect http to https
```go
http.Handle("/" heroku.RedirectToHTTPS(myHandler))
```
