package robottraversal

import "testing"

func TestRobotTraversal(t *testing.T) {
	tests := []struct {
		name      string
		maxX      int
		maxY      int
		startX    int
		startY    int
		direction rune
		commands  string
		wantX     int
		wantY     int
		wantDir   rune
	}{
		{
			name:      "Example test case",
			maxX:      4,
			maxY:      4,
			startX:    0,
			startY:    0,
			direction: 'N',
			commands:  "MMMRMMLM",
			wantX:     2,
			wantY:     4,
			wantDir:   'N',
		},
		{
			name:      "Simple forward movement",
			maxX:      5,
			maxY:      5,
			startX:    0,
			startY:    0,
			direction: 'N',
			commands:  "MMM",
			wantX:     0,
			wantY:     3,
			wantDir:   'N',
		},
		{
			name:      "Turn right and move",
			maxX:      5,
			maxY:      5,
			startX:    0,
			startY:    0,
			direction: 'N',
			commands:  "RMM",
			wantX:     2,
			wantY:     0,
			wantDir:   'E',
		},
		{
			name:      "Turn left and move",
			maxX:      5,
			maxY:      5,
			startX:    2,
			startY:    2,
			direction: 'N',
			commands:  "LMM",
			wantX:     0,
			wantY:     2,
			wantDir:   'W',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotDir := Execute(tt.maxX, tt.maxY, tt.startX, tt.startY, tt.direction, tt.commands)
			if gotX != tt.wantX || gotY != tt.wantY || gotDir != tt.wantDir {
				t.Errorf("Execute() = (%d, %d, %c), want (%d, %d, %c)", gotX, gotY, gotDir, tt.wantX, tt.wantY, tt.wantDir)
			}
		})
	}
}

func TestRobotTurnLeft(t *testing.T) {
	tests := []struct {
		initial  rune
		expected rune
	}{
		{'N', 'W'},
		{'W', 'S'},
		{'S', 'E'},
		{'E', 'N'},
	}

	for _, tt := range tests {
		robot := NewRobot(0, 0, tt.initial)
		robot.TurnLeft()
		if robot.Direction != tt.expected {
			t.Errorf("TurnLeft() from %c = %c, want %c", tt.initial, robot.Direction, tt.expected)
		}
	}
}

func TestRobotTurnRight(t *testing.T) {
	tests := []struct {
		initial  rune
		expected rune
	}{
		{'N', 'E'},
		{'E', 'S'},
		{'S', 'W'},
		{'W', 'N'},
	}

	for _, tt := range tests {
		robot := NewRobot(0, 0, tt.initial)
		robot.TurnRight()
		if robot.Direction != tt.expected {
			t.Errorf("TurnRight() from %c = %c, want %c", tt.initial, robot.Direction, tt.expected)
		}
	}
}
