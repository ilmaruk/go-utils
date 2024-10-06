package random

import (
	"math/rand"
	"time"
)

type Rand[T any] struct {
	rnd *rand.Rand
}

func New[T any]() *Rand[T] {
	return NewWithSeed[T](time.Now().UnixMicro())
}

func NewWithSeed[T any](seed int64) *Rand[T] {
	return NewWithRand[T](rand.New(rand.NewSource(seed)))
}

func NewWithRand[T any](rnd *rand.Rand) *Rand[T] {
	return &Rand[T]{
		rnd: rnd,
	}
}

func (r *Rand[T]) Choice(items []T) T {
	if len(items) == 0 {
		return *new(T)
	}
	return items[r.rnd.Intn(len(items))]
}

func Choice[T any](items []T) T {
	if len(items) == 0 {
		return *new(T)
	}
	rnd := New[T]()
	return rnd.Choice(items)
}
