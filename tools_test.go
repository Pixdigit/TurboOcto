package TurboOcto

import (
    "testing"
)


func TestInSlicce(t *testing.T) {
    ok, err := pathExists("./");    if err != nil {wrapErr(err, "error while checking if local path exists", t)}
    test(ok, "local path not found by pathExists", t)
    ok, err = pathExists("NOPE");    if err != nil {wrapErr(err, "error while checking if local path exists", t)}
    test(!ok, "local path not found by pathExists", t)
}
