package zk

import "github.com/go-kit/log"

type Registrar struct {
	client  Client
	service Service
	logger  log.Logger
}

type Service struct {
	Path string
	Name string
	Data []byte
	node string
}

func NewRegistrar(client Client, service Service, logger log.Logger) *Registrar {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registrar) Register() { _ = "STUB: not implemented"; return }

func (r *Registrar) Deregister() { _ = "STUB: not implemented"; return }
