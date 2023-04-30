package lang

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	var res []painter.Operation
	for scanner.Scan() {
		commandLine := scanner.Text()
		op := getOperationFromString(commandLine) // parse the line to get Operation
		res = append(res, op)
	}

	return res, nil
}

func getOperationFromString(commandString string) painter.Operation {
	splitCommandLine := strings.Fields(commandString)
	command := splitCommandLine[0]

	fmt.Printf(command)
	switch command {
	case "white":
		break
	case "green":
		break
	case "update":
		break
	case "bgrect":
		break
	case "figure":
		break
	case "move":
		break
	case "reset":
		break
	default:
		//
		break
	}

	return nil
}
