# Robot Traversal

## Problem Description

A team of young engineers have built a robot. The robot can travel on a rectangular plane:

### Rectangular Plane
1. The rectangular plane is represented by a X-Y plane with bottom left as (0,0) and top right as (M,N)
2. The rectangular plane contains multiple particles. A particle's position is represented by x and y coordinates

### Robot's Movement
1. Its position is represented by x and y coordinates and a letter representing one of the four directions (N/S/E/W)
2. It takes in commands in the form of single letters. The possible letters are 'L', 'R' and 'M':
   - **'L'** and **'R'** makes the robot turn 90 degrees left or right respectively, without moving from its current position
   - **'M'** means move forward one position and maintain the same direction
3. It stops walking when:
   - It finds any particle on its way
   - It encounters any coordinate on its path which it had travelled earlier
   - The next command leads to a position outside rectangular plane

### Input Format
1. The first line of input is the top-right coordinates of the rectangular plane (M, N)
2. Next two lines of input are about the robot:
   - First line gives the robot's starting position: X Y coordinates and a letter (direction), all 3 separated by spaces
   - Second line is a series of commands for the robot

### Output Format
Robot's final coordinates and direction.

### Example

**Input:**
```
4 4
0 0 N
MMMRMMLM
```

**Output:**
```
2 4 N
```

## Solution Approach

This problem involves:
- Grid traversal and navigation
- State tracking (visited positions for cycle detection)
- Direction management (rotation logic for N/S/E/W)
- Boundary validation
