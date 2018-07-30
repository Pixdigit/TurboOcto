package turboOcto

type buttonState struct {
	IsDown  bool
	Changed bool
}

func (b *buttonState) update(isDown bool) {
	if b.IsDown != isDown {
		b.Changed = true
	} else {
		b.Changed = false
	}
	b.IsDown = isDown
}
func (b *buttonState) copy() *buttonState {
	return &buttonState{b.IsDown, b.Changed}
}

var (
	RELEASED  = buttonState{false, false}
	PRESSING  = buttonState{true, true}
	PRESSED   = buttonState{true, false}
	RELEASING = buttonState{false, true}
)
