package test

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"github.com/roman-mazur/architecture-lab-3/painter/lang"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TestCase struct {
	input     string
	command   string
	operation painter.Operation
}

func TestMain(t *testing.T) {
	testCases := []TestCase{
		{
			input:     "white fill",
			command:   "white",
			operation: painter.OperationFunc(painter.WhiteFill),
		}, {
			input:     "green fill",
			command:   "green",
			operation: painter.OperationFunc(painter.GreenFill),
		}, {
			input:     "update",
			command:   "update",
			operation: painter.UpdateOp,
		}, {
			input:     "move 12 12",
			command:   "move",
			operation: painter.MoveAll(12, 12),
		},
	}
	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		testParserTestCase(testCase, t)
	}
}

func testParserTestCase(testCase TestCase, t *testing.T) {
	parser := lang.Parser{}
	t.Run(testCase.input, func(t *testing.T) {
		op, _ := parser.Parse(strings.NewReader(testCase.input))

		assert.IsType(t, testCase.operation, op[0])
	})

}
