package conn

import (
	"errors"
	"net"
	"time"

	"github.com/go-kit/log"
)

type Dialer func(network, address string) (net.Conn, error)

type AfterFunc func(time.Duration) <-chan time.Time

type Manager struct {
	dialer  Dialer
	network string
	address string
	after   AfterFunc
	logger  log.Logger

	takec chan net.Conn
	putc  chan error
}

func NewManager(d Dialer, network, address string, after AfterFunc, logger log.Logger) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func NewDefaultManager(network, address string, logger log.Logger) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) Take() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (m *Manager) Put(err error) { _ = "STUB: not implemented"; return }

func (m *Manager) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Manager) loop() { _ = "STUB: not implemented"; return }

func dial(d Dialer, network, address string, logger log.Logger) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

func Exponential(d time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

var ErrConnectionUnavailable = errors.New("connection unavailable")
