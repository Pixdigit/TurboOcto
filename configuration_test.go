package TurboOcto

import (
    "testing"
    "fmt"
    "github.com/pkg/errors"
)

func TestSerialization(t *testing.T) {
//    var testVars []interface{}
    testVars := []interface{}{
        true, false, "true",
        3, 456, -234, -0,
        3.1415, -2.1645, -0.0,
        "4", "str:test", ":", "",
    }

    for _, v := range(testVars) {
        enc, err := serialize(v)
        if err != nil {t.Error(errors.Wrap(err, "error while serializing"))} else {
            dec, err := deserialize(enc)
            if err != nil {t.Error(errors.Wrap(err, "error while deserializing"))} else {
                test(dec == v, "var is not the same after de- and serializing: " + fmt.Sprint(v) + " != " + fmt.Sprint(dec), t)
            }
        }
    }
}

func TestConfigurationSystem(t *testing.T) {
    confs := map[string]interface{}{
        "test1": 0,
        "test2": 3.1415,
        "": true,
        "äüß \n ?=):(/&%$§!)": "OKAYdokay",
    }
    for k, v := range confs {
        err := AddConf(k, v)
        if err != nil {t.Error(errors.Wrap(err, "error while adding configuration"))}
        confValue, err := GetConf(k)
        if err != nil {t.Error(errors.Wrap(err, "could not read back configuration"))}
        test(confValue == v, "configuration is not equal to set value", t)
    }

    oldConf := map[string]string{}
    for k, v := range conf {
        oldConf[k] = v
    }

    err := SaveConf("testFilename")
    if err != nil {t.Error(errors.Wrap(err, "could not save conf"))}
    err = LoadConf("testFilename")
    if err != nil {t.Error(errors.Wrap(err, "could not load conf"))}

    for k, v := range oldConf {
        confValue, err := GetConf(k)
        if err != nil {t.Error(errors.Wrap(err, "could not read back configuration \"" + k + "\""))}
        confValueStr, err := serialize(confValue)
        test(confValueStr == v, "configuration " + v + " is not equal to set value " + confValueStr, t)
    }

}
