package eureka

import (
	"sync"
	"time"

	"github.com/hudl/fargo"

	"github.com/go-kit/kit/sd"
	"github.com/go-kit/log"
)

const defaultRenewalInterval = 30 * time.Second

type fargoConnection interface {
	RegisterInstance(instance *fargo.Instance) error
	DeregisterInstance(instance *fargo.Instance) error
	ReregisterInstance(instance *fargo.Instance) error
	HeartBeatInstance(instance *fargo.Instance) error
	ScheduleAppUpdates(name string, await bool, done <-chan struct{}) <-chan fargo.AppUpdate
	GetApp(name string) (*fargo.Application, error)
}

type fargoUnsuccessfulHTTPResponse struct {
	statusCode    int
	messagePrefix string
}

func (u *fargoUnsuccessfulHTTPResponse) Error() string { _ = "STUB: not implemented"; return "" }

type Registrar struct {
	conn     fargoConnection
	instance *fargo.Instance
	logger   log.Logger
	quitc    chan chan struct{}
	sync.Mutex
}

var _ sd.Registrar = (*Registrar)(nil)

func NewRegistrar(conn fargoConnection, instance *fargo.Instance, logger log.Logger) *Registrar {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registrar) Register() { _ = "STUB: not implemented"; return }

func (r *Registrar) Deregister() { _ = "STUB: not implemented"; return }

func (r *Registrar) loop() { _ = "STUB: not implemented"; return }

func httpResponseStatusCode(err error) (code int, present bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func isNotFound(err error) bool { _ = "STUB: not implemented"; return false }

func (r *Registrar) heartbeat() error { _ = "STUB: not implemented"; return nil }
