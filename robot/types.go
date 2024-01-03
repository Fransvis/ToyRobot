package types

type Direction int

// Robot is a struct that holds the current position and facing direction of the robot
type Robot struct {
	X, Y int
	F    Direction
}

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)
