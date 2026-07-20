package etcdv3

import (
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/internal/instance"
	"github.com/go-kit/log"
)

type Instancer struct {
	cache  *instance.Cache
	client Client
	prefix string
	logger log.Logger
	quitc  chan struct{}
}

func NewInstancer(c Client, prefix string, logger log.Logger) (*Instancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Instancer) loop() { _ = "STUB: not implemented"; return }

func (s *Instancer) Stop() { _ = "STUB: not implemented"; return }

func (s *Instancer) Register(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }

func (s *Instancer) Deregister(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }
