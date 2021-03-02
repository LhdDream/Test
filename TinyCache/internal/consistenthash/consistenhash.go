package consistenthash

import (
	"strconv"
)

func JumpHash(m string, buckets int) int {
	var b, j int64
	if buckets <= 0 {
		buckets = 1
	}
	key, _ := strconv.ParseInt(m, 10, 64)
	for j < int64(buckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return int(b)
}
