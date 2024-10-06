package random

import (
	"math/rand"
	"time"
)

var localRand *rand.Rand

func init() {
	localRand = rand.New(rand.NewSource(time.Now().UnixMicro()))
}

func Choice[T any](items []T) T {
	if len(items) == 0 {
		return *new(T)
	}
	return items[localRand.Intn(len(items))]
}

func SetRand(rnd *rand.Rand) {
	localRand = rnd
}
