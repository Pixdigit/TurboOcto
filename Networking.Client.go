package turboOcto

import (
	"bufio"
	"github.com/pkg/errors"
	"net"
)

type client struct {
	protocol      string
	serverAddress string
	isConnected   bool
	rw            *bufio.ReadWriter
}

func TestClient(address string) error {
	return nil
}

func NewClient(address string, protocol string) (client, error) {
	var c client
	c.serverAddress = address
	for _, v := range protocols {
		if v == protocol {
			c.protocol = protocol
		}
	}
	if c.protocol == "" {
		return client{}, errors.New("unknown protocol \"" + protocol + "\"")
	}
	conn, err := net.Dial(c.protocol, c.serverAddress);	if err != nil {return client{}, err}
	c.rw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	return c, nil
}

func (c *client) Send(data interface{}) error {
	dataString, err := serialize(data);	if err != nil {return errors.Wrap(err, "unable to send data")}
	sanitizedStr := ""
	for _, char := range dataString {
		if char == ESCAPE_RUNE {
			sanitizedStr += string(ESCAPE_RUNE)
		}
		sanitizedStr += string(char)
	}
	//Append escape sequence
	sanitizedStr = sanitizedStr + string(ESCAPE_RUNE) + " "
	_, err = c.rw.Write([]byte(sanitizedStr));	if err != nil {return errors.Wrap(err, "unable to send data")}
	err = c.rw.Flush();	if err != nil {return errors.Wrap(err, "unable to send data")}
	return nil
}
