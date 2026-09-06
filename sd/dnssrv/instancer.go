package dnssrv

import (
	"errors"
	"time"

	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/internal/instance"
	"github.com/go-kit/log"
)

var ErrPortZero = errors.New("resolver returned SRV record with port 0")

type Instancer struct {
	cache  *instance.Cache
	name   string
	logger log.Logger
	quit   chan struct{}
}

func NewInstancer(
	name string,
	ttl time.Duration,
	logger log.Logger,
) *Instancer {
	_ = "STUB: not implemented"
	return nil
}

func NewInstancerDetailed(
	name string,
	refresh *time.Ticker,
	lookup Lookup,
	logger log.Logger,
) *Instancer {
	_ = "STUB: not implemented"
	return nil
}

func (in *Instancer) Stop() { _ = "STUB: not implemented"; return }

func (in *Instancer) loop(t *time.Ticker, lookup Lookup) { _ = "STUB: not implemented"; return }

func (in *Instancer) resolve(lookup Lookup) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Instancer) Register(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }

func (in *Instancer) Deregister(ch chan<- sd.Event) { _ = "STUB: not implemented"; return }
