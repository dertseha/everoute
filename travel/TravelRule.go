package travel

type TravelRule interface {
	Compare(sumA *TravelCostSum, sumB *TravelCostSum) float64
}
