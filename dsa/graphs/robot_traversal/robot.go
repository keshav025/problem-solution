package robottraversal

// Robot represents a robot on a rectangular plane.
type Robot struct {
	X         int
	Y         int
	Direction rune
}

// NewRobot creates a new robot with the given position and direction.
func NewRobot(x, y int, direction rune) *Robot {
	return &Robot{
		X:         x,
		Y:         y,
		Direction: direction,
	}
}

// TurnLeft rotates the robot 90 degrees to the left.
func (r *Robot) TurnLeft() {
	switch r.Direction {
	case 'N':
		r.Direction = 'W'
	case 'W':
		r.Direction = 'S'
	case 'S':
		r.Direction = 'E'
	case 'E':
		r.Direction = 'N'
	}
}

// TurnRight rotates the robot 90 degrees to the right.
func (r *Robot) TurnRight() {
	switch r.Direction {
	case 'N':
		r.Direction = 'E'
	case 'E':
		r.Direction = 'S'
	case 'S':
		r.Direction = 'W'
	case 'W':
		r.Direction = 'N'
	}
}

// NextPosition returns the next position if the robot moves forward.
func (r *Robot) NextPosition() (int, int) {
	x, y := r.X, r.Y
	switch r.Direction {
	case 'N':
		y++
	case 'S':
		y--
	case 'E':
		x++
	case 'W':
		x--
	}
	return x, y
}

// Move moves the robot forward one position.
func (r *Robot) Move() {
	r.X, r.Y = r.NextPosition()
}
