package utils

import "time"

// finfing the max request time and compare
func FindMaxRequestTime(t1, t2 time.Duration) time.Duration {
	if t1 > t2 {
		return t1
	}
	return t2
}

// find the min request time and compare
func FindMinRequestTime(t1, t2 time.Duration) time.Duration {
	if t1 < t2 {
		return t1
	}
	return t2
}
