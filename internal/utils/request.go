package utils

import "time"

func FindMaxRequestTime(t1, t2 time.Duration) time.Duration {
	if t1 > t2 {
		return t1
	}
	return t2
}

func FindMinRequestTime(t1, t2 time.Duration) time.Duration {
	if t1 < t2 {
		return t1
	}
	return t2
}
