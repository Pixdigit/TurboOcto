package turboOcto

import (
	"testing"

	tools "gitlab.com/Pixdigit/goTestTools"
)

func TestDeserialization(t *testing.T) {

	exampleConf := []byte(`[turboOcto]
		UpdateOnRefresh    = true
		AllowFrameSkipping = false
		DefaultSpriteTimerMode    = 0
		ResourcePath       = ./

		[internal]
		Fullscreen   = false
		WindowWidth  = 50
		WindowHeight = 50
		VHeight      = 1366
		VWidth       = 768`)

	err := LoadDefaultConf()
	if err != nil {
		tools.WrapErr(err, "could not load default configuration", t)
	}

	err = LoadConf(exampleConf)
	if err != nil {
		tools.WrapErr(err, "failed to load configuration", t)
	}

	tools.Test(!Cfg.AllowFrameSkipping, "default configuration did not change", t)
	tools.Test(Cfg.UpdateOnRefresh, "default configuration did not change", t)
	tools.Test(Cfg.DefaultSpriteTimerMode == 0, "default configuration did not change", t)

	tools.Test(!isFullscreen, "internal configuration did not change", t)
	t.Log(windowSize)
	tools.Test(windowSize.Width == 50, "internal configuration did not change", t)
	tools.Test(windowSize.Height == 50, "internal configuration did not change", t)
	t.Log(vSize)
	tools.Test(vSize.Width == 768, "internal configuration did not change", t)
	tools.Test(vSize.Height == 1366, "internal configuration did not change", t)

	//TODO: test saving

}
