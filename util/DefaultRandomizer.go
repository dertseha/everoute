package util

import (
	"math/rand"
	"time"
)

type defaultRandomizer struct {
	rand *rand.Rand
}

func DefaultRandomizer() Randomizer {
	result := &defaultRandomizer{
		rand: rand.New(rand.NewSource(time.Now().UnixNano()))}

	return result
}

func (rand *defaultRandomizer) Index(limit int) int {
	result := -1

	if limit > 0 {
		result = rand.rand.Intn(limit)
	}

	return result
}
