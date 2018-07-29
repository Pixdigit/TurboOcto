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

var (
	RELEASED  = buttonState{false, false}
	PRESSING  = buttonState{true, true}
	PRESSED   = buttonState{true, false}
	RELEASING = buttonState{false, true}
)
