package dnssrv

import "net"

type Lookup func(service, proto, name string) (cname string, addrs []*net.SRV, err error)
