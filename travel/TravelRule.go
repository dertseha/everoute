package travel

import "github.com/dertseha/everoute/universe"

// TravelRule describes a rule to compare two travel costs.
type TravelRule interface {
	// Compare receives two travel cost sums and compares them. The returned value describes the relationship.
	// A returned value of 0.0 considers the two costs to be equal. <0 considers sumA to be better, >0 considers sumB
	// to be better.
	Compare(sumA *universe.TravelCostSum, sumB *universe.TravelCostSum) float64
}
