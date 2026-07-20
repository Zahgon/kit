package eureka

import (
	"github.com/hudl/fargo"

	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/internal/instance"
	"github.com/go-kit/log"
)

type Instancer struct {
	cache  *instance.Cache
	conn   fargoConnection
	app    string
	logger log.Logger
	quitc  chan chan struct{}
}

func NewInstancer(conn fargoConnection, app string, logger log.Logger) *Instancer {
	_ = "STUB: not implemented"
	return nil
}

func (s *Instancer) Stop() { _ = "STUB: not implemented"; return }

func (s *Instancer) consume(update fargo.AppUpdate) { _ = "STUB: not implemented"; return }

func (s *Instancer) loop(updates <-chan fargo.AppUpdate, done chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

func (s *Instancer) getInstances() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func convertFargoAppToInstances(app *fargo.Application) []string {
	_ = "STUB: not implemented"
	return nil
}

func (s *Instancer) Register(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }

func (s *Instancer) Deregister(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }

func (s *Instancer) state() sd.Event { _ = "STUB: not implemented"; return *new(sd.Event) }
