package zk

import (
	"errors"
	"time"

	"github.com/go-zookeeper/zk"

	"github.com/go-kit/log"
)

var (
	DefaultACL            = zk.WorldACL(zk.PermAll)
	ErrInvalidCredentials = errors.New("invalid credentials provided")
	ErrClientClosed       = errors.New("client service closed")
	ErrNotRegistered      = errors.New("not registered")
	ErrNodeNotFound       = errors.New("node not found")
)

const (
	DefaultConnectTimeout = 2 * time.Second

	DefaultSessionTimeout = 5 * time.Second
)

type Client interface {
	GetEntries(path string) ([]string, <-chan zk.Event, error)

	CreateParentNodes(path string) error

	Register(s *Service) error

	Deregister(s *Service) error

	Stop()
}

type clientConfig struct {
	logger          log.Logger
	acl             []zk.ACL
	credentials     []byte
	connectTimeout  time.Duration
	sessionTimeout  time.Duration
	rootNodePayload [][]byte
	eventHandler    func(zk.Event)
}

type Option func(*clientConfig) error

type client struct {
	*zk.Conn
	clientConfig
	active bool
	quit   chan struct{}
}

func ACL(acl []zk.ACL) Option { _ = "STUB: not implemented"; return *new(Option) }

func Credentials(user, pass string) Option { _ = "STUB: not implemented"; return *new(Option) }

func ConnectTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func SessionTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func Payload(payload [][]byte) Option { _ = "STUB: not implemented"; return *new(Option) }

func EventHandler(handler func(zk.Event)) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewClient(servers []string, logger log.Logger, options ...Option) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (c *client) CreateParentNodes(path string) error { _ = "STUB: not implemented"; return nil }

func (c *client) GetEntries(path string) ([]string, <-chan zk.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *client) Register(s *Service) error { _ = "STUB: not implemented"; return nil }

func (c *client) Deregister(s *Service) error { _ = "STUB: not implemented"; return nil }

func (c *client) Stop() { _ = "STUB: not implemented"; return }
