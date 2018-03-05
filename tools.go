package TurboOcto

import "os"

func inSlice(element interface{}, slice []interface{}) bool {
    for _, v := range(slice) {
        if element == v {
            return true
        }
    }
    return false
}

func pathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
