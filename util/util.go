package util

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

const (
	alphabet  = "abcdefghijklmnopqrstuvwxyz"
	numberbet = "1234567890"
)

func GenRandomString(n int) string {
	src := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(src)
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = alphabet[r.Intn(len(alphabet))]
	}
	return string(b)
}

func GenRandomNumber(n int) string {
	b := make([]byte, n)
	for i, _ := range b {
		num, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(numberbet))))
		b[i] = numberbet[num.Int64()]
	}
	return string(b)
}

func GenRandomName(n int) string {
	return GenRandomString(n)
}

func GenRandomPassword(n int) string {
	return GenRandomNumber(n)
}
