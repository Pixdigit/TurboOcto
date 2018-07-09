package turboOcto

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type sprite struct {
	frames             []*sdl.Texture
	dimensions         []Size
	rect               *Rect
	Delays             []int32
	animationStatus    Runlevel
	timerMode          int32
	timer              int32
	lastBlit           int32
	lastFrameCount     int32
	lastTimer          int32
	AllowFrameSkipping bool
	FrameIndex         int32
	layer              int32
}

const (
	USE_FRAME_COUNT = iota
	USE_TIME_PASSED
)

var sprites []*sprite

func NewSprite() (*sprite, error) {
	newSprite := &sprite{}
	//ensure sprite is the topmost of level 0
	sprites = append([]*sprite{newSprite}, sprites...)
	err := newSprite.ChangeLayer(0);	if err != nil {return &sprite{}, errors.Wrap(err, "could not read default configuration for new sprite")}

	timerMode, err := GetConf("spriteTimerMode");	if err != nil {return &sprite{}, errors.Wrap(err, "could not read configuration for new sprite")}
	newSprite.timerMode = int32(timerMode.(int))
	AllowFrameSkipping, err := GetConf("allowFrameSkipping");	if err != nil {return &sprite{}, errors.Wrap(err, "could not read configuration for new sprite")}
	newSprite.AllowFrameSkipping = AllowFrameSkipping.(bool)
	newSprite.lastFrameCount = frameCount
	newSprite.animationStatus = RUNNING
	newSprite.rect, err = NewRect(Point{0, 0}, Size{0, 0}, AnchorPoint{CENTERX, CENTERY});	if err != nil {return &sprite{}, errors.Wrap(err, "could not create bounding rect for sprite")}

	return newSprite, nil
}

func (s *sprite) ChangeLayer(layer int32) error {
	s.layer = layer
	for i := len(sprites) - 1; i >= 0; i-- {
		sp := sprites[i]
		if sp.layer <= s.layer {
			for i, sp := range sprites {
				if s == sp {
					var newSprites []*sprite
					newSprites = append(newSprites, sprites[:i]...)
					newSprites = append(newSprites, sprites[i+1:]...)
					sprites = newSprites
				}
			}
			var newSprites []*sprite
			newSprites = append(newSprites, sprites[:i]...)
			newSprites = append(newSprites, s)
			newSprites = append(newSprites, sprites[i:]...)
			sprites = newSprites
			break
		}
	}
	return nil
}

func (s *sprite) IncrementTime() error {
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
	return nil
}

func (s *sprite) BlitToScreen() error {
	err := s.rect.SetSize(s.dimensions[s.FrameIndex]);	if err != nil {return errors.Wrap(err, "could not change sprite size to dimension for this frame")}
	dstRect, err := s.rect.SDLRect();	if err != nil {return errors.Wrap(err, "could not convert sprite boundaries to SDL rect")}
	err = screenRenderer.Copy(s.frames[s.FrameIndex], nil, dstRect);	if err != nil {return errors.Wrap(err, "could not copy sprite frame to screenRenderer")}
	return nil
}

func (s *sprite) Blit(dstTexture *sdl.Texture) error {
	err := s.rect.SetSize(s.dimensions[s.FrameIndex]);	if err != nil {return errors.Wrap(err, "could not change sprite size to dimension for this frame")}
	dstRect, err := s.rect.SDLRect();	if err != nil {return errors.Wrap(err, "could not convert sprite boundaries to SDL rect")}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_TARGETTEXTURE);	if err != nil {return errors.Wrap(err, "could not create renderer to render to texture")}
	renderer.SetRenderTarget(dstTexture)
	err = renderer.Copy(s.frames[s.FrameIndex], nil, dstRect);	if err != nil {return errors.Wrap(err, "could not copy sprite frame to texture")}
	return nil
}

func (s *sprite) Start() error {
	s.animationStatus = RUNNING
	s.lastFrameCount = frameCount
	s.lastBlit = int32(sdl.GetTicks())
	s.lastTimer = s.timer - 1
	return nil
}
func (s *sprite) Stop() error {
	s.FrameIndex = 0
	s.timer = 0
	s.animationStatus = STOPPED
	return nil
}
func (s *sprite) Pause() error {
	s.animationStatus = PAUSED
	return nil
}

func BlitAllToScreen() error {
	for _, sp := range sprites {
		err := sp.BlitToScreen();	if err != nil {return errors.Wrap(err, "could not blit all sprites")}
	}
	return nil
}

func UpdateAll() error {
	for _, sp := range sprites {
		err := sp.IncrementTime();	if err != nil {return errors.Wrap(err, "could not increment time for all sprites")}
	}
	return nil
}
