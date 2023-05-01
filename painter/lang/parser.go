package lang

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/GAlexandrD/architecture-lab-3/painter"
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
		op, err := getOperationFromString(commandLine) // parse the line to get Operation
		if err != nil {
			return nil, err
		}
		res = append(res, op)
	}

	return res, nil
}

func getOperationFromString(commandString string) (painter.Operation, error) {
	splitCommandLine := strings.Fields(commandString)
	command := splitCommandLine[0]
	strArgs := splitCommandLine[1:]

	switch command {
	case "white":
		return painter.WhiteFill{}, nil
	case "green":
		return painter.GreenFill{}, nil
	case "update":
		return painter.UpdateOp, nil
	case "bgrect":
		args, err := checkAndParseArgs(4, strArgs)
		if err != nil {
			return nil, err
		}
		return painter.BgRect{X1: args[0], Y1: args[1], X2: args[2], Y2: args[3]}, nil
	case "figure":
		args, err := checkAndParseArgs(2, strArgs)
		if err != nil {
			return nil, err
		}
		return painter.AddT{X: args[0], Y: args[1]}, nil
	case "move":
		args, err := checkAndParseArgs(2, strArgs)
		if err != nil {
			return nil, err
		}
		return painter.MoveAll{X: args[0], Y: args[1]}, nil
	case "reset":
		return painter.Reset{}, nil
	}
	return nil, UnknownOperationError{}
}

func checkAndParseArgs(count int, argsStr []string) ([]float32, error) {
	if len(argsStr) != count {
		return nil, NotEnoughArgsError{}
	}
	args, err := parseStrToFloat(argsStr)
	if err != nil {
		return nil, InvalidArgsError{}
	}
	return args, nil
}

func parseStrToFloat(strArr []string) ([]float32, error) {
	floatArr := make([]float32, len(strArr))
	for i := range strArr {
		float, err := strconv.ParseFloat(strArr[i], 32)
		if err != nil {
			return nil, err
		}
		floatArr[i] = float32(float)
	}
	return floatArr, nil
}

type InvalidArgsError struct{}

func (e InvalidArgsError) Error() string {
	return "Invalid arguments"
}

type NotEnoughArgsError struct{}

func (e NotEnoughArgsError) Error() string {
	return "Not enough arguments provided"
}

type UnknownOperationError struct{}

func (e UnknownOperationError) Error() string {
	return "Unknown Operation"
}
