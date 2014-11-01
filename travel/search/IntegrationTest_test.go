package search

import (
	"github.com/dertseha/everoute/travel"
	"github.com/dertseha/everoute/travel/capabilities"
	"github.com/dertseha/everoute/travel/capabilities/jumpgate"
	"github.com/dertseha/everoute/travel/rules"
	"github.com/dertseha/everoute/travel/rules/security"
	"github.com/dertseha/everoute/travel/rules/transitcount"
	"github.com/dertseha/everoute/universe"

	check "gopkg.in/check.v1"
)

type routeSearchResultCollector struct {
	channel chan *Route
}

func (collector *routeSearchResultCollector) Collect(route *Route) {
	collector.channel <- route
}

type IntegrationTestSuite struct {
	verse                universe.Universe
	solarSystemIdsByName map[string]universe.Id
	solarSystemNamesById map[universe.Id]string

	capabilities []travel.TravelCapability
	rules        []travel.TravelRule
}

var _ = check.Suite(&IntegrationTestSuite{})

func (suite *IntegrationTestSuite) SetUpSuite(c *check.C) {
	suite.solarSystemIdsByName = getSolarSystemIdsByName()
	suite.solarSystemNamesById = make(map[universe.Id]string)
	for name, id := range suite.solarSystemIdsByName {
		suite.solarSystemNamesById[id] = name
	}

	suite.verse = BuildHeimatar()
}

func (suite *IntegrationTestSuite) SetUpTest(c *check.C) {
	suite.capabilities = make([]travel.TravelCapability, 0)
	suite.capabilities = append(suite.capabilities, jumpgate.JumpGateTravelCapability(suite.verse))
	suite.rules = make([]travel.TravelRule, 0)
}

func (suite *IntegrationTestSuite) OptimizedSystemSearchCriterion(to string, rule travel.TravelRule) SearchCriterion {
	toId := suite.solarSystemIdsByName[to]
	criteria := make([]SearchCriterion, 0)

	criteria = append(criteria, DestinationSystemSearchCriterion(toId))
	criteria = append(criteria, CostAwareSearchCriterion(rule))

	return CombiningSearchCriterion(criteria...)
}

func (suite *IntegrationTestSuite) AddRule(rule travel.TravelRule) {
	suite.rules = append(suite.rules, rule)
}

func (suite *IntegrationTestSuite) Rule() travel.TravelRule {
	if len(suite.rules) == 0 {
		suite.AddRule(transitcount.Rule())
	}

	return rules.TravelRuleset(suite.rules...)
}

func (suite *IntegrationTestSuite) VerifyRoute(c *check.C, from string, via []string, to string, expected []string) {
	starts := []travel.Path{travel.NewPath(travel.NewStep(suite.solarSystemIdsByName[from], universe.AnyLocation(), universe.EmptyTravelCostSum(), universe.EmptyTravelCostSum()))}
	searchDone := make(chan int)
	routeChannel := make(chan *Route)
	capability := capabilities.CombiningTravelCapability(suite.capabilities...)
	rule := suite.Rule()
	collector := &routeSearchResultCollector{channel: routeChannel}
	done := false
	var foundRoute *Route = nil

	builder := NewRouteFinder(capability, rule, starts, collector, func() { searchDone <- 1; close(searchDone) })
	for _, viaSystem := range via {
		builder.AddWaypoint(suite.OptimizedSystemSearchCriterion(viaSystem, rule))
	}
	if to != "" {
		builder.ForDestination(suite.OptimizedSystemSearchCriterion(to, rule))
	}

	var _ = builder.Build()

	for !done {
		select {
		case route := <-routeChannel:
			foundRoute = route
		case <-searchDone:
			done = true
		}
	}
	result := make([]string, 0)
	if foundRoute != nil {
		steps := foundRoute.Steps()
		for _, step := range steps {
			result = append(result, suite.solarSystemNamesById[step.SolarSystemId()])
		}
	}
	close(routeChannel)
	c.Assert(result, check.DeepEquals, expected)
}

func (suite *IntegrationTestSuite) TestUniverseContainsRens(c *check.C) {
	result := suite.verse.HasSolarSystem(suite.solarSystemIdsByName["Rens"])

	c.Assert(result, check.Equals, true)
}

func (suite *IntegrationTestSuite) TestRouteFoundWithOnlySourceSystem(c *check.C) {
	suite.VerifyRoute(c, "Rens", make([]string, 0), "", []string{"Rens"})
}

func (suite *IntegrationTestSuite) TestRouteFromRensToPator(c *check.C) {
	suite.VerifyRoute(c, "Rens", make([]string, 0), "Pator", []string{"Rens", "Frarn", "Gyng", "Onga", "Pator"})
}

func (suite *IntegrationTestSuite) TestRouteFromRensToIvar(c *check.C) {
	suite.VerifyRoute(c, "Rens", make([]string, 0), "Ivar", []string{"Rens", "Odatrik", "Trytedald", "Ivar"})
}

func (suite *IntegrationTestSuite) TestRouteFromEmolgranlanToEddarViaHighSec(c *check.C) {
	suite.AddRule(security.MinRule(0.5))
	suite.AddRule(transitcount.Rule())

	suite.VerifyRoute(c, "Emolgranlan", make([]string, 0), "Eddar",
		[]string{"Emolgranlan", "Ammold", "Pator", "Onga", "Gyng",
			"Frarn", "Meirakulf", "Ivar", "Ameinaka", "Hulm",
			"Edmalbrurdus", "Kronsur", "Dumkirinur", "Obrolber", "Austraka",
			"Gerek", "Gerbold", "Offugen", "Eddar"})
}

func (suite *IntegrationTestSuite) TestRouteFromRensToBalginiaViaGyng(c *check.C) {
	suite.AddRule(transitcount.Rule())

	suite.VerifyRoute(c, "Rens", []string{"Gyng"}, "Balginia",
		[]string{"Rens", "Frarn", "Gyng", "Frarn", "Illinfrik", "Balginia"})
}
