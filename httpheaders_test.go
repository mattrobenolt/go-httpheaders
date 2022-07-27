package httpheaders

import (
	"net/http"
	"testing"
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

func BenchmarkBypassingHelpers(b *testing.B) {
	key := http.CanonicalHeaderKey("matt-agent")

	b.Run("Get", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			h.Get(key)
		}
	})

	b.Run("GetDirect", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			getKey(h, key)
		}
	})

	b.Run("Set", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			h.Set(key, "bar")
		}
	})

	b.Run("SetDirect", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			h[key] = []string{"bar"}
		}
	})

	b.Run("SetDirectNoAlloc", func(b *testing.B) {
		h := make(http.Header)
		v := []string{"bar"}
		for i := 0; i < b.N; i++ {
			h[key] = v
		}
	})
}

func getKey(h http.Header, key string) string {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}
