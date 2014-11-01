package rules

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/universe"

	check "gopkg.in/check.v1"
)

type TravelRulesetTestSuite struct {
	costType1 string
	costType2 string
	sumA      *universe.TravelCostSum
	sumB      *universe.TravelCostSum
}

var _ = check.Suite(&TravelRulesetTestSuite{})

func (suite *TravelRulesetTestSuite) SetUpTest(c *check.C) {
	suite.costType1 = "test"
	suite.costType2 = "equal"

	suite.sumA = universe.EmptyTravelCostSum()
	suite.sumA = suite.sumA.Add(universe.SingleTravelCostSum(travel.AddingTravelCost(suite.costType1, 5)))
	suite.sumA = suite.sumA.Add(universe.SingleTravelCostSum(travel.AddingTravelCost(suite.costType2, 100)))

	suite.sumB = universe.EmptyTravelCostSum()
	suite.sumB = suite.sumB.Add(universe.SingleTravelCostSum(travel.AddingTravelCost(suite.costType1, 10)))
	suite.sumB = suite.sumB.Add(universe.SingleTravelCostSum(travel.AddingTravelCost(suite.costType2, 100)))
}

func (suite *TravelRulesetTestSuite) TestCompareReturnsZeroWhenNoRulesDefined(c *check.C) {
	rule := TravelRuleset()
	result := rule.Compare(suite.sumA, suite.sumB)

	c.Assert(result, check.Equals, 0.0)
}

func (suite *TravelRulesetTestSuite) TestCompareConsidersNextRuleWhenPreviousReturnedZero(c *check.C) {
	rule1 := NaturalOrderTravelRule(travel.AddingTravelCost(suite.costType2, 0.0))
	rule2 := NaturalOrderTravelRule(travel.AddingTravelCost(suite.costType1, 0.0))
	rule := TravelRuleset(rule1, rule2)
	result := rule.Compare(suite.sumA, suite.sumB)

	c.Assert(result < 0.0, check.Equals, true)
}
