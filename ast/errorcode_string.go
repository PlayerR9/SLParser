// Code generated by "stringer -type=ErrorCode"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnregisteredType-0]
	_ = x[BadSyntaxTree-1]
}

const _ErrorCode_name = "UnregisteredTypeBadSyntaxTree"

var _ErrorCode_index = [...]uint8{0, 16, 29}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}
