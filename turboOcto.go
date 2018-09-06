package turboOcto

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func init() {
	sdl.Init(sdl.INIT_EVERYTHING)
	img.Init(0x0000000F) // initialize all formats
	err := initializeGraphics()
	if err != nil {
		fmt.Println(errors.Wrap(err, "could not initialize graphics"))
		Quit()
	}
	err = initializeConfiguration()
	if err != nil {
		fmt.Println(errors.Wrap(err, "could not initialize environment"))
		Quit()
	}
	err = initializeTextProcessing()
	if err != nil {
		fmt.Println(errors.Wrap(err, "could not initialize text processing"))
		Quit()
	}

}

func Update() error {
	err := updateAllSprites();	if err != nil {return errors.Wrap(err, "could not update sprites")}
	err = Present();	if err != nil {return errors.Wrap(err, "could not update display")}
	err = updateEvents();	if err != nil {return errors.Wrap(err, "could not update Events")}
	return nil
}

func Quit() error {
	screenRenderer.Destroy()
	window.Destroy()
	img.Quit()
	if ttf.WasInit() {
		ttf.Quit()
	}
	//If anything was init
	sdl.Quit()
	if Cfg.SaveOnQuit {
		err := SaveConf();	if err != nil {return errors.Wrap(err, "Could not save config on quit")}
	}
	return nil
}
