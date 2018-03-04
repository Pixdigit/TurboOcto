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
