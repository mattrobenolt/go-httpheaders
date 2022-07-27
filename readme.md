# httpheaders

[![Go Reference](https://pkg.go.dev/badge/go.withmatt.com/httpheaders.svg)](https://pkg.go.dev/go.withmatt.com/httpheaders)

```go
import "go.withmatt.com/httpheaders"

req.Header.Get(httpheaders.ContentType)
```

This is a package of all the common HTTP header names pre-generated
from running through `http.CanonicalHeaderKey`.

## why?

Unless you type it correctly, it's actually faster in Go to
use the correct normalized form. Go first does a check to make
sure the string _is_ in it's canonical form. If it's not, it
will normalize the string for you.

Go stdlib maintains a list of "common headers" that are optimized
for this transformation, but pre-canonicalizing is twice as
fast even for the optimized path.

For a new header key that doesn't exist here and isn't common,
running `http.CanonicalHeaderKey` over the key first saves a
memory allocation and is about 6-8x faster in my testing.

Also having constants for common header values is beneficial
similar to http status code constants like `http.StatusOK` or
`http.MethodGet`.

It is also significantly more efficient to bypass the
`(http.Header).{Get,Set}` abstractions entirely if you're using
canonicalized keys as proved in the benchmarks. While using the
canonicalized form is faster in general, skipping the `{Get,Set}`
helpers avoid even doing the check necessary to confirm if it's in
the correct form.

### benchmarks

```
BenchmarkFromLibrary
BenchmarkFromLibrary-8                  58604780                19.26 ns/op            0 B/op          0 allocs/op
BenchmarkCanonicalAndCommon
BenchmarkCanonicalAndCommon-8           59234570                18.79 ns/op            0 B/op          0 allocs/op
BenchmarkNotCanonicalButCommon
BenchmarkNotCanonicalButCommon-8        31013407                38.36 ns/op            0 B/op          0 allocs/op
BenchmarkCanonicalNotCommon
BenchmarkCanonicalNotCommon-8           65812561                19.44 ns/op            0 B/op          0 allocs/op
BenchmarkNotCanonicalNotCommon
BenchmarkNotCanonicalNotCommon-8        12636105               108.2 ns/op            16 B/op          1 allocs/op
BenchmarkBypassingHelpers
BenchmarkBypassingHelpers/Get
BenchmarkBypassingHelpers/Get-8         60684620                18.98 ns/op            0 B/op          0 allocs/op
BenchmarkBypassingHelpers/GetDirect
BenchmarkBypassingHelpers/GetDirect-8           561376302                1.976 ns/op           0 B/op          0 allocs/op
BenchmarkBypassingHelpers/Set
BenchmarkBypassingHelpers/Set-8                 14616615                88.58 ns/op           16 B/op          1 allocs/op
BenchmarkBypassingHelpers/SetDirect
BenchmarkBypassingHelpers/SetDirect-8           16508486                65.67 ns/op           16 B/op          1 allocs/op
BenchmarkBypassingHelpers/SetDirectNoAlloc
BenchmarkBypassingHelpers/SetDirectNoAlloc-8    123816476               13.00 ns/op            0 B/op          0 allocs/op
```

From this, we can see that the pre-canonicalized forms clock in at half the speed as a non-canonicalized common header.

And for a non-canonicalized not common header, this is significantly ~6x faster with 1 less allocation.
