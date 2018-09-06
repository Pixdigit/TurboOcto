package turboOcto

type buttonState struct {
	isDown  bool
	changed bool
}

func (b *buttonState) update(isDown bool) {
	if b.isDown != isDown {
		b.changed = true
	} else {
		b.changed = false
	}
	b.isDown = isDown
}
func (b *buttonState) copy() *buttonState {
	return &buttonState{b.isDown, b.changed}
}

func (b *buttonState) equals(b2 buttonState) bool {
	return b.isDown == b2.isDown && b.changed == b2.changed
}

func (b *buttonState) Is(b2 buttonState) (bool, error) {
	if b == nil {
        //Default state is RELEASED
		if b2.equals(RELEASED) {
			return true, nil
		} else {
			return false, nil
		}
	}
	return b.equals(b2), nil
}

var (
	RELEASED  = buttonState{false, false}
	PRESSING  = buttonState{true, true}
	PRESSED   = buttonState{true, false}
	RELEASING = buttonState{false, true}
)
