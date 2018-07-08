package turboOcto

func min(a, b int32) (int32, error) {
	if a < b {
		return a, nil
	} else {
		return b, nil
	}
}
func max(a, b int32) (int32, error) {
	if a > b {
		return a, nil
	} else {
		return b, nil
	}
}
