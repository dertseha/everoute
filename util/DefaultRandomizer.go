package util

import (
	"math/rand"
	"time"
)

// defaultRandomizer implements the Randomizer interface based on the VM randomizer.
type defaultRandomizer struct {
	rand *rand.Rand
}

// DefaultRandomizer returns a Ranomizer instance used for regular processing.
func DefaultRandomizer() Randomizer {
	result := &defaultRandomizer{
		rand: rand.New(rand.NewSource(time.Now().UnixNano()))}

	return result
}

// Index is the Randomizer interface implementation.
func (rand *defaultRandomizer) Index(limit int) int {
	result := -1

	if limit > 0 {
		result = rand.rand.Intn(limit)
	}

	return result
}
