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

func initializeEnvironment() (err error) {
    //TODO: Actually use configuration

    if ok, err := pathExists("./conf"); err != nil {return errors.Wrap(err, "could not check for configuration")
    } else if ok {
        if ok, err := pathExists("./conf/last.csv"); err != nil {return errors.Wrap(err, "failed to check for last configuration")
        } else if ok {
            if err := LoadConf("last"); err != nil {return errors.Wrap(err, "could not load last configuration")}
        } else {
            if err := LoadDefaultConf(); err != nil {return errors.Wrap(err, "could not initialize environment")}
        }
    } else {
        if err := LoadDefaultConf(); err != nil {return errors.Wrap(err, "could not initialize environment")}
    }

    displayIndex, err := window.GetDisplayIndex()
    if err != nil {return errors.Wrap(err, "could not get display index")}
    dmode, err := sdl.GetDesktopDisplayMode(displayIndex)
    if err != nil {return errors.Wrap(err, "could not get display mode")}
    maxWidth, maxHeight = dmode.W, dmode.H

    return nil
}
func LoadDefaultConf() error {
    if ok, err := pathExists("./conf/default.csv"); err != nil {return errors.Wrap(err, "failed to check for default configuration")
    } else if ok {
        if err := LoadConf("default"); err != nil {return errors.Wrap(err, "could not load default configuration")}
    } else {
        //Default Configuration
        conf["updateOnRefresh"] = "bool:true"
        conf["fullscreen"] = "bool:true"
    }
    return nil
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

func typeOfSerialized(s string) (string, error) {
    if !strings.Contains(s, ":") {return "", errors.New("record has untyped values")}
    return s[:strings.Index(s, ":")], nil
}

func deserialize(raw string) (result interface{}, err error) {
    varType, err := typeOfSerialized(raw)
    if err != nil {return nil, errors.Wrap(err, "could not deserialize \"" + raw + "\"")}
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
    if err != nil {return nil, errors.New("invalid configuration value \"" + config + "\"")}
    return value, nil
}

func SetConf(confName string, confValue interface{}) error {
    config, ok := conf[confName]
    if !ok {return errors.New("configuration " + confName + " does not exist")}
    newConfig, err := serialize(confValue);    if err != nil {return errors.Wrap(err, "could not change configuration")}
    oldVarType, err := typeOfSerialized(config);    if err != nil {return errors.Wrap(err, "could not check var type")}
    newVarType, err := typeOfSerialized(config);    if err != nil {return errors.Wrap(err, "could not check var type")}

    if oldVarType != newVarType {return errors.New("configuration must have same type")}
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
    r.FieldsPerRecord = 2
	data, err := r.ReadAll()
    if err != nil {return errors.Wrap(err, "could not read configuration from file " + filename)}

    for _, configuration := range(data) {
        //test if value can be deserialized == valid value
        _, err := deserialize(configuration[1]);    if err != nil {return errors.Wrap(err, "could not load conf file \"" + filename + "\"")}
        conf[configuration[0]] = configuration[1]
    }

    return nil
}
