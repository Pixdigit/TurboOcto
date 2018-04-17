package TurboOcto

import (
    "testing"
    "github.com/veandco/go-sdl2/sdl"
)

func TestSpriteLayers(t *testing.T) {
    LoadDefaultConf()
    sp, err := NewSprite();    if err != nil {wrapErr(err, "could not create test sprite", t)}
    sp2, err := NewSprite();    if err != nil {wrapErr(err, "could not create test sprite", t)}
    //create unreferenced sprite
    _, err = NewSprite();     if err != nil {wrapErr(err, "could not create test sprite", t)}
    err = sp2.ChangeLayer(2);    if err != nil {wrapErr(err, "could not create test sprite", t)}
    err = sp2.ChangeLayer(3);    if err != nil {wrapErr(err, "could not create test sprite", t)}
    //create unreferenced sprite
    _, err = NewSprite();     if err != nil {wrapErr(err, "could not create test sprite", t)}
    test(sprites[0] == sp, "sprite did not remain in position", t)
    test(len(sprites) == 4, "unexpected sprite count", t)
    test(sprites[len(sprites) - 1] == sp2, "sprite on layer 3 did not move to correct position", t)
}

func TestSpriteRendering(t *testing.T) {
    testTextures := make([]*sdl.Texture, 0)
    for i := 0; i < 5; i++ {
        tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, 10, 10);    if err != nil {wrapErr(err, "error while creating test textures", t)}
        testTextures = append(testTextures, tex)
    }
    dFrames := int32(0)
    err := SetConf("spriteTimerMode", USE_FRAME_COUNT);    if err != nil {wrapErr(err, "error while setting configuration", t)}


    err = SetConf("allowFrameSkipping", false);    if err != nil {wrapErr(err, "error while setting configuration", t)}
    sp, err := LoadAnimatedSpriteFromTextures(testTextures, []int32{0, 0, 0, 0, 0});    if err != nil {wrapErr(err, "could not create test sprite", t)}
    frameCount = 0
    sp.lastFrameCount = 0

    for _, v := range sp.Delays {if v == 0 {dFrames++} else {dFrames = dFrames + v}}
    for i := int32(0); i < dFrames; i++ {
        err = sp.Blit();    if err != nil {wrapErr(err, "could not blit sprite", t)}
        test(sp.FrameIndex == i, "FrameIndex mismatch without FrameSkip", t);
        Present()
    }


    err = SetConf("allowFrameSkipping", true);    if err != nil {wrapErr(err, "error while setting configuration", t)}
    sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{0, 1, 0, 1, 1});    if err != nil {wrapErr(err, "could not create test sprite", t)}
    expectedFrameIndexes := []int32{1, 3, 4}

    for _, v := range sp.Delays {if v == 0 {dFrames++} else {dFrames = dFrames + v}}
    for i := int32(0); i < dFrames; i++ {
        err = sp.Blit();    if err != nil {wrapErr(err, "could not blit sprite", t)}
        test(sp.FrameIndex == expectedFrameIndexes[int(i) % len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and singe blit frames", t);
        Present()
    }


    err = SetConf("allowFrameSkipping", true);    if err != nil {wrapErr(err, "error while setting configuration", t)}
    sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{0, 3, 0, 5, 1});    if err != nil {wrapErr(err, "could not create test sprite", t)}
    expectedFrameIndexes = []int32{1, 1, 1, 3, 3, 3, 3, 3, 4}

    for _, v := range sp.Delays {if v == 0 {dFrames++} else {dFrames = dFrames + v}}
    for i := int32(0); i < dFrames; i++ {
        err = sp.Blit();    if err != nil {wrapErr(err, "could not blit sprite", t)}
        test(sp.FrameIndex == expectedFrameIndexes[int(i) % len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and various delays", t);
        Present()
    }


    err = SetConf("allowFrameSkipping", true);    if err != nil {wrapErr(err, "error while setting configuration", t)}
    sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{1, 1, -2, 1, 1});    if err != nil {wrapErr(err, "could not create test sprite", t)}
    expectedFrameIndexes = []int32{0, 1}

    for _, v := range sp.Delays {if v == 0 {dFrames++} else {dFrames = dFrames + v}}
    for i := int32(0); i < dFrames; i++ {
        err = sp.Blit();    if err != nil {wrapErr(err, "could not blit sprite", t)}
        test(sp.FrameIndex == expectedFrameIndexes[int(i) % len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and unexpected delays", t);
        Present()
    }


    err = SetConf("allowFrameSkipping", true);    if err != nil {wrapErr(err, "error while setting configuration", t)}
    sp, err = LoadAnimatedSpriteFromTextures(testTextures, []int32{1, 2, 1, 2, 1});    if err != nil {wrapErr(err, "could not create test sprite", t)}
    expectedFrameIndexes = []int32{0, 1, 3, 4, 1, 2, 3}

    for _, v := range sp.Delays {if v == 0 {dFrames++} else {dFrames = dFrames + v}}
    for i := int32(0); i < dFrames; i++ {
        err = sp.Blit();    if err != nil {wrapErr(err, "could not blit sprite", t)}

        test(sp.FrameIndex == expectedFrameIndexes[int(i) % len(expectedFrameIndexes)], "FrameIndex mismatch with FrameSkip and multiple present", t);
        Present()
        Present()
    }


}
