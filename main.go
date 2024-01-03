package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Fransvis/toyrobot/robot/handlers"
	"github.com/Fransvis/toyrobot/robot/services"
)

// PLACE X,Y,F: where X is x-axis, Y is y-axis and F is facing direction
// MOVE: moves one forward in the direction it is facing
// LEFT: rotates 90 degrees to the left
// RIGHT: rotates 90 degrees to the right
// REPORT: outputs the current position of the robot

func main() {
	r := services.Robot{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Welcome to the Toy Robot Simulator!\n")

	for {
		err := handlers.InstructRobot(&r, scanner)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue // Move to the next iteration of the loop - this is to prevent the program from exiting
		}
	}

}
