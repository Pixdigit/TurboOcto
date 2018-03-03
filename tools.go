package TurboOcto


func inSlice(element interface{}, slice []interface{}) bool {
    for _, v := range(slice) {
        if element == v {
            return true
        }
    }
    return false
}
