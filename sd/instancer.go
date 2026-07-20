package sd

type Event struct {
	Instances []string
	Err       error
}

type Instancer interface {
	Register(chan<- Event)
	Deregister(chan<- Event)
	Stop()
}

type FixedInstancer []string

func (d FixedInstancer) Register(ch chan<- Event) { _ = "STUB: not implemented"; return }

func (d FixedInstancer) Deregister(ch chan<- Event) { _ = "STUB: not implemented"; return }

func (d FixedInstancer) Stop() { _ = "STUB: not implemented"; return }
