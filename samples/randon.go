package samples

import (
	"math/rand"
	"time"
)

func getRandonString() string {
	rand.Seed(time.Now().UnixNano())
	n := 10
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(rand.Intn(26) + 65)
	}
	return string(b)
}
