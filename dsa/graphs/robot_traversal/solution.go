package robottraversal

import "fmt"

// RobotState manages the state of the robot simulation.
type RobotState struct {
	MaxX    int
	MaxY    int
	Visited map[string]bool
}

// NewRobotState creates a new robot state with the given boundaries.
func NewRobotState(maxX, maxY int) *RobotState {
	return &RobotState{
		MaxX:    maxX,
		MaxY:    maxY,
		Visited: make(map[string]bool),
	}
}

// IsValidPosition checks if the position is within the boundaries.
func (rs *RobotState) IsValidPosition(x, y int) bool {
	return x >= 0 && x <= rs.MaxX && y >= 0 && y <= rs.MaxY
}

// IsVisited checks if the position has been visited before.
func (rs *RobotState) IsVisited(x, y int) bool {
	key := fmt.Sprintf("%d,%d", x, y)
	return rs.Visited[key]
}

// MarkVisited marks the position as visited.
func (rs *RobotState) MarkVisited(x, y int) {
	key := fmt.Sprintf("%d,%d", x, y)
	rs.Visited[key] = true
}

// Execute runs the robot simulation with the given commands.
func Execute(maxX, maxY, startX, startY int, direction rune, commands string) (int, int, rune) {
	robot := NewRobot(startX, startY, direction)
	state := NewRobotState(maxX, maxY)
	
	state.MarkVisited(robot.X, robot.Y)
	
	for _, cmd := range commands {
		switch cmd {
		case 'L':
			robot.TurnLeft()
		case 'R':
			robot.TurnRight()
		case 'M':
			nextX, nextY := robot.NextPosition()
			
			if !state.IsValidPosition(nextX, nextY) {
				return robot.X, robot.Y, robot.Direction
			}
			
			if state.IsVisited(nextX, nextY) {
				return robot.X, robot.Y, robot.Direction
			}
			
			robot.Move()
			state.MarkVisited(robot.X, robot.Y)
		}
	}
	
	return robot.X, robot.Y, robot.Direction
}
