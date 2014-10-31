package util

import (
	check "gopkg.in/check.v1"
)

type Vector3dTestSuite struct {
}

var _ = check.Suite(&Vector3dTestSuite{})

func (suite *Vector3dTestSuite) TestStringReturnsStringPresentation(c *check.C) {
	vec := Vector3d{0, 1, 2}
	result := vec.String()

	c.Assert(result, check.Equals, "[0.000000, 1.000000, 2.000000]")
}

func (suite *Vector3dTestSuite) TestLengthReturnsVectorLength(c *check.C) {
	vec := Vector3d{10, 20, 30}
	result := vec.Length()

	c.Assert(int(result*1000.0), check.Equals, 37416) // 37.416...
}
