package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

type Rect struct {
	geometry.Rect
}

func NewRect(p geometry.Point, s geometry.Size, fixPoint geometry.AnchorPoint) (*Rect, error) {
	//Calculate reference after size is defined
	geoR, err := geometry.NewRect(p, s, fixPoint);	if err != nil {return &Rect{}, errors.Wrap(err, "unable to create a new Rect")}
	r := &Rect{}
	r.Rect = *geoR

	return r, nil
}
func NewRectFromGeometryRect(r geometry.Rect) (*Rect, error) {
	rect := &Rect{r}
	return rect, nil
}

func (r *Rect) DrawBoundaries(red, g, b, a uint8) error {
	SDLRect, err := r.SDLRect();	if err != nil {return errors.Wrap(err, "could not Rect boundaries for drawing them")}
	err = screenRenderer.SetDrawColor(red, g, b, a);	if err != nil {return errors.Wrap(err, "could not set draw color")}
	err = screenRenderer.DrawRect(SDLRect);	if err != nil {return errors.Wrap(err, "could not draw Rect")}
	return nil
}

func (r *Rect) SDLRect() (*sdl.Rect, error) {
	size, err := r.Size();	if err != nil {return &sdl.Rect{}, errors.Wrap(err, "could not get sprite size for blitting")}
	topLeft, err := r.AnchorPosition(geometry.AnchorPoint{geometry.LEFT, geometry.TOP});	if err != nil {return &sdl.Rect{}, errors.Wrap(err, "could not get sprite position for blitting")}
	SDLRect := &sdl.Rect{int32(topLeft.X), int32(topLeft.Y), int32(size.Width), int32(size.Height)}
	return SDLRect, nil
}
func (r *Rect) BaseRect() (*geometry.Rect, error) {
	return &r.Rect, nil
}
