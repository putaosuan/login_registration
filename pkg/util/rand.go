package util

import (
	"math/rand"
	"time"
)

const letterBytes = "012345679abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandBytes(n int) []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return b
}

func RandString(n int) string {
	s := RandBytes(n)
	return string(s)
}
