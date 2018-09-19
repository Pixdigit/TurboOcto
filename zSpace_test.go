package turboOcto

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestZSpace(t *testing.T) {
	// LoadDefaultConf()
	surf, err := sdl.CreateRGBSurface(0, 1, 1, 32, rmask, gmask, bmask, amask);	if err != nil {tools.WrapErr(err, "could not create a new pixel buffer", t)}
	surf.FillRect(nil, sdl.Color{0, 0, 0, 0}.Uint32());	if err != nil {tools.WrapErr(err, "could not create dummy pixel data", t)}
	frame1, err := NewEmptyFrame();	if err != nil {tools.WrapErr(err, "could not create empty frame for new Sprite", t)}
	frame2, err := NewEmptyFrame();	if err != nil {tools.WrapErr(err, "could not create empty frame for new Sprite", t)}
	frame3, err := NewEmptyFrame();	if err != nil {tools.WrapErr(err, "could not create empty frame for new Sprite", t)}
	frame4, err := NewEmptyFrame();	if err != nil {tools.WrapErr(err, "could not create empty frame for new Sprite", t)}
	AddElement(frame1, 0)
	AddElement(frame2, 1)
	AddElement(frame3, 2)
	AddElement(frame4, 3)
	_ = frame2
	_ = frame3
	_ = frame4
	// sp, err := NewSprite();	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	// _, _ = sp, err
	// sp2, err := NewSprite();	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	// //create unreferenced sprite
	// _, err = NewSprite();	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	// //TODO:
	// // err = sp2.ChangeLayer(2);	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	// // err = sp2.ChangeLayer(3);	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	//
	// //create unreferenced sprite
	// _, err = NewSprite();	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	//
	// tools.Test(sprites[0] == sp, "sprite did not remain in position", t)
	// tools.Test(len(sprites) == 4, "unexpected sprite count", t)
	// tools.Test(sprites[len(sprites)-1] == sp2, "sprite on layer 3 did not move to correct position", t)
}
