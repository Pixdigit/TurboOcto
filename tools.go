package TurboOcto

import (
    "os"
)

type Runlevel int32
const STOPPED Runlevel = 0;
const RUNNING Runlevel = 1;
const PAUSED Runlevel = 2;

func pathExists(path string) (bool, error) {
    _, err := os.Stat(path);    if err == nil {return true, nil}
    if os.IsNotExist(err) {return false, nil}
    return true, err
}

func pushErr(errChan chan error, err error) {
    go func() {
        errChan <- err
    }()
}
