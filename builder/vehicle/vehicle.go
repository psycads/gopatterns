package vehicle

type Vehicle interface {
	Start() error
	Speed() int
	PassengerCap() int
	Passengers() int
}

type VehicleBuilder interface {
	Build() Vehicle
	SetSeats(int) VehicleBuilder
	SetMaxSpeed(int) VehicleBuilder
}
