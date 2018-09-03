package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	geo "gitlab.com/Pixdigit/geometry"
)

func initializeTextProcessing() error {
	return ttf.Init()
}

type renderMethod int

const (
	BLENDED = renderMethod(iota)
	SHADED
	SOLID
)

type GlyphMetrics struct {
	ttf.GlyphMetrics
}

type Font struct {
	*ttf.Font
	Method     renderMethod
	ColorRed   uint8
	ColorGreen uint8
	ColorBlue  uint8
	ColorAlpha uint8
}

func colorToSDLColor(r, g, b, a uint8) sdl.Color {
	return sdl.Color{r, g, b, a}
}

func OpenFont(file string, size int) (Font, error) {
	sdlFont, err := ttf.OpenFont(file, size)
	f := Font{
		sdlFont,
		BLENDED,
		255,
		255,
		255,
		255,
	}
	return f, err
}

func (f *Font) RenderString(text string) (*sdl.Texture, error) {
	color := colorToSDLColor(f.ColorRed, f.ColorGreen, f.ColorBlue, f.ColorAlpha)

	var fontSurf *sdl.Surface
	var err error
	switch f.Method {
	case BLENDED:
		fontSurf, err = f.RenderUTF8Blended(text, color);	if err != nil {return nil, errors.Wrap(err, "could not render text blended")}
	case SHADED:
		fontSurf, err = f.RenderUTF8Shaded(text, color, sdl.Color{0, 0, 0, 0});	if err != nil {return nil, errors.Wrap(err, "could not render text shaded")}
	case SOLID:
		fontSurf, err = f.RenderUTF8Solid(text, color);	if err != nil {return nil, errors.Wrap(err, "could not render text solid")}
	default:
		return nil, errors.New("render method does not exist")
	}

	fontTexture, err := screenRenderer.CreateTextureFromSurface(fontSurf);	if err != nil {return nil, errors.Wrap(err, "could not convert render output to texture")}

	return fontTexture, nil
}

func (f *Font) TextSize(text string) (geo.Size, error) {
	w, h, err := f.SizeUTF8(text);	if err != nil {return geo.Size{}, errors.Wrap(err, "could not texture size of text")}
	return geo.Size{geo.Scalar(w), geo.Scalar(h)}, nil
}

func (f *Font) GlyphMetrics() (*GlyphMetrics, error) {
	metrics, err := f.GlyphMetrics();	if err != nil {return nil, errors.Wrap(err, "could not get metrics")}
	turboOctoGlyphMetrics := GlyphMetrics(*metrics)
	return &turboOctoGlyphMetrics, nil
}
