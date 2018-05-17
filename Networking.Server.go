package turboOcto

import (
	"bufio"
	"io"
	"net"

	"github.com/pkg/errors"
)

const TCP string = "tcp"
const TCP4 string = "tcp4"
const TCP6 string = "tcp6"
const UDP string = "udp"
const UDP4 string = "udp4"
const UDP6 string = "udp6"

var protocols = [...]string{TCP, TCP4, TCP6, UDP, UDP4, UDP6}

//End of key indicator
const ESCAPE_RUNE = rune('/')

//End of value (and transmission) indicator
const EOT_RUNE = rune('!')

type server struct {
	listener     net.Listener
	prototcol    string
	incomingData chan map[string]string
	address      string
	state        Runlevel
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
	s.incomingData = make(chan map[string]string)
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

func (s *server) Recv() (string, string, bool) {
	select {
	case data := <-s.incomingData:
		//Data only consists of one pair
		//Just read out both variables
		for k, v := range data {
			return k, v, true
		}
	default:
	}
	return "", "", false
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
				//one "character" or more if waiting for another escape Rune
				var token []rune
				key := ""
				strData := ""
				dataIsKey := true
				for {
					thisRune, _, err := r.ReadRune()
					token = append(token, thisRune)
					if err != nil {
						//TODO: send notification of faulty msg to client
						pushErr(errChan, err)
						strData = ""
						token = []rune{}
					}
					//Token is of max length 2
					if dataIsKey {
						token, strData, err = readTokenWithEscapeRune(token, strData, ESCAPE_RUNE)
						if err != nil {
							if err == io.EOF {
								key = strData[:len(strData)-1]
								//reset strData buffer to first rune of value strData
								strData = string(strData[len(strData)-1])
								dataIsKey = false
							} else {
								pushErr(errChan, err)
							}
						}
					} else {
						token, strData, err = readTokenWithEscapeRune(token, strData, EOT_RUNE)
						if err != nil {
							if err == io.EOF {
								dataMap := make(map[string]string)
								dataMap[key] = strData[:len(strData)-1]
								go func() { s.incomingData <- dataMap }()
								strData = ""
								key = ""
							} else {
								pushErr(errChan, err)
							}
							dataIsKey = true
						}
					}
				}
			}
		}
	}()
}

func readTokenWithEscapeRune(token []rune, data string, escapeRune rune) ([]rune, string, error) {
	var err error = nil
	if len(token) == 1 && token[0] != escapeRune {
		//Single rune
		data += string(token[0])
		token = []rune{}

	} else if len(token) == 2 && token[0] == escapeRune {
		if token[0] == escapeRune && token[1] == escapeRune {
			//Escaped escape rune
			token = []rune{escapeRune}
		} else {
			//Recieved single escape rune as end statement
			err = io.EOF
			token = []rune{token[1]}
		}
		data += string(token[0])
		token = []rune{}

	} else {
		//Token not correctly formatted
		if len(token) > 2 || len(token) == 0 {
			err = errors.New("token of unusable size")
		} else if len(token) == 2 && token[0] != escapeRune {
			err = errors.New("token longer than 1 but does not begin with escape rune")
		}
	}
	return token, data, err
}
