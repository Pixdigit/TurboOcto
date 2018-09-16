package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	geo "gitlab.com/Pixdigit/geometry"
)

type Frame struct {
	*geo.Rect
	*sdl.Texture
	Visible bool
}

func NewFrame(texture *sdl.Texture) (*Frame, error) {
	_, _, w, h, err := texture.Query();	if err != nil {return nil, errors.Wrap(err, "Could not queury info for new Frame")}
	size := geo.Size{geo.Scalar(w), geo.Scalar(h)}
	rect := geo.NewRect(geo.Point{0, 0}, size, geo.CENTER)
	visible := Cfg.FramesVisibleOnLoad

	frame := &Frame{
		&rect,
		texture,
		visible,
	}

	return frame, nil
}

func NewEmptyFrame() (*Frame, error) {
	//size needs to be at least 1
	surf, err := sdl.CreateRGBSurface(0, 1, 1, 32, rmask, gmask, bmask, amask)
	//r = g = b = alpha = 0
	surf.FillRect(nil, 0);	if err != nil {return nil, errors.Wrap(err, "could not create dummy pixel data")}
	texture, err := screenRenderer.CreateTextureFromSurface(surf)
	frame, err := NewFrame(texture);	if err != nil {return nil, errors.Wrap(err, "could not create empty frame for new Sprite")}
	return frame, nil
}

func (f *Frame) Render() error {
	if !f.Visible {
		return nil
	}
	size := f.Size()
	topLeft := f.PositionFrom(geo.TOPLEFT)
	dstRect := &sdl.Rect{int32(topLeft.X), int32(topLeft.Y), int32(size.Width), int32(size.Height)}
	err := screenRenderer.Copy(f.Texture, nil, dstRect);	if err != nil {return errors.Wrap(err, "could not copy Sprite frame to screenRenderer")}
	return nil
}

func (f *Frame) RenderToFrame(dstFrame *Frame) error {
	if !f.Visible {
		return nil
	}
	size := f.Size()
	topLeft := f.PositionFrom(geo.TOPLEFT)
	dstRect := &sdl.Rect{int32(topLeft.X), int32(topLeft.Y), int32(size.Width), int32(size.Height)}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_TARGETTEXTURE);	if err != nil {return errors.Wrap(err, "could not create renderer to render to texture")}
	renderer.SetRenderTarget(dstFrame.Texture)
	err = renderer.Copy(f.Texture, nil, dstRect);	if err != nil {return errors.Wrap(err, "could not copy Sprite frame to texture")}
	return nil
}
