package lang

import (
	"strings"
	"testing"

	"github.com/roman-mazur/architecture-lab-3/painter"
	"github.com/stretchr/testify/assert"
)

type TCRecognOp struct {
	test      string
	input     string
	operation painter.Operation
}

type TCMultipleOps struct {
	test       string
	input      string
	operations []painter.Operation
}

type TCInvalidInput struct {
	test  string
	input string
	err   error
}

func TestRecognizeOperations(t *testing.T) {
	testCases := []TCRecognOp{
		{
			test:      "green",
			input:     "green",
			operation: painter.GreenFill{},
		}, {
			test:      "figure",
			input:     "figure 0.5 0.4",
			operation: painter.AddT{},
		}, {
			test:      "move",
			input:     "move 0.5 0.2",
			operation: painter.MoveAll{},
		}, {
			test:      "bgrect",
			input:     "bgrect 0.3 0.2 0.5 0.7",
			operation: painter.BgRect{},
		}, {
			test:      "reset",
			input:     "reset",
			operation: painter.Reset{},
		}, {
			test:      "white",
			input:     "white",
			operation: painter.WhiteFill{},
		}, {
			test:      "update",
			input:     "update",
			operation: painter.UpdateOp,
		},
	}
	parser := Parser{}
	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		op, err := parser.Parse(strings.NewReader(testCase.input))
		if err != nil {
			t.Fatalf("Operation %s wasn't recognized", testCase.test)
		}
		assert.IsType(t, testCase.operation, op[0])
	}
}

func TestMultipleOperations(t *testing.T) {
	testCases := []TCMultipleOps{
		{
			test:       "multiple",
			input:      "green\nwhite\nfigure 0.5 0.5\nbgrect 0.2 0.2 0.3 0.3\nupdate",
			operations: []painter.Operation{painter.GreenFill{}, painter.WhiteFill{}, painter.AddT{}, painter.BgRect{}, painter.UpdateOp},
		},
	}
	parser := Parser{}
	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		ops, err := parser.Parse(strings.NewReader(testCase.input))
		if err != nil {
			t.Fatalf("Error accured in %s test: %s", testCase.test, err.Error())
		}
		for i := range ops {
			assert.IsType(t, testCase.operations[i], ops[i])
		}
	}
}

func TestInvalidInputs(t *testing.T) {
	testCases := []TCInvalidInput{
		{
			test:  "multiple",
			input: "fdgdfgfdgdf",
			err:   UnknownOperationError{},
		}, {
			test:  "multiple",
			input: "bgrect 3 4",
			err:   NotEnoughArgsError{},
		}, {
			test:  "multiple",
			input: "bgrect fsdfds 4 fsd 3",
			err:   InvalidArgsError{},
		},
	}
	parser := Parser{}
	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		_, err := parser.Parse(strings.NewReader(testCase.input))
		if err == nil {
			t.Fatalf("Error wasn't catch in test %s", testCase.test)
		}
		assert.IsType(t, testCase.err, err)
	}
}
