package ast

import (
	gers "github.com/PlayerR9/go-errors"
	gerr "github.com/PlayerR9/go-errors/error"
)

// TODO: Delete this once go-errors is updated.

// NewErrNilReceiver returns a new error.Err error representing a
// nil receiver.
//
// Returns:
//   - *error.Err: A pointer to the newly created error. Never returns nil.
func NewErrNilReceiver() *gerr.Err {
	err := gerr.New(gers.OperationFail, "receiver must not be nil")
	err.AddSuggestion("Did you forget to initialize the receiver?")

	return err
}
