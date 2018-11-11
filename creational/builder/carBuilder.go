package builder

//CarBuilder struct
type CarBuilder struct {
	v VehicleProduct
}

//SetWheels function
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

//SetSeats function
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

//SetStructure function
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

//GetVehicle function
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
