package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("Wheels of car must be 4")
		return
	}
	if car.Seats != 5 {
		t.Errorf("Seats of car must be 4")
		return
	}
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.Build()
	motorbike.Wheels = 2
	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
		return
	}
	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was % s\n", motorbike.Structure)
		return
	}
}
