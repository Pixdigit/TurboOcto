package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/Pixdigit/geometry"
)

func LoadAnimatedSpriteFromTextures(textures []*sdl.Texture, delays []int32) (*Sprite, error) {
	if len(textures) != len(delays) {
		return &Sprite{}, errors.New("argument lengths must be equal \"textures " + string(len(textures)) + "  delays " + string(len(delays)))
	}

	var dimensions []geometry.Size
	newSprite, _ := NewSprite()
	for _, frame := range textures {
		_, _, w, h, err := frame.Query();	if err != nil {return &Sprite{}, errors.Wrap(err, "could not determine texture size")}
		dimensions = append(dimensions, geometry.Size{geometry.Scalar(w), geometry.Scalar(h)})
	}

	newSprite.frames = textures
	newSprite.delays = delays
	newSprite.dimensions = dimensions
	newSprite.FrameIndex = 0
	//Update size
	newSprite.Rect.SetSize(dimensions[0])
	//ensure Sprite has some delay at any frame
	err := newSprite.SetDelay(delays[0]);	if err != nil {return &Sprite{}, errors.Wrap(err, "new Sprite has invalid delay")}
	return newSprite, nil
}

func LoadAnimatedSpriteFromTexture(texture *sdl.Texture) (*Sprite, error) {
	return LoadAnimatedSpriteFromTextures([]*sdl.Texture{texture}, []int32{0})
}

func LoadAnimatedSpriteFromFiles(fileNames []string, delays []int32) (*Sprite, error) {
	var textures []*sdl.Texture
	for _, fileName := range fileNames {
		texture, err := img.LoadTexture(screenRenderer, Cfg.ResourcePath+fileName);	if err != nil {return &Sprite{}, errors.Wrap(err, "could not load Sprite file \""+Cfg.ResourcePath+fileName+"\"")}
		textures = append(textures, texture)
	}
	return LoadAnimatedSpriteFromTextures(textures, delays)
}

func LoadSpriteFromFile(filename string) (*Sprite, error) {
	return LoadAnimatedSpriteFromFiles([]string{filename}, []int32{0})
}

func LoadAnimatedSpriteFromFile(filename string, rects []sdl.Rect, delays []int32) (*Sprite, error) {
	surface, err := img.Load(filename);	if err != nil {return &Sprite{}, errors.Wrap(err, "could not load Sprite image")}
	if len(rects) == 0 {
		//D == Amount of
		DSprites := surface.W / surface.H
		for i := int32(0); i < DSprites; i++ {
			rects = append(rects, sdl.Rect{i * surface.H, 0, surface.H, surface.H})
		}
	}
	var textures []*sdl.Texture
	xOffset := int32(0)
	for _, rect := range rects {
		if rect.W == 0 || rect.H == 0 {
			rect = sdl.Rect{0, 0, surface.H, surface.H}
		}
		tmpSurface, err := sdl.CreateRGBSurface(0, rect.W, rect.H, 32, rmask, gmask, bmask, amask);	if err != nil {return &Sprite{}, errors.Wrap(err, "could not create tmpSurface for transfer")}
		rect.X += xOffset
		xOffset += rect.W
		surface.Blit(&rect, tmpSurface, nil)
		texture, err := screenRenderer.CreateTextureFromSurface(tmpSurface);	if err != nil {return &Sprite{}, errors.Wrap(err, "could not transfer surface to texture")}
		textures = append(textures, texture)
	}
	return LoadAnimatedSpriteFromTextures(textures, delays)
}
