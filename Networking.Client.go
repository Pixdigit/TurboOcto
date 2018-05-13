package turboOcto

import (
	"bufio"
	"net"
)

func TestClient(address string) error {
	conn, err := net.Dial("tcp", address);	if err != nil {return err}
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	rw.Write([]byte("str:äh/ int:8/ str://test/// "))
	rw.Flush()
	return nil
}
