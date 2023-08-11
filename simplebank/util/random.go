package util

import (
	"math/rand"
	"strings"
	"time"
)

const alph = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randString(n int) string {
	var sb strings.Builder
	k := len(alph)

	for i := 0; i < n; i++ {
		c := alph[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()

}

func randOwner() string {
	return randString(6)
}

func randMoney() int64 {
	return randInt(0, 1000)
}

func randCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
