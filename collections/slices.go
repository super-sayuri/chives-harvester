package collections

import (
	"math/rand"
	"time"
)

func SliceRandomPick[T any](ts []T) T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(ts))
	return ts[i]
}
