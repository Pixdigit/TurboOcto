package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "github.com/pkg/errors"
)

type Sprite struct {
    frames  []sdl.Texture
    delays  []int32
    XCenter, YCenter int32
}

func NewSprite() (Sprite, error) {
    return Sprite{}, nil
}
func LoadAnimatedSpriteFromTextures(textures []sdl.Texture, delays []int32) (Sprite, error) {
    if (len(textures) != len(delays)) {return Sprite{}, errors.New("argument lengths must be equal \"textures " + string(len(textures)) + "  delays " + string(len(delays)))}

    return Sprite{
        frames: textures,
        delays: delays,
    }, nil
}
func LoadAnimatedSpriteFromFiles(fileNames []string, delays []int32) (Sprite, error) {
    var textures []sdl.Texture
    for _, fileName := range fileNames {
        texture, err := img.LoadTexture(renderer, "./assets/sprites/" + fileName);    if err != nil {return Sprite{}, errors.Wrap(err, "could not load sprite file \"./assets/sprites/" + fileName)}
        textures = append(textures, *texture)
    }
    return LoadAnimatedSpriteFromTextures(textures, delays)
}
func LoadSpriteFromFile(filename string) (Sprite, error) {
    return LoadAnimatedSpriteFromFiles([]string{filename}, []int32{0})
}
//IDEA: Support GIF as textures and delays source
func LoadAnimatedSpriteFromFile(filename string, rects []sdl.Rect, delays []int32) (Sprite, error) {
    surface, err := img.Load(filename);    if err != nil {return Sprite{}, errors.Wrap(err, "could not load sprite image")}
    if len(rects) == 0 {
        //D == Amount of
        DSprites := surface.W / surface.H
        for i := int32(0); i < DSprites; i++ {
            rects = append(rects, sdl.Rect{i * surface.H, 0, surface.H, surface.H})
        }
    }
    var textures []sdl.Texture
    xOffset := int32(0)
    for _, rect := range rects {
        if rect.W == 0 || rect.H == 0 {
            rect = sdl.Rect{0, 0, surface.H, surface.H}
        }
        tmpSurface, err := sdl.CreateRGBSurface(0, rect.W, rect.H, 32, rmask, gmask, bmask, amask);    if err != nil {return Sprite{}, errors.Wrap(err, "could not create tmpSurface for transfer")}
        rect.X += xOffset
        xOffset += rect.W
        surface.Blit(&rect, tmpSurface, nil)
        texture, err := renderer.CreateTextureFromSurface(tmpSurface);    if err != nil {return Sprite{}, errors.Wrap(err, "could not transfer surface to texture")}
        textures = append(textures, *texture)

    }
    return Sprite{
        frames: textures,
        delays: delays,
        }, nil
}
