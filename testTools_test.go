package TurboOcto

import (
    "testing"
    "github.com/pkg/errors"
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
        t.Error(errMsg)
    }
}

func testAgainstStrings(set func(s string)(error), get func()(string), errMsg string, t *testing.T) {
    for _, testString := range(testStrings) {
        err := set(testString);    if err != nil {t.Error(errors.Wrap(err, "failed to set string"))}
        //TODO: check for errs while getting
        result := get()
        errorMsg := errMsg + ": failed at string \"" + testString + "\"; is " + result
        test(result == testString, errorMsg, t)
    }
}
