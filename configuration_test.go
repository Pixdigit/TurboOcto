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
        if err != nil {t.Error("error while serializing:", err)} else {
            dec, err := deserialize(enc)
            if err != nil {t.Error("error while deserializing:", err)} else {
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
        "äüß \n ?=)(/&%$§!)": "OKAYdokay",
    }
    for k, v := range confs {
        err := AddConf(k, v)
        if err != nil {t.Error("error while adding configuration", err)}
        confValue, err := GetConf(k)
        if err != nil {t.Error("could not read back configuration", err)}
        test(confValue == v, "configuration is not equal to set value", t)
    }

}
