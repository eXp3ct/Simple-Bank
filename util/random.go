package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var stringBuilder strings.Builder
	var k = len(alphabet)

	for i := 0; i < n; i++ {
		var char = alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(char)
	}

	return stringBuilder.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	var currencies = []string{
		"EUR",
		"RUB",
		"USD",
	}
	var curLen = len(currencies)
	return currencies[rand.Intn(curLen)]
}
