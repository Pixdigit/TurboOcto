package turboOcto

import (
	"github.com/pkg/errors"
	"testing"
)

var testStrings []string = []string{"TEST", "ẞönDérZäíſĉh€Ń", "1234567890", "", "\n"}

func assert(success bool) error {
	if !success {
		return errors.New("assertion failed!")
	} else {
		return nil
	}
}

func test(success bool, errMsg string, t *testing.T) {
	err := assert(success)
	if err != nil {
		wrapErr(err, errMsg, t)
	}
}

func testAgainstStrings(set func(s string) error, get func() (error, string), errMsg string, t *testing.T) {
	for _, testString := range testStrings {
		err := set(testString)
		if err != nil {
			wrapErr(err, "failed to set string", t)
		}
		err, result := get()
		if err != nil {
			wrapErr(err, "failed to get string", t)
		}
		errorMsg := errMsg + ": failed at string \"" + testString + "\"; is " + result
		test(result == testString, errorMsg, t)
	}
}

func wrapErr(err error, msg string, t *testing.T) {
	t.Error(errors.Wrap(err, msg))
}
