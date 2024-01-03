package services_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/Fransvis/toyrobot/robot/handlers"
	"github.com/Fransvis/toyrobot/robot/services"
	"github.com/stretchr/testify/assert"
)

func TestInstructRobot(t *testing.T) {
	testCases := []struct {
		name     string
		commands string
		expected services.Robot
		err      error
	}{
		{
			name:     "Valid PLACE and MOVE commands",
			commands: "PLACE 1,2,EAST, MOVE, MOVE, LEFT, MOVE, REPORT",
			expected: services.Robot{X: 3, Y: 3, Direction: services.NORTH, IsPlaced: true},
		},
		{
			name:     "Invalid PLACE command",
			commands: "PLACE",
			expected: services.Robot{X: 0, Y: 0, Direction: services.WEST, IsPlaced: true},
			err:      services.ErrInvalidPlaceCommand,
		},
		{
			name:     "Invalid PLACE coordinates",
			commands: "PLACE 1,2,NORTHEAST, MOVE, MOVE, LEFT, MOVE, REPORT",
			expected: services.Robot{X: 0, Y: 4, Direction: services.WEST, IsPlaced: true},
		},
		{
			name:     "Invalid MOVE command",
			commands: "PLACE 0,0 WEST, MOVE, MOVE, LEFT, MOVE",
			expected: services.Robot{},
			err:      services.ErrInvalidMoveCommand,
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := services.Robot{}
			s := bufio.NewScanner(strings.NewReader(tc.commands))
			err := handlers.InstructRobot(&r, s)

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {

				if r != tc.expected {
					t.Errorf("Expected robot state %v, got %v", tc.expected, r)
				}
			}
		})
	}
}
