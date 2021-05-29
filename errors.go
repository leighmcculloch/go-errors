package errors

import "errors"

// Assign checks if err is assignable to the target value.
func Assign(err error, target interface{}) bool {
	return errors.As(err, target)
}

// EquatableError is an error that can be compared for equality against other
// errors of the same or different type.
type EquatableError interface {
	error
	Equal(err error) bool
}

// Equal checks if the err is equal to target.
//
// If target has an Equal function, satisfing the EquatableError interface, it
// is called.  Otherwise a basic equality check takes place.
func Equal(err, target error) bool {
	if err == nil && target == nil {
		return true
	}
	eqTarget, equatable := target.(EquatableError)
	for {
		if equatable && eqTarget.Equal(err) {
			return true
		}
		if func() bool {
			defer func() {
				recover()
			}()
			return err == target
		}() {
			return true
		}
		err = errors.Unwrap(err)
		if err == nil {
			return false
		}
	}
}

// MatchableError is an error that can be partially matched against other errors
// of the same or different type.
type MatchableError interface {
	error
	Match(err error) bool
}

// Match checks if the err is matchable to the target.
//
// The target must be an error with a Match function, satisfying the
// MatchableError interface.
func Match(err, target MatchableError) bool {
	return errors.Is(err, target)
}

func New(text string) error {
	return errors.New(text)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}
