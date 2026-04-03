package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	robottraversal "problem-solution/dsa/graphs/robot_traversal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter grid dimensions (M N):")
	scanner.Scan()
	var maxX, maxY int
	fmt.Sscanf(scanner.Text(), "%d %d", &maxX, &maxY)

	fmt.Println("Enter robot position (X Y Direction):")
	scanner.Scan()
	var startX, startY int
	var direction string
	fmt.Sscanf(scanner.Text(), "%d %d %s", &startX, &startY, &direction)

	fmt.Println("Enter commands:")
	scanner.Scan()
	commands := strings.TrimSpace(scanner.Text())

	finalX, finalY, finalDir := robottraversal.Execute(
		maxX, maxY,
		startX, startY,
		rune(direction[0]),
		commands,
	)

	fmt.Printf("\nRobot final position: %d %d %c\n", finalX, finalY, finalDir)
}
