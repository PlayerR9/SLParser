package Test

import (
	"testing"

	prx "github.com/PlayerR9/SLParser"
)

func TestGeneration(t *testing.T) {
	_, err := prx.ParseEbnf([]byte("Source = typeID space assignment_op space op_square variable_id space asterisk_op space variable_id space for space variable_id space in space range op_paren number cl_paren cl_square newline typeID EOF ."))
	if err != nil {
		t.Error(err)
	}
}
