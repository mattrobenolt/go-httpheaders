//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
	"time"
)

var commonHeaders = []string{
	"Accept",
	"Accept-Charset",
	"Accept-Encoding",
	"Accept-Language",
	"Accept-Ranges",
	"Access-Control-Allow-Credentials",
	"Access-Control-Allow-Headers",
	"Access-Control-Allow-Methods",
	"Access-Control-Allow-Origin",
	"Access-Control-Max-Age",
	"Alt-Svc",
	"Age",
	"Authorization",
	"Cache-Control",
	"Connection",
	"Content-Disposition",
	"Content-Encoding",
	"Content-Language",
	"Content-Length",
	"Content-Range",
	"Content-Security-Policy",
	"Content-Security-Policy-Report-Only",
	"Content-Type",
	"Cookie",
	"Date",
	"Dnt",
	"Etag",
	"Expect-Ct",
	"Expect",
	"Expires",
	"Forwarded",
	"Host",
	"If-Match",
	"If-Modified-Since",
	"If-None-Match",
	"If-Unmodified-Since",
	"Keep-Alive",
	"Last-Modified",
	"Link",
	"Location",
	"Origin",
	"Pragma",
	"Referer",
	"Request-Id",
	"Retry-After",
	"Server",
	"Set-Cookie",
	"Strict-Transport-Security",
	"Upgrade",
	"User-Agent",
	"Vary",
	"Via",
	"Www-Authenticate",
	"X-Content-Type-Options",
	"X-Frame-Options",
	"X-Forwarded-For",
	"X-Forwarded-Host",
	"X-Forwarded-Proto",
}

func main() {
	const name = "httpheaders.go"
	generate(name)
	format(name)
}

func generate(name string) {
	fp, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	sort.Strings(commonHeaders)

	headers := make([][2]string, 0)
	for _, header := range commonHeaders {
		headers = append(headers, [2]string{makeVar(header), http.CanonicalHeaderKey(header)})
	}

	tmpl.Execute(fp, struct {
		Timestamp time.Time
		Headers   [][2]string
	}{
		Timestamp: time.Now(),
		Headers:   headers,
	})
}

func format(name string) {
	exec.Command("gofmt", "-w", name).Run()
}

func makeVar(k string) string {
	return strings.ReplaceAll(k, "-", "")
}

var tmpl = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// {{ .Timestamp }}

package httpheaders
//go:generate go run gen.go

const (
{{range $header := .Headers}}	{{index $header 0}} = "{{index $header 1}}"
{{end}}
)
`))
