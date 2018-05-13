package turboOcto

import (
	"bufio"
	//"encoding/gob"
	"github.com/pkg/errors"
	//"io"
	"fmt"
	"net"
)

var protocols = [...]string{"tcp", "tcp4", "tcp6", "udp", "udp4", "udp6"}

type server struct {
	listener  net.Listener
	prototcol string
	address   string
	state     Runlevel
}

func NewServer(address string, protocol string) (server, error) {
	s := server{}
	s.address = address
	s.state = PAUSED
	for _, v := range protocols {
		if v == protocol {
			s.prototcol = protocol
		}
	}
	if s.prototcol == "" {
		return server{}, errors.New("unknown protocol \"" + protocol + "\"")
	}
	return s, nil
}

func (s *server) Start(errChan chan error) {
	var err error
	s.Continue()
	s.listener, err = net.Listen(s.prototcol, s.address)
	if err != nil {
		pushErr(errChan, errors.Wrap(err, "unable to listen on "+s.listener.Addr().String()))
	}

	connChan := make(chan net.Conn)

	s.acceptConnections(connChan, errChan)
	s.handleConnections(connChan, errChan)

}

func (s *server) acceptConnections(connChan chan net.Conn, errChan chan error) {
	go func() {
		for s.state != STOPPED {
			conn, err := s.listener.Accept()
			if err != nil {
				pushErr(errChan, errors.Wrap(err, "failed accepting connection request"))
			} else {
				connChan <- conn
			}
		}
	}()
}

func (s *server) Pause() {
	s.state = PAUSED
}
func (s *server) Continue() {
	s.state = RUNNING
}
func (s *server) Stop() {
	s.state = STOPPED
}

func (s *server) handleConnections(connChan chan net.Conn, errChan chan error) {
	go func() {
		for s.state != STOPPED {
			select {
			case conn := <-connChan:
				s.handleConnection(conn, errChan)
			default:
				//Wait until state is running
				for s.state == PAUSED {
					//TODO: implement return package to indicate paused server
				}
			}
		}
	}()
}

func (s *server) handleConnection(conn net.Conn, errChan chan error) {
	go func() {
		defer conn.Close()
		r := bufio.NewReader(conn)
		for s.state != STOPPED {
			for s.state == RUNNING {
				//dec := gob.NewDecoder(rw)
				//type dataForm struct{ Test string }
				//data := dataForm{Test: ""}
				//err := dec.Decode(&data)
				var datArray []string
				data := ""
				var token []rune
				for {
					thisRune, _, _ := r.ReadRune()
					//Token is of max length 2
					token = append(token, thisRune)
					if len(token) == 1 && token[0] != rune('/'){
						data += string(token[0])
						token = []rune{}
					} else if len(token) == 2 && token[0] == rune('/') {
						if token[0] == rune('/') && token[1] == rune('/') {
							token = []rune{rune('/')}
						} else {
							//Recieved single / as end statement
							datArray = append(datArray, data)
							data = ""
							fmt.Println("NEW")
							for _, dat := range datArray {
								fmt.Println(deserialize(dat))
							}
							token = []rune{token[1]}
						}
						data += string(token[0])
						token = []rune{}
					}

					/*if err != nil {
						if err != io.EOF {
							pushErr(errChan, err)
						}
						return
					}*/
				}
			}
		}
	}()
}
