package turboOcto

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestSpriteRendering(t *testing.T) {
	testFrames := make([]*Frame, 0)
	for i := 0; i < 5; i++ {
		tex, err := screenRenderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, 10, 10);	if err != nil {tools.WrapErr(err, "error while creating test textures", t)}
		frame, err := NewFrame(tex);	if err != nil {tools.WrapErr(err, "could not create testing frame", t)}
		testFrames = append(testFrames, frame)
	}
	dFrames := 0
	Cfg.DefaultSpriteTimerMode = USE_FRAME_COUNT
	Cfg.AllowFrameSkipping = false
	sp, err := LoadAnimatedSpriteFromFrames(testFrames, []int{0, 0, 0, 0, 0});	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	AddElement(sp, 0)
	frameCount = 0
	sp.Start()

	for _, v := range sp.delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := 0; i < dFrames; i++ {
		tools.Test(sp.FrameIndex == i, "FrameIndex mismatch without FrameSkip", t)
		Render()
	}

	Cfg.AllowFrameSkipping = true
	sp, err = LoadAnimatedSpriteFromFrames(testFrames, []int{0, 1, 0, 1, 1});	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	AddElement(sp, 0)
	sp.Start()
	expectedFrameIndexes := []int{1, 3, 4}

	for _, v := range sp.delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := 0; i < dFrames; i++ {
		Render()
		tools.Test(sp.FrameIndex == expectedFrameIndexes[i%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and singe blit frames", t)
	}

	Cfg.AllowFrameSkipping = true
	sp, err = LoadAnimatedSpriteFromFrames(testFrames, []int{0, 3, 0, 5, 1});	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	AddElement(sp, 0)
	sp.Start()
	expectedFrameIndexes = []int{1, 1, 1, 3, 3, 3, 3, 3, 4}

	for _, v := range sp.delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := 0; i < dFrames; i++ {
		Render()
		tools.Test(sp.FrameIndex == expectedFrameIndexes[i%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and various delays", t)
	}

	Cfg.AllowFrameSkipping = true
	sp, err = LoadAnimatedSpriteFromFrames(testFrames, []int{1, 1, -2, 1, 1});	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	AddElement(sp, 0)
	sp.Start()
	expectedFrameIndexes = []int{0, 1}

	for _, v := range sp.delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := 0; i < dFrames; i++ {
		Render()
		tools.Test(sp.FrameIndex == expectedFrameIndexes[i%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and unexpected delays", t)
	}

	Cfg.AllowFrameSkipping = true
	sp, err = LoadAnimatedSpriteFromFrames(testFrames, []int{1, 2, 1, 2, 1});	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	AddElement(sp, 0)
	sp.Start()
	//increment in 2 frame steps
	expectedFrameIndexes = []int{1, 2, 3, 0, 1, 3, 4}

	for _, v := range sp.delays {
		if v == 0 {
		} else {
			dFrames++
			dFrames = dFrames + v
		}
	}
	for i := 0; i < dFrames; i++ {
		Render()
		Render()
		tools.Test(sp.FrameIndex == expectedFrameIndexes[i%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and multiple present", t)
	}
}

func TestSpriteControl(t *testing.T) {
	testFrames := make([]*Frame, 0)
	for i := 0; i < 5; i++ {
		tex, err := screenRenderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, 10, 10);	if err != nil {tools.WrapErr(err, "error while creating test textures", t)}
		frame, err := NewFrame(tex);	if err != nil {tools.WrapErr(err, "could not create testing frame", t)}
		testFrames = append(testFrames, frame)
	}
	Cfg.DefaultSpriteTimerMode = USE_FRAME_COUNT
	Cfg.AllowFrameSkipping = true

	sp, err := LoadAnimatedSpriteFromFrames(testFrames, []int{1, 1, 1, 1, 1});	if err != nil {tools.WrapErr(err, "could not create test sprite", t)}
	AddElement(sp, 0)
	sp.Start()

	tools.Test(sp.FrameIndex == 0, "framecount changed before start of animation", t)

	Render()
	Render()
	Render()
	sp.Pause()
	Render()
	Render()
	Render()
	tools.Test(sp.FrameIndex == 2, "pausing sprite did not halt frameCount", t)

	sp.Stop()
	sp.Start()
	tools.Test(sp.FrameIndex == 0, "stopping sprite did not reset frameCount", t)
	Render()
	Render()
	Render()
	Render()
	tools.Test(sp.FrameIndex == 3, "sprite did not start frameCount after stop", t)
}
