package rules

import "github.com/dertseha/everoute/travel"

type TravelRule interface {
	Compare(sumA travel.TravelCostSum, sumB travel.TravelCostSum) float64
}
