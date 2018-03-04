package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/pkg/errors"
    "fmt"
    "strings"
    "strconv"
)

var renderer *sdl.Renderer
var window *sdl.Window
var displayIndex int //TODO: Dynamically update when window moved

var conf map[string]string = map[string]string{}

func init() {
    conf["updateOnRefresh"] = "bool:true"
}

func serialize(variable interface{}) (string, error) {
    var result string
    var err error

    switch t := variable.(type) {
    case bool:
        if t {
            result = "bool:true"
        } else {
            result = "bool:false"
        }
    case int:
        result = "int:" + strconv.Itoa(t)
    case float32:
        result = "float:" + strconv.FormatFloat(float64(t), 'G', -1, 64)
    case float64:
        result = "float:" + strconv.FormatFloat(t, 'G', -1, 64)
    case string:
        result = "str:" + t
    default:
        err = errors.New(fmt.Sprintf("Can not deserialize %#v of type %T", t, t))
    }
    return result, err
}

func deserialize(raw string) (result interface{}, err error) {
    varType := raw[:strings.Index(raw, ":")]
    varValue := raw[strings.Index(raw, ":") + 1:]

    switch varType {
    case "bool":
        if varValue == "true" {
            result = true
        } else if varValue == "false" {
            result = false
        } else {
            err = errors.New("could not deserialize " + varValue + " as boolean")
        }
    case "int":
        result, err = strconv.Atoi(varValue)
        errors.Wrap(err, "could not deserialize " + varValue + " as int")
    case "float":
        result, err = strconv.ParseFloat(varValue, 64) //TODO: use future conf infrastructure
        errors.Wrap(err, "could not deserialize " + varValue + " as float")
    case "str":
        result = varValue
    default:
        err = errors.New("unknown var type in deserialisation: " + varType)
    }
    return
}
