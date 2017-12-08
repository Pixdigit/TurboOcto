package TurboOcto

import (
    "github.com/ByteArena/box2d"
	"github.com/veandco/go-sdl2/sdl"
    "github.com/Pixdigit/TurboOcto/internal/sharedStructs"
    "github.com/go-ini/ini"
    "github.com/pkg/errors"
    "strings"
)


func LoadSpritesFromFile(filename string, env Environment) ([]sharedStructs.Sprite, error) {
    conf, err := ini.Load(filename)
    if err != nil {return []sharedStructs.Sprite{}, errors.Wrap(err, "Could not load sprite file")}
    //TODO: this error is irrelevant since there is always default section
    if len(conf.Sections()) == 0 {return []sharedStructs.Sprite{}, errors.New("Config file has no sections " + filename)}

    var sprites []sharedStructs.Sprite

    for _, section := range conf.Sections() {
        name := section.Name()

        if name == "DEFAULT" {
            continue
        }

        imagePath := section.Key("image").String()
        image, err := env.renderEngine.LoadImage(imagePath, env.world)
        if err != nil {return []sharedStructs.Sprite{}, errors.Wrap(err, "Error reading image from config " + filename)}

        _, _, width, height, err := image.Query()
        if err != nil {return []sharedStructs.Sprite{}, errors.Wrap(err, "Error reading image from config " + filename)}
        rect := &sdl.Rect{0, 0, width, height}

        bodyDef := box2d.MakeB2BodyDef()
    	bodyDef.Position.Set(0, 0)
        bodyType := section.Key("type").String()
        //solid B2BodyType
        if pos := strings.IndexRune(bodyType, 's'); pos == 0 {
            bodyDef.Type = box2d.B2BodyType.B2_staticBody
            //dynamic B2BodyType
        } else if pos := strings.IndexRune(bodyType, 'd'); pos == 0 {
            bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
            //kinematic B2BodyType
        } else if pos := strings.IndexRune(bodyType, 'k'); pos == 0 {
            bodyDef.Type = box2d.B2BodyType.B2_kinematicBody
        } else {
            return []sharedStructs.Sprite{}, errors.New("Could not understand BodyType " + bodyType + " from config file " + filename + " in section " + name)
        }
    	body := env.world.CreateBody(&bodyDef)


        fixtureDef := box2d.MakeB2FixtureDef()

        //TODO: read neccessary values
        bodyShape := section.Key("shape").String()
        if bodyShape == "circle" {
            shape := box2d.MakeB2CircleShape()
        	if section.HasKey("radius") {
                shape.B2Shape.M_radius, err = section.Key("radius").Float64()
                if err != nil {return []sharedStructs.Sprite{}, errors.Wrap(err, "Could not read radius from " + filename + " in section " + name)}
            } else {
                return []sharedStructs.Sprite{}, errors.New("Shape is circle but sections has no radius in config file " + filename + " in section " + name)
            }
            fixtureDef.Shape = shape
        } else if bodyShape == "polygon" {
            //TODO
            shape := box2d.MakeB2CircleShape()
        	shape.B2Shape.M_radius = 0.5
            fixtureDef.Shape = shape
        } else if bodyType == "chain" {
            //TODO
            shape := box2d.MakeB2CircleShape()
        	shape.B2Shape.M_radius = 0.5
            fixtureDef.Shape = shape
        } else {
            return []sharedStructs.Sprite{}, errors.New("Could not understand Shape " + bodyShape + " from config file " + filename + " in section " + name)
        }

        fixtureDef.Density = 1
        fixtureDef.Friction = 0.3

        body.CreateFixtureFromDef(&fixtureDef)

        sprites = append(sprites, sharedStructs.Sprite{Texture: image, Rect: rect, Body: body})
    }

    return sprites, nil
}
