package turboOcto

import (
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

type Rect struct {
	geometry.Rect
	// COMBAK: are constrains really a good idea?
	constraint func(*Rect) error
}

func NewRect(p geometry.Point, s geometry.Size, fixPoint geometry.AnchorPoint) (*Rect, error) {
	//Calculate reference after size is defined
	geoR := geometry.NewRect(p, s, fixPoint)
	r := &Rect{}
	r.Rect = geoR

	return r, nil
}
func NewRectFromGeometryRect(r geometry.Rect) (*Rect, error) {
	rect := &Rect{r, func(r *Rect) error { return nil }}
	return rect, nil
}

func (r *Rect) IsClicked(which buttonPosition) (bool, error) {
	return r.Rect.Contains(Mouse.Pos) && (*Mouse.Buttons[which] == PRESSING), nil
}
func (r *Rect) HasMouseState(which buttonPosition, state buttonState) (bool, error) {
	return r.Rect.Contains(Mouse.Pos) && (*Mouse.Buttons[which] == state), nil
}

func (r *Rect) SetConstraint(constraint func(*Rect) error) error {
	r.constraint = constraint
	r.constraint(r)
	return nil
}

func (r *Rect) toSDLRect() (*sdl.Rect, error) {
	size := r.Size()
	topLeft := r.PositionFrom(geometry.TOPLEFT)
	SDLRect := &sdl.Rect{int32(topLeft.X), int32(topLeft.Y), int32(size.Width), int32(size.Height)}
	return SDLRect, nil
}
func (r *Rect) BaseRect() *geometry.Rect {
	return &r.Rect
}
