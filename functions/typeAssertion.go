package main

import (
	"fmt"
)

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Beep")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func main() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	// type Assertion to initial type
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}
