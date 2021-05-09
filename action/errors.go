package action

import "errors"

var (
	errNotImplemented  = errors.New("not implemented")
	errUnknownActionID = errors.New("unknown action id")
)
