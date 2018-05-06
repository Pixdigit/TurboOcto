package TurboOcto

import (
    "net"
    "bufio"
    "encoding/gob"
)

func TestClient(address string) error {
    conn, err := net.Dial("tcp", address);    if (err != nil) {return err}
    rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    type dataForm struct {Test string}
    data := dataForm{Test: "uiuiui"}
    enc := gob.NewEncoder(rw)
    err = enc.Encode(data);    if (err != nil) {return err}
    rw.Flush()
    return nil
}
