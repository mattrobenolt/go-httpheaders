package httpheaders

type header = map[string][]string

// Add adds the key, value pair to the header. It appends to any
// existing values associated with key. The key is case sensitive.
func Add(h header, key, value string) {
	h[key] = append(h[key], value)
}

// Del deletes the values associated with key.
// The key is case sensitive.
func Del(h header, key string) {
	delete(h, key)
}

// Get gets the first value associated with the given key. If
// there are no values associated with the key, Get returns "".
// The key is case sensitive.
func Get(h header, key string) string {
	if h != nil {
		if v := h[key]; len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

// Set sets the header entries associated with key to the
// single element value. It replaces any existing values
// associated with key. The key is case sensitive.
func Set(h header, key, value string) {
	h[key] = []string{value}
}

// Values returns all values associated with the given key.
// It is case sensitive. The returned slice is not a copy.
func Values(h header, key string) []string {
	return h[key]
}
