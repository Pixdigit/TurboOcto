package sharedStructs

import (
    "github.com/ByteArena/box2d"
	"github.com/veandco/go-sdl2/sdl"
    //"github.com/pkg/errors"
)

type Sprite struct {
    Name string
    Texture *sdl.Texture
    Rect *sdl.Rect
    Body *box2d.B2Body
}
