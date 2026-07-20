package consul

import (
	consul "github.com/hashicorp/consul/api"
)

type Client interface {
	Register(r *consul.AgentServiceRegistration) error

	Deregister(r *consul.AgentServiceRegistration) error

	Service(service, tag string, passingOnly bool, queryOpts *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error)
}

type client struct {
	consul *consul.Client
}

func NewClient(c *consul.Client) Client { _ = "STUB: not implemented"; return *new(Client) }

func (c *client) Register(r *consul.AgentServiceRegistration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) Deregister(r *consul.AgentServiceRegistration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) Service(service, tag string, passingOnly bool, queryOpts *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
