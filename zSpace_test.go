package turboOcto

import (
	"testing"

	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestZSpace(t *testing.T) {
	frame1, err := NewEmptyFrame()
	if err != nil {
		tools.WrapErr(err, "could not create empty frame for new Sprite", t)
	}
	frame2, err := NewEmptyFrame()
	if err != nil {
		tools.WrapErr(err, "could not create empty frame for new Sprite", t)
	}
	err = AddElement(frame1, 0)
	if err != nil {
		tools.WrapErr(err, "could not add Frame to ZSpace", t)
	}
	err = AddElement(frame1, 0)
	if err == nil {
		t.Error("ZSpace allowed adding the same element twice")
	}
	err = AddElement(frame2, 0)
	if err != nil {
		tools.WrapErr(err, "could not add second Frame to ZSpace", t)
	}
}
