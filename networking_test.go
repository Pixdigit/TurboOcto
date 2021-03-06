package turboOcto

import (
	"testing"
	"time"
)

func TestNetworking(t *testing.T) {
	address := "localhost:49125"
	protocol := "tcp"
	s, err := NewServer(address, protocol)
	if err != nil {
		wrapErr(err, "could not start server", t)
		t.FailNow()
	}
	end := time.After(100 * time.Millisecond)
	errChan := make(chan error)

	s.Start(errChan)
	defer s.Stop()

	c, err := NewClient(address, protocol)
	if err != nil {
		wrapErr(err, "could not create client", t)
		t.FailNow()
	}

	//TODO: test other data types
	testAgainstStrings(func(str string) error {
		err := c.Send(str, str)
		if err != nil {
			wrapErr(err, "failed to send data \""+str+"\"", t)
		}
		return nil
	}, func() (string, error) {
		gotNew := false
		str := ""
		for !gotNew {
			_, str, gotNew = s.Recv()
		}
		result, err := deserialize(str)
		if err != nil {
			wrapErr(err, "got invalid data", t)
		}
		return result.(string), nil
	}, "failure while data transfer", t)

	run := true
	for run {
		select {
		case err := <-errChan:
			if err != nil {
				wrapErr(err, "error ", t)
			} else {
				t.Log("some function pushed error with value nil")
			}
		case <-end:
			run = false
		default:
		}
	}
}
