package entx

import (
	"math/rand"
	
	"github.com/thoohv5/template/internal/ent"
)

type Policy interface {
	Resolve([]*ent.Client) *ent.Client
}

type RandomClientSetPolicy struct{}

func (p *RandomClientSetPolicy) Resolve(clients []*ent.Client) *ent.Client {
	n := rand.Intn(len(clients))
	return clients[n]
}

type RoundRobinPolicy struct {
	n int
}

func (p *RoundRobinPolicy) Resolve(clients []*ent.Client) *ent.Client {
	n := p.n % len(clients)
	p.n++
	return clients[n]
}
