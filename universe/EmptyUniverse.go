package universe

import "fmt"

type emptyUniverse struct{}

func New() Universe {
	return new(emptyUniverse)
}

func (this *emptyUniverse) Extend() *UniverseBuilder {
	return newUniverseBuilder(this)
}

func (this *emptyUniverse) HasSolarSystem(id Id) bool {
	return false
}

func (this *emptyUniverse) SolarSystem(id Id) SolarSystem {
	panic(fmt.Sprintf("SolarSystem with ID <%d> not found", id))
}

func (this *emptyUniverse) SolarSystemIds() []Id {
	return nil
}
