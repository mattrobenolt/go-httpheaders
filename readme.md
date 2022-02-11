# httpheaders

```go
import "go.withmatt.com/httpheaders"

request.Headers.Get(httpheaders.ContentType)
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

### benchmarks

```
BenchmarkFromLibrary
BenchmarkFromLibrary-8                  61118968                19.02 ns/op            0 B/op          0 allocs/op
BenchmarkCanonicalAndCommon
BenchmarkCanonicalAndCommon-8           61921510                19.36 ns/op            0 B/op          0 allocs/op
BenchmarkNotCanonicalButCommon
BenchmarkNotCanonicalButCommon-8        30440584                39.97 ns/op            0 B/op          0 allocs/op
BenchmarkCanonicalNotCommon
BenchmarkCanonicalNotCommon-8           59834079                19.42 ns/op            0 B/op          0 allocs/op
BenchmarkNotCanonicalNotCommon
BenchmarkNotCanonicalNotCommon-8        12968305                96.71 ns/op           16 B/op          1 allocs/op
```

From this, we can see that the pre-canonicalized forms clock in at half the speed as a non-canonicalized common header.

And for a non-canonicalized not common header, this is significantly ~6x faster with 1 less allocation.
