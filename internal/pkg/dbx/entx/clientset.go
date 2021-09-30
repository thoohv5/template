package entx

import "github.com/thoohv5/template/internal/ent"

type ClientType int

type ClientSet struct {
	primary  *ent.Client
	replicas []*ent.Client
	policy   Policy
}

type ClientSetOption func(*ClientSet)

func NewClientSet(primary *ent.Client, opts ...ClientSetOption) *ClientSet {
	cs := &ClientSet{
		primary: primary,
		policy:  &RandomClientSetPolicy{},
	}
	for _, fn := range opts {
		fn(cs)
	}
	return cs
}

func WithReplicas(replicas []*ent.Client) ClientSetOption {
	return func(cs *ClientSet) {
		cs.replicas = replicas
	}
}

func WithPolicy(policy Policy) ClientSetOption {
	return func(cs *ClientSet) {
		cs.policy = policy
	}
}

func (c *ClientSet) Primary() *ent.Client {
	return c.primary
}

func (c *ClientSet) Replica() *ent.Client {
	if len(c.replicas) == 0 {
		return c.primary
	}

	return c.policy.Resolve(c.replicas)
}

func (c *ClientSet) Close() error {
	for _, replica := range c.replicas {
		if err := replica.Close(); nil != err {
			return err
		}
	}

	return c.primary.Close()
}
