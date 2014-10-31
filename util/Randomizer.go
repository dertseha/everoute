package util

// Randomizer represents an interface for any random value generator in use.
// It is provided to hide the source of random values and allow reproducible tests.
type Randomizer interface {
	// Index returns an integer value in the range of [0..limit[ or -1 if the limit is 0 or negative.
	// It is used to create a random value for an array index.
	Index(limit int) int
}
