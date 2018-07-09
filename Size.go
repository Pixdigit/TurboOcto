package turboOcto

type Size struct {
	Width, Height Scalar
}

func (s *Size) AreaIsPositive() (bool, error) {
	return (s.Width >= 0 && s.Height >= 0), nil
}
