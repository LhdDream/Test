package service

type fnv64a struct{}

const (
	offset64 = 14695981039346656037
	prime64 = 1099511628211
)

func newDefaultHash() fnv64a {
	return fnv64a{}
}

func (f fnv64a) Sum64(Key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i<len(Key) ; i++ {
		hash ^= uint64(Key[i])
		hash *= prime64
	}
	return hash
}