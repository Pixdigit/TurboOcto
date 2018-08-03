package turboOcto

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

type internals struct {
	Fullscreen   bool
	WindowWidth  int32
	WindowHeight int32
	VHeight      int32
	VWidth       int32
}
type configuration struct {
	UpdateOnRefresh    bool
	AllowFrameSkipping bool
	SpriteTimerMode    timerMode
	ResourcePath       string
	internal           internals
}

const confSectionName = "turboOcto"
const internalConfSectionName = "internal"

var Cfg configuration

func initializeConfiguration() error {
	err := LoadDefaultConf();	if err != nil {return errors.Wrap(err, "could not initialize configuration with default values")}
	return nil
}

func LoadDefaultConf() error {
	Cfg = configuration{
		UpdateOnRefresh:    true,
		AllowFrameSkipping: true,
		SpriteTimerMode:    USE_FRAME_COUNT,
		ResourcePath:       "./",
		internal: internals{
			Fullscreen: true,
		},
	}
	err := setDefaultInternals();	if err != nil {return errors.Wrap(err, "could not set internal configuration to default")}
	err = updateFromInternals();	if err != nil {return errors.Wrap(err, "could not update based on conf")}
	return nil
}

func setDefaultInternals() error {
	isFullscreen = true
	windowWidth = screenWidth / 2
	windowHeight = screenHeight / 2
	vWidth = windowWidth
	vHeight = windowHeight
	return nil
}

func LoadConf(dataSrc interface{}) error {
	//dataSrc can be file path
	err := LoadDefaultConf();	if err != nil {return errors.Wrap(err, "set defaults before loading configuration")}
	cfgIniFile, err := ini.Load(dataSrc);	if err != nil {return errors.Wrap(err, "could not load configuration")}
	err = cfgIniFile.Section(confSectionName).MapTo(&Cfg);	if err != nil {return errors.Wrap(err, "could not load configuration")}
	err = cfgIniFile.Section(internalConfSectionName).MapTo(&Cfg.internal);	if err != nil {return errors.Wrap(err, "could not load configuration")}
	err = updateFromInternals();	if err != nil {return errors.Wrap(err, "could not update based on conf")}
	return nil
}

func SaveConf(filePath string) error {
	ok, err := pathExists(filePath);	if err != nil {return errors.Wrap(err, "could not check if path to configuration exists")}
	if !ok {
		_, err = os.Create(filePath);	if err != nil {return errors.Wrap(err, "could not create file to save conf in")}
	}

	Cfg.internal = internals{
		isFullscreen,
		windowWidth,
		windowHeight,
		vWidth,
		vHeight,
	}

	cfgIniFile, err := ini.Load(filePath);	if err != nil {return errors.Wrap(err, "could not load configuration file \""+filePath+"\"")}
	err = cfgIniFile.Section(confSectionName).ReflectFrom(&Cfg);	if err != nil {return errors.Wrap(err, "could not reflect configuration into ini")}
	err = cfgIniFile.Section(internalConfSectionName).ReflectFrom(&Cfg.internal);	if err != nil {return errors.Wrap(err, "could not reflect internal configuration into ini")}
	err = cfgIniFile.SaveTo(filePath);	if err != nil {return errors.Wrap(err, "could not save configuration to file")}

	return nil
}

func updateFromInternals() error {
	//TODO: refine error management
	errMsg := "could not process internal"

	if Cfg.internal.Fullscreen {
		err := Fullscreen();	if err != nil {return errors.Wrap(err, errMsg)}
	} else {
		err := Windowed();	if err != nil {return errors.Wrap(err, errMsg)}
	}

	err := SetWindowSize(Cfg.internal.WindowWidth, Cfg.internal.WindowWidth);	if err != nil {return errors.Wrap(err, errMsg)}
	err = SetVirtualSize(Cfg.internal.VWidth, Cfg.internal.VHeight);	if err != nil {return errors.Wrap(err, errMsg)}

	return nil
}
