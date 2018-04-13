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
    err := SetConf("spriteTimerMode", USE_FRAME_COUNT);    if err != nil {wrapErr(err, "error while setting configuration", t)}
    err = SetConf("allowFrameSkipping", false);    if err != nil {wrapErr(err, "error while setting configuration", t)}

    testTextures := make([]*sdl.Texture, 0)
    for i := 0; i < 10; i++ {
        tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_STREAMING, 10, 10);    if err != nil {wrapErr(err, "error while creating test textures", t)}
        testTextures = append(testTextures, tex)
    }

    sp, err := LoadAnimatedSpriteFromTextures(testTextures, []int32{1, 2, 1, 5, 0, 1, 1, 1, 1, 1});    if err != nil {wrapErr(err, "could not create test sprite", t)}

    for i := 0; i < 14; i++ {
        sp.Blit()
        print(sp.FrameIndex)
        print("\n")
        //TODO: adjust to delay array
        //test(testTextures[i] == sp.frames[i], "sprite frame is not as expected (either pointer dref or something with FrameIndex)", t)
        //test(sp.FrameIndex == int32(i), "FrameIndex has not been incremented", t)
        Present()
    }


}
