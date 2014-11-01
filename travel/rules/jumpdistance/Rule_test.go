package jumpdistance

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"

	check "gopkg.in/check.v1"
)

type RuleTestSuite struct {
	rule travel.TravelRule
}

var _ = check.Suite(&RuleTestSuite{Rule()})

func (suite *RuleTestSuite) TestCompareConsidersLowerAmountBetter(c *check.C) {
	sumA := universe.EmptyTravelCostSum().Add(universe.SingleTravelCostSum(Cost(10.0)))
	sumB := universe.EmptyTravelCostSum().Add(universe.SingleTravelCostSum(Cost(20.0)))
	result := suite.rule.Compare(sumA, sumB)

	c.Assert(result < 0.0, check.Equals, true)
}
