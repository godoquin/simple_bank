package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RamdomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RamdomString(n int) string {
	var sb strings.Builder
	k := len(letterBytes)

	for i := 0; i < n; i++ {
		c := letterBytes[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RamdomOwner() string {
	return RamdomString(6)
}

func RamdomMoney() int64 {
	return RamdomInt(0, 1000)
}

func RamdomCurrency() string {
	currencies := []string{"EUR", "USD", "COP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
