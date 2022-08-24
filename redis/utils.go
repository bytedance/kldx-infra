package redis

import "time"

// KeepTTL is an option for Set command to keep key's existing TTL.
// For example:
// Set(key, value, redis.KeepTTL)
const KeepTTL = -1

func usePrecise(param time.Duration) bool {
	return param < time.Second || param%time.Second != 0
}

func formatMils(param time.Duration) int64 {
	if param < time.Millisecond && param > 0 {
		return 1
	}
	return int64(param / time.Millisecond)
}

func formatSecond(param time.Duration) int64 {
	if param < time.Second && param > 0 {
		return 1
	}
	return int64(param / time.Second)
}

func appendArguments(destination, source []interface{}) []interface{} {
	destination = append(destination, source...)
	return destination
}

func appendArgument(destination []interface{}, source interface{}) []interface{} {
	switch argument := source.(type) {
	case []interface{}:
		destination = append(destination, argument...)
		return destination
	case map[string]interface{}:
		for key, value := range argument {
			destination = append(destination, key, value)
		}
		return destination
	case []string:
		for _, arg := range argument {
			destination = append(destination, arg)
		}
		return destination
	default:
		return append(destination, argument)
	}
}
