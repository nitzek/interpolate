package util

func Merge(dst map[string]string, src map[string]string) {
	for k, v := range src {
		dst[k] = v
	}
}
