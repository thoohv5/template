package entx

import (
	"testing"

	"github.com/stretchr/testify/require"
	
	"github.com/thoohv5/template/internal/ent"
)

func TestRandomClientSetPolicy(t *testing.T) {
	policy := &RandomClientSetPolicy{}
	clients := []*ent.Client{
		{},
		{},
		{},
		{},
		{},
	}

	visitedClients := make(map[*ent.Client]int)
	for i := 0; i < len(clients)*2; i++ {
		cli := policy.Resolve(clients)
		visitedClients[cli]++
	}

	require.Greater(t, len(visitedClients), 1)
	require.LessOrEqual(t, len(visitedClients), len(clients))
}

func TestRoundRobinPolicy(t *testing.T) {
	policy := &RoundRobinPolicy{}
	clients := []*ent.Client{
		{},
		{},
		{},
		{},
		{},
	}

	for i := 0; i < len(clients)*2; i++ {
		cli := policy.Resolve(clients)
		require.Equal(t, clients[i%len(clients)], cli)
	}
}
