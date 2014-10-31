package util

// MetersPerAu specifies the number of meters one Astronomical Unit (AU) has.
// This value is the 2009 defined number ( See http://en.wikipedia.org/wiki/Astronomical_unit ).
// Note that EVE Online was created way earlier than this definition and might thus deviate from this value.
const MetersPerAu float64 = 149597870700

// MetersPerLy specifies the number of meters one Light Year (LY) has.
// This value is based on the constant MetersPerAu.
const MetersPerLy float64 = MetersPerAu * 63241
