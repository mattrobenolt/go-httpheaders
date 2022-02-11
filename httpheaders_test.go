package httpheaders

import (
	"testing"
	"net/http"
)

func BenchmarkFromLibrary(b *testing.B) {
	h := make(http.Header)
	key := UserAgent
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
}

func BenchmarkCanonicalAndCommon(b *testing.B) {
	h := make(http.Header)
	key := "User-Agent"
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
}

func BenchmarkNotCanonicalButCommon(b *testing.B) {
	h := make(http.Header)
	key := "user-agent"
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
}

func BenchmarkCanonicalNotCommon(b *testing.B) {
	h := make(http.Header)
	key := http.CanonicalHeaderKey("matt-agent")
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
}

func BenchmarkNotCanonicalNotCommon(b *testing.B) {
	h := make(http.Header)
	key := "matt-agent"
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
}
