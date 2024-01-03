package handlers

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/Fransvis/toyrobot/robot/services"
)

func InstructRobot(r *services.Robot, s *bufio.Scanner) error {
	for s.Scan() {
		line := s.Text()
		commands := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ','
		})

		if !r.IsPlaced && !checkValidPlaceCommand(commands) {
			return services.ErrInvalidPlaceCommand
		}

		for i := 0; i < len(commands); {
			commands[i] = strings.ToUpper(commands[i]) // ensure safety for upppercase and lowercase
			// TODO(FRANSVIS): implement logic for typos?
			switch commands[i] {
			case "PLACE":
				if i+3 < len(commands) {
					x, _ := strconv.Atoi(commands[i+1])
					y, _ := strconv.Atoi(commands[i+2])
					var direction services.Direction

					// TODO(Fransvis): Improve
					switch commands[i+3] {
					case "NORTH":
						direction = services.NORTH
					case "SOUTH":
						direction = services.SOUTH
					case "EAST":
						direction = services.EAST
					case "WEST":
						direction = services.WEST
					}

					err := r.Place(x, y, direction)
					if err != nil {
						return err
					}
					i += 3 // Skip the next three commands as they are part of the PLACE command
				}
			case "MOVE":
				err := r.Move()
				if err != nil {
					return err
				}
			case "LEFT":
				r.Left()
			case "RIGHT":
				r.Right()
			case "REPORT":
				r.Report()
			}
			i++
		}
	}

	return nil
}

func checkValidPlaceCommand(commands []string) bool {
	if len(commands) < 4 {
		return false
	}

	return true
}
