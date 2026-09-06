package consul

import (
	stdconsul "github.com/hashicorp/consul/api"

	"github.com/go-kit/log"
)

type Registrar struct {
	client       Client
	registration *stdconsul.AgentServiceRegistration
	logger       log.Logger
}

func NewRegistrar(client Client, r *stdconsul.AgentServiceRegistration, logger log.Logger) *Registrar {
	_ = "STUB: not implemented"
	return nil
}

func (p *Registrar) Register() { _ = "STUB: not implemented"; return }

func (p *Registrar) Deregister() { _ = "STUB: not implemented"; return }
