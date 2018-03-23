package TurboOcto

import (
    "testing"
    "github.com/pkg/errors"
)

func TestSpriteLayers(t *testing.T) {
    sp, err := NewSprite();    if err != nil {t.Error(errors.New("Could not create test sprite"))}
    sp2, err := NewSprite();    if err != nil {t.Error(errors.New("Could not create test sprite"))}
    NewSprite()
    sp2.ChangeLayer(2)
    sp2.ChangeLayer(3)
    NewSprite()
    test(sprites[0] == sp, "sprite not kept track of", t)
    test(sprites[0] == sp, "sprite did not remain in position", t)
    test(len(sprites) == 4, "unexpected sprite count", t)
    test(sprites[len(sprites) - 1] == sp2, "sprite on layer 3 did not move to correct position", t)
}
