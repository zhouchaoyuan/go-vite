package p2p

import (
	"math/rand"
	"net"
	"testing"
)

func TestMsg(t *testing.T) {
	for i := 0; i < 100; i++ {
		length := rand.Intn(int(maxPayloadSize))

		t.Logf("rand payload length %d\n", length)
		payload := make([]byte, length)

		rand.Read(payload)

		m := Msg{
			Code:    rand.Uint64(),
			Payload: payload,
		}

		_, err := pack(m)
		if err != nil {
			t.Fatalf("pack msg error: %v\n", err)
		}
	}
}

type mockConn struct {
	net.Conn
}

func (c mockConn) Read(buf []byte) (int, error) {
	rand.Read(buf)
	return len(buf), nil
}
func TestReadFullBytes(t *testing.T) {
	for i := 0; i < 100; i++ {
		conn, _ := net.Pipe()
		mConn := mockConn{
			Conn: conn,
		}

		length := rand.Intn(int(maxPayloadSize))
		t.Logf("rand read length %d\n", length)

		buf := make([]byte, length)
		err := readFullBytes(mConn, buf)
		if err != nil {
			t.Fatalf("readfullbytes error: %v\n", err)
		}
	}
}
