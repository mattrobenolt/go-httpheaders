package httpheaders

import (
	"net/http"
	"reflect"
	"testing"
)

func BenchmarkFromLibrary(b *testing.B) {
	h := make(http.Header)
	key := UserAgent
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
	b.ReportAllocs()
}

func BenchmarkCanonicalAndCommon(b *testing.B) {
	h := make(http.Header)
	key := "User-Agent"
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
	b.ReportAllocs()
}

func BenchmarkNotCanonicalButCommon(b *testing.B) {
	h := make(http.Header)
	key := "user-agent"
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
	b.ReportAllocs()
}

func BenchmarkCanonicalNotCommon(b *testing.B) {
	h := make(http.Header)
	key := http.CanonicalHeaderKey("matt-agent")
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
	b.ReportAllocs()
}

func BenchmarkNotCanonicalNotCommon(b *testing.B) {
	h := make(http.Header)
	key := "matt-agent"
	for i := 0; i < b.N; i++ {
		h.Get(key)
	}
	b.ReportAllocs()
}

func BenchmarkBypassingHelpers(b *testing.B) {
	key := http.CanonicalHeaderKey("matt-agent")

	b.Run("Get", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			h.Get(key)
		}
		b.ReportAllocs()
	})

	b.Run("GetDirect", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			Get(h, key)
		}
		b.ReportAllocs()
	})

	b.Run("Set", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			h.Set(key, "bar")
		}
		b.ReportAllocs()
	})

	b.Run("SetDirect", func(b *testing.B) {
		h := make(http.Header)
		for i := 0; i < b.N; i++ {
			h[key] = []string{"bar"}
		}
		b.ReportAllocs()
	})

	b.Run("SetDirectNoAlloc", func(b *testing.B) {
		h := make(http.Header)
		v := []string{"bar"}
		for i := 0; i < b.N; i++ {
			h[key] = v
		}
		b.ReportAllocs()
	})
}

func TestGet(t *testing.T) {
	var h http.Header
	var expect, out string

	expect = h.Get("Foo")
	out = Get(h, "Foo")
	if expect != out {
		t.Fatalf("%v: not equal to: %v", out, expect)
	}

	h = make(http.Header)
	expect = h.Get("Foo")
	out = Get(h, "Foo")
	if expect != out {
		t.Fatalf("%v: not equal to: %v", out, expect)
	}

	h["Foo"] = []string{}
	expect = h.Get("Foo")
	out = Get(h, "Foo")
	if expect != out {
		t.Fatalf("%v: not equal to: %v", out, expect)
	}

	h["Foo"] = []string{"bar"}
	expect = h.Get("Foo")
	out = Get(h, "Foo")
	if expect != out {
		t.Fatalf("%v: not equal to: %v", out, expect)
	}

	h["Foo"] = []string{"bar", "baz"}
	expect = h.Get("Foo")
	out = Get(h, "Foo")
	if expect != out {
		t.Fatalf("%v: not equal to: %v", out, expect)
	}

	expect = h.Get("Nothing")
	out = Get(h, "Nothing")
	if expect != out {
		t.Fatalf("%v: not equal to: %v", out, expect)
	}
}

func TestAdd(t *testing.T) {
	h1 := make(http.Header)
	h2 := make(http.Header)

	h1.Add("Foo", "bar")
	Add(h2, "Foo", "bar")
	if !reflect.DeepEqual(h1, h2) {
		t.Fatalf("%v: should be equal to: %v", h1, h2)
	}

	h1.Add("Foo", "baz")
	Add(h2, "Foo", "baz")
	if !reflect.DeepEqual(h1, h2) {
		t.Fatalf("%v: should be equal to: %v", h1, h2)
	}

	h1.Add("A", "b")
	Add(h2, "A", "b")
	if !reflect.DeepEqual(h1, h2) {
		t.Fatalf("%v: should be equal to: %v", h1, h2)
	}
}

func TestDel(t *testing.T) {
	h1 := http.Header{
		"Foo": []string{"bar"},
		"A":   []string{"b"},
	}
	h2 := h1.Clone()

	h1.Del("Foo")
	Del(h2, "Foo")
	if !reflect.DeepEqual(h1, h2) {
		t.Fatalf("%v: should be equal to: %v", h1, h2)
	}
}

func TestSet(t *testing.T) {
	h1 := make(http.Header)
	h2 := make(http.Header)

	h1.Set("Foo", "bar")
	Set(h2, "Foo", "bar")
	if !reflect.DeepEqual(h1, h2) {
		t.Fatalf("%v: should be equal to: %v", h1, h2)
	}

	h1.Set("Foo", "baz")
	Set(h2, "Foo", "baz")
	if !reflect.DeepEqual(h1, h2) {
		t.Fatalf("%v: should be equal to: %v", h1, h2)
	}

	h1.Set("A", "b")
	Set(h2, "A", "b")
	if !reflect.DeepEqual(h1, h2) {
		t.Fatalf("%v: should be equal to: %v", h1, h2)
	}
}

func TestValues(t *testing.T) {
	var h1, h2 http.Header
	var v1, v2 []string

	v1 = h1.Values("Foo")
	v2 = Values(h2, "Foo")
	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("%v: should be equal to: %v", v1, v2)
	}

	h1 = make(http.Header)
	h2 = make(http.Header)

	v1 = h1.Values("Foo")
	v2 = Values(h2, "Foo")
	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("%v: should be equal to: %v", v1, v2)
	}

	h1 = http.Header{
		"Foo": []string{"bar"},
		"A":   []string{"b"},
	}
	h2 = h1.Clone()

	v1 = h1.Values("Foo")
	v2 = Values(h2, "Foo")
	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("%v: should be equal to: %v", v1, v2)
	}

	v1 = h1.Values("Nothing")
	v2 = Values(h2, "Nothing")
	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("%v: should be equal to: %v", v1, v2)
	}
}
