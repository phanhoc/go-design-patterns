package builder

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}
