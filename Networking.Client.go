package turboOcto

import (
	"bufio"
	"net"

	"github.com/pkg/errors"
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
	conn, err := net.Dial(c.protocol, c.serverAddress)
	if err != nil {
		return client{}, err
	}
	c.rw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	return c, nil
}

func (c *client) Send(key string, data interface{}) error {
	dataString, err := serialize(data)
	if err != nil {
		return errors.Wrap(err, "unable to send data")
	}
	sanitizedDataStr, err := sanitize(dataString, EOT_RUNE)
	if err != nil {
		return errors.Wrap(err, "could not sanitize data")
	}
	sanitizedKey, err := sanitize(key, ESCAPE_RUNE)
	if err != nil {
		return errors.Wrap(err, "could not sanitize key")
	}
	//Append escape sequence
	sanitizedStr := sanitizedKey + string(ESCAPE_RUNE) + sanitizedDataStr + string(EOT_RUNE) + " "
	_, err = c.rw.Write([]byte(sanitizedStr))
	if err != nil {
		return errors.Wrap(err, "unable to send data")
	}
	err = c.rw.Flush()
	if err != nil {
		return errors.Wrap(err, "unable to send data")
	}
	return nil
}

func sanitize(str string, escapeRunes ...rune) (string, error) {
	sanitizedStr := ""
	for _, char := range str {
		for _, escapeRune := range escapeRunes {
			if char == escapeRune {
				sanitizedStr += string(escapeRune)
			}
		}
		sanitizedStr += string(char)
	}
	return sanitizedStr, nil
}
