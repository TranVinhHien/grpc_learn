package generate

import (
	"math/rand"
	"strings"
	"time"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}
func RandomInt(min, max int) int {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return seededRand.Intn(max-min+1) + min
}
func RandomBool() bool {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return seededRand.Intn(2) == 0
}
