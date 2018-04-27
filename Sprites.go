package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "github.com/pkg/errors"
)

const STOPPED = 0;
const RUNNING = 1;
const PAUSED = 2;

type Sprite struct {
    frames  []*sdl.Texture
    dimensions [][2]int32
    XCenter, YCenter int32
    Delays  []int32
    animationStatus int32
    timerMode int32
    timer int32
    lastBlit int32
    lastFrameCount int32
    lastTimer int32
    AllowFrameSkipping bool
    FrameIndex int32
    layer int32
}

var USE_FRAME_COUNT int32 = 1
var USE_TIME_PASSED int32 = 2

var sprites []*Sprite

func NewSprite() (*Sprite, error) {
    sprite := &Sprite{}
    //ensure sprite is the topmost of level 0
    sprites = append([]*Sprite{sprite}, sprites...)
    err := sprite.ChangeLayer(0);    if err != nil {return &Sprite{}, errors.Wrap(err, "could not read default configuration for new sprite")}

    timerMode, err := GetConf("spriteTimerMode"); if err != nil {return &Sprite{}, errors.Wrap(err, "could not read configuration for new sprite")}
    sprite.timerMode = int32(timerMode.(int))
    AllowFrameSkipping, err := GetConf("allowFrameSkipping"); if err != nil {return &Sprite{}, errors.Wrap(err, "could not read configuration for new sprite")}
    sprite.AllowFrameSkipping = AllowFrameSkipping.(bool)
    sprite.lastFrameCount = frameCount
    sprite.animationStatus = RUNNING

    return sprite, nil
}
func LoadAnimatedSpriteFromTextures(textures []*sdl.Texture, delays []int32) (*Sprite, error) {
    if (len(textures) != len(delays)) {return &Sprite{}, errors.New("argument lengths must be equal \"textures " + string(len(textures)) + "  delays " + string(len(delays)))}

    var dimensions [][2]int32
    sprite, _ := NewSprite()
    for _, frame := range textures {
        _, _, w, h, err := frame.Query();    if err != nil {return &Sprite{}, errors.Wrap(err, "could not determine texture size")}
        dimensions = append(dimensions, [2]int32{w, h})
    }

    sprite.frames = textures
    sprite.Delays = delays
    sprite.dimensions = dimensions
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


func (s *Sprite) ChangeLayer(layer int32) error {
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
    return nil
}

func (s *Sprite) Blit() error {
    currentTime := int32(sdl.GetTicks())
    if s.animationStatus == RUNNING {
        if s.timerMode == USE_FRAME_COUNT {
            s.timer += frameCount - s.lastFrameCount
        } else if s.timerMode == USE_TIME_PASSED {
            s.timer += currentTime - s.lastBlit
        }
    }
    s.lastBlit = int32(currentTime)
    s.lastFrameCount = frameCount

    if s.timer >= s.Delays[s.FrameIndex] {
        if s.AllowFrameSkipping {
            for s.timer >= s.Delays[s.FrameIndex] {
                s.timer = s.timer - s.Delays[s.FrameIndex]
                s.FrameIndex = (s.FrameIndex + 1) % int32(len(s.frames))
            }
        } else {
            //If we have no frame skipping ensure at least one blit
            s.timer = s.timer - s.Delays[s.FrameIndex]
            if s.timer > s.lastTimer || (s.FrameIndex == 0 && s.timer == 0) {
                s.FrameIndex = (s.FrameIndex + 1) % int32(len(s.frames))
            }
        }
    }
    s.lastTimer = s.timer

    dstRect := sdl.Rect{s.XCenter - (s.dimensions[s.FrameIndex][0] >> 2), s.YCenter - (s.dimensions[s.FrameIndex][1] >> 2), s.dimensions[s.FrameIndex][0], s.dimensions[s.FrameIndex][1]}
    err := renderer.Copy(s.frames[s.FrameIndex], nil, &dstRect);    if err != nil {return errors.Wrap(err, "could not copy sprite frame to renderer")}

    return nil
}

func (s *Sprite) Start() error {
    s.animationStatus = RUNNING
    s.lastFrameCount = frameCount
    s.lastBlit = int32(sdl.GetTicks())
    s.lastTimer = s.timer - 1
    return nil
}
func (s *Sprite) Stop() error {
    s.FrameIndex = 0
    s.timer = 0
    s.animationStatus = STOPPED
    return nil
}
func (s *Sprite) Pause() error {
    s.animationStatus = PAUSED
    return nil
}


func BlitAll() error {
    for _, sp := range sprites {
        err := sp.Blit();    if err != nil {return errors.Wrap(err, "could not blit all sprites")}
    }
    return nil
}
