/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */

package contract

import "errors"

var (
	errVmConstructorIsNotPayable = errors.New("constructor is not payable")
)

type ErrSystem interface {
	System() bool
}

func isSystemError(err error) bool {
	sErr, ok := err.(ErrSystem)
	return ok && sErr.System()
}

type vmStartError struct {}

func newVmStartError() error {
	return &vmStartError{}
}

func (e *vmStartError) Error() string {
	return "cannot start a VM"
}

func (e *vmStartError) System() bool {
	return e != nil
}

type DbSystemError struct {
	error
}

func newDbSystemError(e error) error {
	return &DbSystemError{e}
}

func (e *DbSystemError) System() bool {
	return e != nil
}

type ErrRuntime interface {
	Runtime() bool
}

func IsRuntimeError(err error) bool {
	rErr, ok := err.(ErrRuntime)
	return ok && rErr.Runtime()
}

type vmError struct {
	error
}

func newVmError(err error) error {
	return &vmError{err}
}

func (e *vmError) Runtime() bool {
	return e != nil
}
