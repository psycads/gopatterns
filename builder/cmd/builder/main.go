package main

import (
	"fmt"

	"github.com/psycads/gopattern/builder/vehicle"
	"github.com/psycads/gopattern/builder/vehicle/car"
)

func main() {
	var vb vehicle.VehicleBuilder
	vb = car.NewBuilder()
	vb = vb.SetMaxSpeed(120).
		SetSeats(6)
	car := vb.Build()
	car.Start()
	fmt.Printf("%v\n", car)
	fmt.Println("hi")
}
