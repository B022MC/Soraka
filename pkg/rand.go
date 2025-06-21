package pkg

import (
	"math/rand"
	"time"
)

var (
	defaultRd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandomAlphaNum(lengthParam ...int) []byte {
	length := 16
	if len(lengthParam) > 0 {
		length = lengthParam[0]
	}
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[defaultRd.Intn(len(bytes))])
	}
	return result
}
func InArray[T comparable](a T, arr []T) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}
