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

func (r *Rand[T]) Shuffle(items []T) []T {
	output := make([]T, len(items))
	copy(output, items)
	for i := 0; i < len(output); i++ {
		j := r.rnd.Intn(len(output))
		output[i], output[j] = output[j], output[i]
	}
	return output
}

func Choice[T any](items []T) T {
	rnd := New[T]()
	return rnd.Choice(items)
}

func Shuffle[T any](items []T) []T {
	rnd := New[T]()
	return rnd.Shuffle(items)
}
