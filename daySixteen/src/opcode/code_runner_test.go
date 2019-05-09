package opcode_test

import (
	"opcode"
	"testing"
)

func TestOpCodeRunner(t *testing.T) {
	val := opcode.Parse("../input1.txt")
	num := opcode.TryCodesAndGetCountOfMatches(val[0])
	t.Error(num)
}
