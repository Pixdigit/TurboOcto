package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "github.com/pkg/errors"
)

type Sprite struct {
    Frames  []*sdl.Texture
    Delays  []int32
    dimensionsHalved [][2]int32
    FrameIndex int32
    XCenter, YCenter int32
    layer int32
}

var sprites []*Sprite

func NewSprite() (*Sprite, error) {
    sprite := &Sprite{}
    //ensure sprite is the topmost of level 0
    sprites = append([]*Sprite{sprite}, sprites...)
    sprite.ChangeLayer(0)
    return sprite, nil
}
func LoadAnimatedSpriteFromTextures(textures []*sdl.Texture, delays []int32) (*Sprite, error) {
    if (len(textures) != len(delays)) {return &Sprite{}, errors.New("argument lengths must be equal \"textures " + string(len(textures)) + "  delays " + string(len(delays)))}

    sprite, _ := NewSprite()

    for _, frame := range textures {
        _, _, w, h, err := frame.Query();    if err != nil {return &Sprite{}, errors.Wrap(err, "could not determine sprite dimensions")}
        sprite.dimensionsHalved = append(sprite.dimensionsHalved, [2]int32{w >> 1, h >> 1})
    }

    sprite.Frames = textures
    sprite.Delays = delays
    return sprite, nil
}
func LoadAnimatedSpriteFromFiles(fileNames []string, delays []int32) (*Sprite, error) {
    var textures []*sdl.Texture
    for _, fileName := range fileNames {
        texture, err := img.LoadTexture(renderer, "./assets/sprites/" + fileName);    if err != nil {return &Sprite{}, errors.Wrap(err, "could not load sprite file \"./assets/sprites/" + fileName)}
        textures = append(textures, texture)
    }
    return LoadAnimatedSpriteFromTextures(textures, delays)
}
func LoadSpriteFromFile(filename string) (*Sprite, error) {
    return LoadAnimatedSpriteFromFiles([]string{filename}, []int32{0})
}
func LoadAnimatedSpriteFromFile(filename string, rects []sdl.Rect, delays []int32) (*Sprite, error) {
    surface, err := img.Load(filename);    if err != nil {return &Sprite{}, errors.Wrap(err, "could not load sprite image")}
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
        tmpSurface, err := sdl.CreateRGBSurface(0, rect.W, rect.H, 32, rmask, gmask, bmask, amask);    if err != nil {return &Sprite{}, errors.Wrap(err, "could not create tmpSurface for transfer")}
        rect.X += xOffset
        xOffset += rect.W
        surface.Blit(&rect, tmpSurface, nil)
        texture, err := renderer.CreateTextureFromSurface(tmpSurface);    if err != nil {return &Sprite{}, errors.Wrap(err, "could not transfer surface to texture")}
        textures = append(textures, texture)
    }
    return LoadAnimatedSpriteFromTextures(textures, delays)
}


func (s *Sprite) ChangeLayer(layer int32) {
    s.layer = layer
    for i := len(sprites) - 1; i >= 0 ; i-- {
        sp := sprites[i]
        if sp.layer <= s.layer {
            for i, sp := range sprites {
                if s == sp {
                    var newSprites []*Sprite
                    newSprites = append(newSprites, sprites[:i]...)
                    newSprites = append(newSprites, sprites[i + 1:]...)
                    sprites = newSprites
                }
            }
            var newSprites []*Sprite
            newSprites = append(newSprites, sprites[:i]...)
            newSprites = append(newSprites, s)
            newSprites = append(newSprites, sprites[i:]...)
            sprites = newSprites
            break
        }
    }
}

func (s *Sprite) Blit() error {
    //Calculate topleft point from dimensions and center
    destRect := sdl.Rect{s.XCenter - s.dimensionsHalved[s.FrameIndex][0], s.YCenter - s.dimensionsHalved[s.FrameIndex][1], s.dimensionsHalved[s.FrameIndex][0] << 1, s.dimensionsHalved[s.FrameIndex][1] << 1}
    renderer.Copy(s.Frames[s.FrameIndex], nil, &destRect)
    return nil
}
