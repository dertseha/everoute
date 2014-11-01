package rules

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"

	check "gopkg.in/check.v1"
)

type NaturalOrderTravelRuleTestSuite struct {
	nullCost universe.TravelCost
	rule     travel.TravelRule
}

var _ = check.Suite(&NaturalOrderTravelRuleTestSuite{})

func (suite *NaturalOrderTravelRuleTestSuite) SetUpTest(c *check.C) {
	suite.nullCost = universe.AddingTravelCost("test", 0.0)
	suite.rule = NaturalOrderTravelRule(suite.nullCost)
}

func (suite *NaturalOrderTravelRuleTestSuite) TestCompareReturnsZeroWhenSumDoesNotContainCost(c *check.C) {
	sumA := universe.EmptyTravelCostSum()
	sumB := universe.EmptyTravelCostSum()
	result := suite.rule.Compare(sumA, sumB)

	c.Assert(result, check.Equals, 0.0)
}

func (suite *NaturalOrderTravelRuleTestSuite) TestCompareReturnsPositiveWhenFirstSumIsBigger(c *check.C) {
	sumA := universe.EmptyTravelCostSum().Add(universe.SingleTravelCostSum(universe.AddingTravelCost(suite.nullCost.Type(), 10)))
	sumB := universe.EmptyTravelCostSum().Add(universe.SingleTravelCostSum(universe.AddingTravelCost(suite.nullCost.Type(), 5)))
	result := suite.rule.Compare(sumA, sumB)

	c.Assert(result > 0.0, check.Equals, true)
}
