package TurboOcto

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/pkg/errors"
    "fmt"
    "strings"
    "strconv"
    "encoding/csv"
    "os"
)

var renderer *sdl.Renderer
var window *sdl.Window
var displayIndex int //TODO: Dynamically update when window moved

var conf map[string]string = map[string]string{}

func init() {
    //TODO: Actually use configuration
    conf["updateOnRefresh"] = "bool:true"
    conf["fullscreen"] = "bool:true"
    displayIndex, _ := window.GetDisplayIndex()
    dmode, _ := sdl.GetDesktopDisplayMode(displayIndex)
    maxWidth, maxHeight = dmode.W, dmode.H
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

func typeOfSerialized(s string) string {
    return s[:strings.Index(s, ":")]
}

func deserialize(raw string) (result interface{}, err error) {
    varType := typeOfSerialized(raw)
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
        err = errors.New("unknown var type in deserialisation " + varType)
    }
    return
}

func GetConf(confName string) (interface{}, error) {
    config, ok := conf[confName]
    if !ok {return nil, errors.New("configuration " + confName + " does not exist")}
    value, err := deserialize(config)
    if err != nil {return nil, errors.New("invalid configuration value")}
    return value, nil
}

func SetConf(confName string, confValue interface{}) error {
    config, ok := conf[confName]
    if !ok {return errors.New("configuration " + confName + " does not exist")}
    newConfig, err := serialize(confValue)
    if err != nil {return errors.Wrap(err, "could not change configuration")}
    if typeOfSerialized(config) != typeOfSerialized(newConfig) {return errors.New("configuration must have same type")}
    conf[confName] = newConfig
    return nil
}

func AddConf(confName string, initConfValue interface{}) error {
    newConfig, err := serialize(initConfValue)
    if err != nil {return errors.Wrap(err, "could not add configuration")}
    conf[confName] = newConfig
    return nil
}

func SaveConf(filename string) error {
    var data [][]string
    for k, v := range(conf) {
        configuration := []string{k, v}
        data = append(data, configuration)
    }

    if ok, err := pathExists("./conf/"); err != nil {
        return errors.Wrap(err, "could not check wether conf folder exists")
    } else if !ok {
        err = os.Mkdir("./conf", os.ModePerm)
        if err != nil {return errors.Wrap(err, "could not create conf folder")}
    }

    file, err := os.Create("./conf/" + filename + ".csv")
    defer file.Close()
    if err != nil {return errors.Wrap(err, "could not open file " + filename + " to save configuration")}
    w := csv.NewWriter(file)
	err = w.WriteAll(data)
    if err != nil {return errors.Wrap(err, "could not write configuration to file" + filename)}

    return nil
}

func LoadConf(filename string) error {
    file, err := os.Open("./conf/" + filename + ".csv")
    defer file.Close()
    if err != nil {return errors.Wrap(err, "could not open file " + filename + " to load configuration")}
    r := csv.NewReader(file)
	data, err := r.ReadAll()
    if err != nil {return errors.Wrap(err, "could not read configuration from file " + filename)}

    for _, configuration := range(data) {
        conf[configuration[0]] = configuration[1]
    }

    return nil
}
