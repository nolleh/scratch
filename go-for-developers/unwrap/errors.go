package error

import "errors"
import "fmt"

type ErrorA struct {
	err error
}

func (e ErrorA) Error() string {
	return fmt.Sprintf("ErrorA %s", e.err)
}

func (e ErrorA) Unwrap() error {
	if _, ok := e.err.(Unwrapper); ok {
		return errors.Unwrap(e.err)
	}
	return e.err
}

type ErrorB struct {
	err error
}

func (e ErrorB) Error() string {
	return fmt.Sprintf("[ErrorB] %s", e.err)
}

func (e ErrorB) Unwrap() error {
	if _, ok := e.err.(Unwrapper); ok {
		return errors.Unwrap(e.err)
	}
	return e.err
}

type ErrorC struct {
	err error
}

func (e ErrorC) Error() string {
	return fmt.Sprintf("[ErrorC] %s", e.err)
}

func (e ErrorC) Unwrap() error {
  if _, ok := e.err.(Unwrapper); ok {
    return errors.Unwrap(e.err)
  }
  return e.err
}
