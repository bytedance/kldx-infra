package redis

import "time"

// KeepTTL is an option for Set command to keep key's existing TTL.
// For example:
// Set(key, value, redis.KeepTTL)
const KeepTTL = -1

func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func formatMs(dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		return 1
	}
	return int64(dur / time.Millisecond)
}

func formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		return 1
	}
	return int64(dur / time.Second)
}

func appendArgs(dst, src []interface{}) []interface{} {
	if len(src) == 1 {
		return appendArg(dst, src[0])
	}

	dst = append(dst, src...)
	return dst
}

func appendArg(dst []interface{}, arg interface{}) []interface{} {
	switch arg := arg.(type) {
	case []string:
		for _, s := range arg {
			dst = append(dst, s)
		}
		return dst
	case []interface{}:
		dst = append(dst, arg...)
		return dst
	case map[string]interface{}:
		for k, v := range arg {
			dst = append(dst, k, v)
		}
		return dst
	default:
		return append(dst, arg)
	}
}
