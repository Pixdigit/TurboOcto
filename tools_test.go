package turboOcto

import (
	"testing"

	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestPathExistance(t *testing.T) {
	ok, err := pathExists("./");	if err != nil {tools.WrapErr(err, "error while checking if local path exists", t)}
	tools.Test(ok, "local path not found by pathExists", t)
	ok, err = pathExists("NOPE");	if err != nil {tools.WrapErr(err, "error while checking if local path exists", t)}
	tools.Test(!ok, "local path not found by pathExists", t)
}
