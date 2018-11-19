package builder

//BuildProcess interface to define the vehicle building steps
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//ManufacturingDirector the struct
type ManufacturingDirector struct {
	builder BuildProcess
}

//Construct function to construct a vehicle
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

//SetBuilder function
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
