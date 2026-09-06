package etcdv3

import (
	"context"
	"errors"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
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

	LeaseID() int64
}

type client struct {
	cli *clientv3.Client
	ctx context.Context

	kv clientv3.KV

	watcher clientv3.Watcher

	wctx context.Context

	wcf context.CancelFunc

	leaseID clientv3.LeaseID

	hbch <-chan *clientv3.LeaseKeepAliveResponse

	leaser clientv3.Lease
}

type ClientOptions struct {
	Cert          string
	Key           string
	CACert        string
	DialTimeout   time.Duration
	DialKeepAlive time.Duration

	DialOptions []grpc.DialOption

	Username string
	Password string
}

func NewClient(ctx context.Context, machines []string, options ClientOptions) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (c *client) LeaseID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *client) GetEntries(key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) WatchPrefix(prefix string, ch chan struct{}) { _ = "STUB: not implemented"; return }

func (c *client) Register(s Service) error { _ = "STUB: not implemented"; return nil }

func (c *client) Deregister(s Service) error { _ = "STUB: not implemented"; return nil }

func (c *client) close() { _ = "STUB: not implemented"; return }
