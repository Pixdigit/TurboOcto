package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "github.com/pkg/errors"
)

type Sprite struct {
    frames  []*sdl.Texture
    delays  []int32
    XCenter, YCenter int32
    Width, Height int32
    animationDelay int32
}

func NewSprite() (Sprite, error) {
    return Sprite{}, nil
}

func LoadAnimatedSprite(fileNames []string, delays []int32, rects []sdl.Rect) (Sprite, error) {
    if (len(fileNames) != len(delays)) || (len(delays) != len(rects)) {return Sprite{}, errors.New("argument lengths must be equal \"files " + string(len(fileNames)) + "  delays " + string(len(delays)) + "  rects " + string(len(rects)))}

    sprite := Sprite{}

    for i, fileName := range fileNames {
        ok, err := pathExists("./assets/sprites/" + fileName);    if err != nil {return Sprite{}, errors.Wrap(err, "could not check if path exists \"./assets/sprites/" + fileName + "\"")}
        if !ok {return Sprite{}, errors.New("sprite file does not exist \"./assets/sprites/" + fileName + "\"")}
        surface, err := img.Load("./assets/sprites/" + fileName);    if err != nil {return Sprite{}, errors.Wrap(err, "could not load sprite file \"./assets/sprites/" + fileName)}

        tmpSurface, err := sdl.CreateRGBSurface(0, rects[i].W, rects[i].H, 32, rmask, gmask, bmask, amask);    if err != nil {return Sprite{}, errors.Wrap(err, "could not create tmpSurface for transfer")}
        //Centered to tmpSurface
        srcRect := sdl.Rect{(rects[i].W - surface.W) / 2, (rects[i].H - surface.H) / 2, rects[i].W, rects[i].H}
        surface.Blit(&srcRect, tmpSurface, nil)
        texture, err := renderer.CreateTextureFromSurface(tmpSurface);    if err != nil {return Sprite{}, errors.Wrap(err, "could not transfer surface to texture")}

        sprite.frames = append(sprite.frames, texture)
        sprite.delays = append(sprite.delays, delays[i])
    }

    return sprite, nil
}

/*
spriteSheet, err := img.Load("./assets/sprites/" + fileName);    if err != nil {return nil, errors.Wrap(err, "could not load sprite sheet \"./assets/sprites/" + fileName + "\"")}
spriteSheet.Lock()
animationLength := spriteSheet.W / spriteSheet.H
*/
