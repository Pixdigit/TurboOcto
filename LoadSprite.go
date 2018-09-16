package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadAnimatedSpriteFromFrames(frames []*Frame, delays []int) (*Sprite, error) {
	if len(frames) != len(delays) {
		return &Sprite{}, errors.New("argument lengths must be equal \"frames " + string(len(frames)) + "  delays " + string(len(delays)))
	}

	newSprite, err := NewSprite();	if err != nil {return nil, errors.Wrap(err, "could not create empty sprite to load data into")}

	newSprite.frames = frames
	newSprite.delays = delays
	newSprite.FrameIndex = 0
	newSprite.timer.Duration = float64(delays[newSprite.FrameIndex])

	//ensure Sprite has some delay at any frame
	ok := newSprite.validateDelays()
	if !ok {
		return &Sprite{}, errors.New("new Sprite has invalid delay")
	}
	return newSprite, nil
}

func LoadAnimatedSpriteFromTexture(frame *Frame) (*Sprite, error) {
	return LoadAnimatedSpriteFromFrames([]*Frame{frame}, []int{0})
}

func LoadAnimatedSpriteFromFiles(fileNames []string, delays []int) (*Sprite, error) {
	var frames []*Frame
	for _, fileName := range fileNames {
		texture, err := img.LoadTexture(screenRenderer, Cfg.ResourcePath+fileName);	if err != nil {return &Sprite{}, errors.Wrap(err, "could not load Sprite file \""+Cfg.ResourcePath+fileName+"\"")}
		frame, err := NewFrame(texture);	if err != nil {return nil, errors.Wrap(err, "could not load frames for new Sprite")}
		frames = append(frames, frame)
	}
	return LoadAnimatedSpriteFromFrames(frames, delays)
}

func LoadSpriteFromFile(filename string) (*Sprite, error) {
	return LoadAnimatedSpriteFromFiles([]string{filename}, []int{0})
}

func LoadAnimatedSpriteFromFile(filename string, rects []sdl.Rect, delays []int) (*Sprite, error) {
	surface, err := img.Load(filename);	if err != nil {return &Sprite{}, errors.Wrap(err, "could not load Sprite image")}
	if len(rects) == 0 {
		//D == Amount of
		DSprites := surface.W / surface.H
		for i := int32(0); i < DSprites; i++ {
			rects = append(rects, sdl.Rect{i * surface.H, 0, surface.H, surface.H})
		}
	}
	var frames []*Frame
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
		frame, err := NewFrame(texture);	if err != nil {return nil, errors.Wrap(err, "could not create new frame for new sprite")}
		frames = append(frames, frame)
	}
	return LoadAnimatedSpriteFromFrames(frames, delays)
}
