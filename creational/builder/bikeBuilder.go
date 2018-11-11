package builder

//BikeBuilder struct
type BikeBuilder struct {
	v VehicleProduct
}

//SetWheels function
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

//SetSeats function
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

//SetStructure function
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

//GetVehicle function
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
