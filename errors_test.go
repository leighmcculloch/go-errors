package errors_test

import (
	"strconv"
	"testing"

	"4d63.com/errors"
	"4d63.com/test"
)

type nonComparableError []struct{}

func (nonComparableError) Error() string {
	return "nonComparableError"
}

type comparableError []struct{}

func (comparableError) Equal(err error) bool {
	_, ok := err.(comparableError)
	return ok
}

func (comparableError) Error() string {
	return "comparableError"
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		Err    error
		Target error
		Equal  bool
	}{
		{nil, nil, true},
		{errors.New("a"), errors.New("a"), false},
		{errors.New("a"), errors.New("b"), false},
		{nonComparableError{}, nonComparableError{}, false},
		{comparableError{}, comparableError{}, true},
		// TODO: Unwrap
		// TODO: All types
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			equal := errors.Equal(tc.Err, tc.Target)
			test.Eq(t, equal, tc.Equal)
		})
	}
}
