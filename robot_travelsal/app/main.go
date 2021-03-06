package main

import (
	"fmt"
	robot "problem-solution/robot_travelsal/robottravelsal"
)

func main() {
	r := &robot.Robot{State: robot.GetRobotNorthState(), MaxX: 5, MaxY: 5}
	r.Move("MMMRMMLM")
	fmt.Println(r.State.GetCordinate(), r.State.GetDirection())
}
