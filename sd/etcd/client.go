package etcd

import (
	"context"
	"errors"
	"time"

	etcd "go.etcd.io/etcd/client/v2"
)

var (
	ErrNoKey = errors.New("no key provided")

	ErrNoValue = errors.New("no value provided")
)

type Client interface {
	GetEntries(prefix string) ([]string, error)

	WatchPrefix(prefix string, ch chan struct{})

	Register(s Service) error

	Deregister(s Service) error
}

type client struct {
	keysAPI etcd.KeysAPI
	ctx     context.Context
}

type ClientOptions struct {
	Cert                    string
	Key                     string
	CACert                  string
	DialTimeout             time.Duration
	DialKeepAlive           time.Duration
	HeaderTimeoutPerRequest time.Duration
}

func NewClient(ctx context.Context, machines []string, options ClientOptions) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (c *client) GetEntries(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) WatchPrefix(prefix string, ch chan struct{}) { _ = "STUB: not implemented"; return }

func (c *client) Register(s Service) error { _ = "STUB: not implemented"; return nil }

func (c *client) Deregister(s Service) error { _ = "STUB: not implemented"; return nil }
