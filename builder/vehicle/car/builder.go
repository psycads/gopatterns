package car

import "github.com/psycads/gopattern/builder/vehicle"

type Builder struct {
	config config
}

func NewBuilder() vehicle.VehicleBuilder {
	return &Builder{defaultConfig}
}

func (b *Builder) Build() vehicle.Vehicle {
	return &Car{
		config: b.config,
	}
}

func (b *Builder) SetMaxSpeed(speed int) vehicle.VehicleBuilder {
	b.config.MaxSpeed = speed
	return b
}

func (b *Builder) SetSeats(count int) vehicle.VehicleBuilder {
	b.config.Seats = count
	return b
}
