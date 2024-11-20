package main

import (
	"fmt"
	"sort"
)

type Car struct {
	position int
	timeToTarget float64
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}
	cars := make([]Car, n)
	
	for i := 0; i < n; i++ {
		timeToTarget := float64(target - position[i]) / float64(speed[i])
		cars = append(cars, Car{position: position[i], timeToTarget: timeToTarget})		
	}

	// Sort cars by position in descending order
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	// fleets count
	fleets := 0
	latestTime := 0.0

	for _, car := range cars {
		if car.timeToTarget > latestTime {
			fleets++
			latestTime = car.timeToTarget
		}
	}

	return fleets
}

func main() {
	target := 12
	pos := []int{10,8,2,6}
	speed := []int{2,4,2,1}

	fleets := carFleet(target, pos, speed)
	fmt.Println(fleets)
}
