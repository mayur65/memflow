package protocol

import (
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func ParseCommand(line string) (Command, error) {

	parts := strings.Split(strings.TrimSpace(line), " ")

	//String should have at least 1 space
	if len(parts) < 2 {
		return Command{}, nil
	}

	cmd := Command{Name: parts[0], Args: parts[1:]}

	return cmd, nil
}
