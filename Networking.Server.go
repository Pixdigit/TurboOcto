package turboOcto

import (
	"bufio"
	"github.com/pkg/errors"
	"fmt"
	"net"
)

var protocols = [...]string{"tcp", "tcp4", "tcp6", "udp", "udp4", "udp6"}
const ESCAPE_RUNE = rune('/')


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
				//one "character" or more if waiting for another /
				var token []rune
				data := ""
				var datArray []string
				for {
					thisRune, _, err := r.ReadRune()
					token = append(token, thisRune)
					if err != nil {
						//TODO: send notification of faulty msg to client
						errChan <- err
						data = ""
						token = []rune{}
					}
					//Token is of max length 2
					if len(token) == 1 && token[0] != ESCAPE_RUNE {
						//Single "character"
						data += string(token[0])
						token = []rune{}
					} else if len(token) == 2 && token[0] == ESCAPE_RUNE {
						if token[0] == ESCAPE_RUNE && token[1] == ESCAPE_RUNE {
							//Escaped escape character
							token = []rune{ESCAPE_RUNE}
						} else {
							//Recieved single / as end statement
							fmt.Println(data)
							datArray = append(datArray, data)
							data = ""
							token = []rune{token[1]}
						}
						data += string(token[0])
						token = []rune{}
					}

				}
			}
		}
	}()
}
