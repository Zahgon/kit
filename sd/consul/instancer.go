package consul

import (
	"errors"

	consul "github.com/hashicorp/consul/api"

	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/internal/instance"
	"github.com/go-kit/log"
)

const defaultIndex = 0

var errStopped = errors.New("quit and closed consul instancer")

type Instancer struct {
	cache       *instance.Cache
	client      Client
	logger      log.Logger
	service     string
	tags        []string
	passingOnly bool
	quitc       chan struct{}
}

func NewInstancer(client Client, logger log.Logger, service string, tags []string, passingOnly bool) *Instancer {
	_ = "STUB: not implemented"
	return nil
}

func (s *Instancer) Stop() { _ = "STUB: not implemented"; return }

func (s *Instancer) loop(lastIndex uint64) { _ = "STUB: not implemented"; return }

func (s *Instancer) getInstances(lastIndex uint64, interruptc chan struct{}) ([]string, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Instancer) Register(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }

func (s *Instancer) Deregister(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }

func filterEntries(entries []*consul.ServiceEntry, tags ...string) []*consul.ServiceEntry {
	_ = "STUB: not implemented"
	return nil
}

func makeInstances(entries []*consul.ServiceEntry) []string { _ = "STUB: not implemented"; return nil }
