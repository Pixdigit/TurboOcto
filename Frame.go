package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	geo "gitlab.com/Pixdigit/geometry"
	"gitlab.com/Pixdigit/uniqueID"
)

type Frame struct {
	*geo.Rect
	*sdl.Texture
	id      uniqueID.ID
	Visible bool
}

func NewFrame(texture *sdl.Texture) (*Frame, error) {
	_, _, w, h, err := texture.Query();	if err != nil {return nil, errors.Wrap(err, "Could not queury info for new Frame")}
	size := geo.Size{geo.Scalar(w), geo.Scalar(h)}
	rect := geo.NewRect(geo.Point{0, 0}, size, geo.CENTER)
	visible := Cfg.FramesVisibleOnLoad

	ID := uniqueID.NewID()

	frame := &Frame{
		&rect,
		texture,
		ID,
		visible,
	}

	return frame, nil
}

func NewEmptyFrame() (*Frame, error) {
	//size needs to be at least 1
	surf, err := sdl.CreateRGBSurface(0, 1, 1, 32, rmask, gmask, bmask, amask);	if err != nil {return &Frame{}, errors.Wrap(err, "could not create a new pixel buffer")}
	//r = g = b = alpha = 0
	surf.FillRect(nil, sdl.Color{0, 0, 0, 0}.Uint32());	if err != nil {return nil, errors.Wrap(err, "could not create dummy pixel data")}
	texture, err := screenRenderer.CreateTextureFromSurface(surf);	if err != nil {return &Frame{}, errors.Wrap(err, "could not copy pixel buffer into frame")}
	frame, err := NewFrame(texture);	if err != nil {return nil, errors.Wrap(err, "could not create empty frame for new Sprite")}
	return frame, nil
}

func (f *Frame) ID() uniqueID.ID {
	return f.id
}

func (f *Frame) render() error {
	if !f.Visible {
		return nil
	}
	size := f.Size()
	topLeft := f.PositionFrom(geo.TOPLEFT)
	dstRect := &sdl.Rect{int32(topLeft.X), int32(topLeft.Y), int32(size.Width), int32(size.Height)}
	err := screenRenderer.Copy(f.Texture, nil, dstRect);	if err != nil {return errors.Wrap(err, "could not copy Sprite frame to screenRenderer")}
	return nil
}
