package services

import (
	"errors"
	"fmt"
)

type Direction int

// Robot is a struct that holds the current position and facing direction of the robot
type Robot struct {
	X, Y      int
	Direction Direction
	IsPlaced  bool
}

var ErrInvalidPlaceCommand = errors.New("please provide a valid PLACE command")
var ErrInvalidMoveCommand = errors.New("be careful or you will fall to your doom")

// Directions are the possible facing directions of the robot
// An iota is used to increase the value of each constant by 1 -
// This can become messy if the list evolves to much, but it is a nice way to start
const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (r *Robot) Place(x, y int, f Direction) error {
	if x < 0 || x > 4 || y < 0 || y > 4 {
		return ErrInvalidPlaceCommand
	}

	r.X = x
	r.Y = y
	r.Direction = f
	r.IsPlaced = true

	return nil
}

func (r *Robot) Move() error {
	newX, newY := r.X, r.Y
	switch r.Direction {
	case NORTH:
		newY++
	case EAST:
		newX++
	case SOUTH:
		newY--
	case WEST:
		newX--
	}

	if newX < 0 || newX > 4 || newY < 0 || newY > 4 {
		return fmt.Errorf("%w", ErrInvalidMoveCommand)
	}

	r.X, r.Y = newX, newY

	return nil
}

func (r *Robot) Left() {
	r.Direction = (r.Direction + 3) % 4
}

func (r *Robot) Right() {
	r.Direction = (r.Direction + 1) % 4
}

func (r *Robot) Report() {
	Directions := []string{"NORTH", "EAST", "SOUTH", "WEST"}
	fmt.Printf("Your current position is %d,%d,%s\n", r.X, r.Y, Directions[r.Direction])
}
