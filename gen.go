//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
	"os/exec"
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
	"Age",
	"Authorization",
	"Cache-Control",
	"Cc",
	"Connection",
	"Content-Id",
	"Content-Language",
	"Content-Length",
	"Content-Security-Policy",
	"Content-Security-Policy-Report-Only",
	"Content-Transfer-Encoding",
	"Content-Type",
	"Cookie",
	"Date",
	"Dkim-Signature",
	"Etag",
	"Expires",
	"From",
	"Host",
	"If-Modified-Since",
	"If-None-Match",
	"In-Reply-To",
	"Last-Modified",
	"Location",
	"Message-Id",
	"Mime-Version",
	"Origin",
	"Pragma",
	"Received",
	"Referer",
	"Return-Path",
	"Server",
	"Set-Cookie",
	"Subject",
	"To",
	"User-Agent",
	"Vary",
	"Via",
	"X-Forwarded-For",
	"X-Imforwards",
	"X-Powered-By",
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
