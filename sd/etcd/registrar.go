package etcd

import (
	"sync"
	"time"

	etcd "go.etcd.io/etcd/client/v2"

	"github.com/go-kit/log"
)

const minHeartBeatTime = 500 * time.Millisecond

type Registrar struct {
	client  Client
	service Service
	logger  log.Logger

	quitmtx sync.Mutex
	quit    chan struct{}
}

type Service struct {
	Key           string
	Value         string
	TTL           *TTLOption
	DeleteOptions *etcd.DeleteOptions
}

type TTLOption struct {
	heartbeat time.Duration
	ttl       time.Duration
}

func NewTTLOption(heartbeat, ttl time.Duration) *TTLOption { _ = "STUB: not implemented"; return nil }

func NewRegistrar(client Client, service Service, logger log.Logger) *Registrar {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registrar) Register() { _ = "STUB: not implemented"; return }

func (r *Registrar) loop() { _ = "STUB: not implemented"; return }

func (r *Registrar) Deregister() { _ = "STUB: not implemented"; return }
