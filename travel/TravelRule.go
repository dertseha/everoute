package travel

import "github.com/dertseha/everoute/universe"

type TravelRule interface {
	Compare(sumA *universe.TravelCostSum, sumB *universe.TravelCostSum) float64
}
