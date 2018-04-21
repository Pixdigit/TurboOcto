package TurboOcto

import (
    "testing"
    "fmt"
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
        if err != nil {wrapErr(err, "error while serializing", t)} else {
            dec, err := deserialize(enc)
            if err != nil {wrapErr(err, "error while deserializing", t)} else {
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
        err := AddConf(k, v);    if err != nil {wrapErr(err, "error while adding configuration", t)}
        confValue, err := GetConf(k);    if err != nil {wrapErr(err, "could not read back configuration", t)}
        test(confValue == v, "configuration is not equal to set value", t)
    }

    oldConf := map[string]string{}
    for k, v := range conf {
        oldConf[k] = v
    }

    err := SaveConf("testFilename");    if err != nil {wrapErr(err, "could not save conf", t)}
    err = LoadConf("testFilename");    if err != nil {wrapErr(err, "could not load conf", t)}

    for k, v := range oldConf {
        confValue, err := GetConf(k);    if err != nil {wrapErr(err, "could not read back configuration \"" + k + "\"", t)}
        confValueStr, err := serialize(confValue);    test(confValueStr == v, "configuration " + v + " is not equal to set value " + confValueStr, t)
    }

    //should not exist
    err = DelConf("sdfg");    if err == nil {wrapErr(err, "did not recieve error for deleting non existant conf", t)}
    err = DelConf("test1");    if err != nil {wrapErr(err, "could not delete test conf value", t)}
    _, ok := conf["test1"]
    test(!ok, "conf still exists after deleting", t)

}
