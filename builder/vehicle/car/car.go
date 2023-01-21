package car

var defaultConfig config = config{
	Seats:    5,
	MaxSpeed: 90,
}

type Car struct {
	config
	Status string
}

func (c *Car) Start() error {
	c.Status = "Running"
	return nil
}

func (c *Car) Speed() int {
	return c.MaxSpeed
}

func (c *Car) PassengerCap() int {
	return c.Seats
}

func (c *Car) Passengers() int {
	return 0
}

type config struct {
	Seats    int
	MaxSpeed int
}
