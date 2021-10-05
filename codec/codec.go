package codec

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

func GenerateShort() string {
	//the following couple of lines are a weak point,
	//but i'd ususally use the database's ID here
	rand.Seed(time.Now().UnixNano())
	s := rand.Uint64()

	var b strings.Builder
	b.Grow(11)

	for ; s > 0; s = s / length {
		b.WriteByte(alphabet[(s % length)])
	}

	return b.String()
}
