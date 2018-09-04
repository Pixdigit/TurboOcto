package turboOcto

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	geo "gitlab.com/Pixdigit/geometry"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestFontRenderMethods(t *testing.T) {
	SetWindowSize(geo.Size{500, 500})
	Windowed()
	screenRenderer.SetDrawColor(0, 0, 0, 255)
	screenRenderer.Clear()
	screenRenderer.Present()

	testText := "blablabla"
	f, err := OpenFont("./assets/fonts/aileron/Aileron-Black.otf", 28);	if err != nil {tools.WrapErr(err, "could not open font", t)}
	f.ColorRed = 0
	f.ColorGreen = 255
	f.ColorBlue = 255
	f.ColorAlpha = 255
	texture, err := f.RenderString(testText);	if err != nil {tools.WrapErr(err, "could not render blended with font", t)}
	size, err := f.TextSize(testText);	if err != nil {tools.WrapErr(err, "could not determine size of rendered text", t)}
	rect := sdl.Rect{0, 0, int32(size.Width), int32(size.Height)}
	screenRenderer.Copy(texture, nil, &rect)

	f.Method = SHADED
	f.ColorRed = 255
	f.ColorGreen = 0
	f.ColorBlue = 255
	f.ColorAlpha = 180
	texture, err = f.RenderString(testText);	if err != nil {tools.WrapErr(err, "could not render shaded with font", t)}
	size, err = f.TextSize(testText);	if err != nil {tools.WrapErr(err, "could not determine size of rendered text", t)}
	rect = sdl.Rect{0, 100, int32(size.Width), int32(size.Height)}
	screenRenderer.Copy(texture, nil, &rect)

	f.Method = SOLID
	f.ColorRed = 255
	f.ColorGreen = 255
	f.ColorBlue = 0
	f.ColorAlpha = 100
	texture, err = f.RenderString(testText);	if err != nil {tools.WrapErr(err, "could not render solid with font", t)}
	size, err = f.TextSize(testText);	if err != nil {tools.WrapErr(err, "could not determine size of rendered text", t)}
	rect = sdl.Rect{0, 200, int32(size.Width), int32(size.Height)}
	screenRenderer.Copy(texture, nil, &rect)
	screenRenderer.Present()
    if !testing.Short() {
        window.SetTitle("Check fonts")
        sdl.Delay(3000)
    }
}
