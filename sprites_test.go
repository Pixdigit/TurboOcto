package turboOcto

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestSpriteLayers(t *testing.T) {
	LoadDefaultConf()
	sp, err := NewSprite()
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	sp2, err := NewSprite()
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	//create unreferenced sprite
	_, err = NewSprite()
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	err = sp2.ChangeLayer(2)
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	err = sp2.ChangeLayer(3)
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	//create unreferenced sprite
	_, err = NewSprite()
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	tools.Test(sprites[0] == sp, "sprite did not remain in position", t)
	tools.Test(len(sprites) == 4, "unexpected sprite count", t)
	tools.Test(sprites[len(sprites)-1] == sp2, "sprite on layer 3 did not move to correct position", t)
}

func TestSpriteRendering(t *testing.T) {
	testTextures := make([]*sdl.Texture, 0)
	for i := 0; i < 5; i++ {
		tex, err := screenRenderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, 10, 10)
		if err != nil {
			tools.WrapErr(err, "error while creating test textures", t)
		}
		testTextures = append(testTextures, tex)
	}
	dFrames := int32(0)
	err := SetConf("spriteTimerMode", USE_FRAME_COUNT)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}

	err = SetConf("allowFrameSkipping", false)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}
	sp, err := LoadAnimatedSpriteFromTextures(testTextures, []int32{0, 0, 0, 0, 0})
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	frameCount = 0
	sp.lastFrameCount = 0

	for _, v := range sp.Delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := int32(0); i < dFrames; i++ {
		tools.Test(sp.FrameIndex == i, "FrameIndex mismatch without FrameSkip", t)
		err = sp.IncrementTime()
		if err != nil {
			tools.WrapErr(err, "could not blit sprite", t)
		}
		Present()
	}

	err = SetConf("allowFrameSkipping", true)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}
	sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{0, 1, 0, 1, 1})
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	expectedFrameIndexes := []int32{1, 3, 4}

	for _, v := range sp.Delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := int32(0); i < dFrames; i++ {
		err = sp.IncrementTime()
		if err != nil {
			tools.WrapErr(err, "could not blit sprite", t)
		}
		tools.Test(sp.FrameIndex == expectedFrameIndexes[int(i)%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and singe blit frames", t)
		Present()
	}

	err = SetConf("allowFrameSkipping", true)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}
	sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{0, 3, 0, 5, 1})
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	expectedFrameIndexes = []int32{1, 1, 1, 3, 3, 3, 3, 3, 4}

	for _, v := range sp.Delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := int32(0); i < dFrames; i++ {
		err = sp.IncrementTime()
		if err != nil {
			tools.WrapErr(err, "could not blit sprite", t)
		}
		tools.Test(sp.FrameIndex == expectedFrameIndexes[int(i)%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and various delays", t)
		Present()
	}

	err = SetConf("allowFrameSkipping", true)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}
	sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{1, 1, -2, 1, 1})
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	expectedFrameIndexes = []int32{0, 1}

	for _, v := range sp.Delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := int32(0); i < dFrames; i++ {
		err = sp.IncrementTime()
		if err != nil {
			tools.WrapErr(err, "could not blit sprite", t)
		}
		tools.Test(sp.FrameIndex == expectedFrameIndexes[int(i)%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and unexpected delays", t)
		Present()
	}

	err = SetConf("allowFrameSkipping", true)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}
	sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{1, 2, 1, 2, 1})
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	expectedFrameIndexes = []int32{0, 1, 3, 4, 1, 2, 3}

	for _, v := range sp.Delays {
		if v == 0 {
			dFrames++
		} else {
			dFrames = dFrames + v
		}
	}
	for i := int32(0); i < dFrames; i++ {
		err = sp.IncrementTime()
		if err != nil {
			tools.WrapErr(err, "could not blit sprite", t)
		}

		tools.Test(sp.FrameIndex == expectedFrameIndexes[int(i)%len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and multiple present", t)
		Present()
		Present()
	}
}

func TestSpriteControl(t *testing.T) {
	testTextures := make([]*sdl.Texture, 0)
	for i := 0; i < 5; i++ {
		tex, err := screenRenderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, 10, 10)
		if err != nil {
			tools.WrapErr(err, "error while creating test textures", t)
		}
		testTextures = append(testTextures, tex)
	}
	err := SetConf("spriteTimerMode", USE_FRAME_COUNT)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}
	err = SetConf("allowFrameSkipping", false)
	if err != nil {
		tools.WrapErr(err, "error while setting configuration", t)
	}

	sp, err := LoadAnimatedSpriteFromTextures(testTextures, []int32{1, 1, 1, 1, 1})
	if err != nil {
		tools.WrapErr(err, "could not create test sprite", t)
	}
	sp.lastFrameCount = 0

	Present()
	sp.IncrementTime()
	sp.Pause()
	Present()
	Present()
	Present()
	sp.IncrementTime()
	tools.Test(sp.FrameIndex == 1, "pausing sprite did not halt frameCount", t)

	sp.Stop()
	sp.Start()
	sp.IncrementTime()
	tools.Test(sp.FrameIndex == 0, "stopping sprite did not reset frameCount", t)
	//FIXME: why do I need to present twice
	Present()
	//Present()
	sp.IncrementTime()
	tools.Test(sp.FrameIndex == 1, "sprite did not start frameCount after stop", t)
	Quit()

}
