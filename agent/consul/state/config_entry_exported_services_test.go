// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package state

import (
	"testing"

	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/private/pbconfigentry"
	"github.com/hashicorp/go-memdb"
	"github.com/stretchr/testify/require"
)

func TestStore_ResolvedExportingServices(t *testing.T) {
	s := NewStateStore(nil)
	var c indexCounter

	{
		require.NoError(t, s.EnsureNode(c.Next(), &structs.Node{
			Node: "foo", Address: "127.0.0.1",
		}))

		require.NoError(t, s.EnsureService(c.Next(), "foo", &structs.NodeService{
			ID: "db", Service: "db", Port: 5000,
		}))

		require.NoError(t, s.EnsureService(c.Next(), "foo", &structs.NodeService{
			ID: "cache", Service: "cache", Port: 5000,
		}))

		entry := &structs.ExportedServicesConfigEntry{
			Name: "default",
			Services: []structs.ExportedService{
				{
					Name: "db",
					Consumers: []structs.ServiceConsumer{
						{
							Peer: "east",
						},
						{
							Peer: "west",
						},
					},
				},
				{
					Name: "cache",
					Consumers: []structs.ServiceConsumer{
						{
							Peer: "east",
						},
					},
				},
			},
		}
		err := s.EnsureConfigEntry(c.Next(), entry)
		require.NoError(t, err)

		// Adding services to check wildcard config later on

		require.NoError(t, s.EnsureService(c.Next(), "foo", &structs.NodeService{
			ID: "frontend", Service: "frontend", Port: 5000,
		}))

		require.NoError(t, s.EnsureService(c.Next(), "foo", &structs.NodeService{
			ID: "backend", Service: "backend", Port: 5000,
		}))

		// The consul service should never be exported.
		require.NoError(t, s.EnsureService(c.Next(), "foo", &structs.NodeService{
			ID: structs.ConsulServiceID, Service: structs.ConsulServiceName, Port: 8000,
		}))

	}

	type testCase struct {
		expect []*pbconfigentry.ResolvedExportedService
		idx    uint64
	}

	run := func(t *testing.T, tc testCase) {
		ws := memdb.NewWatchSet()
		defaultMeta := structs.DefaultEnterpriseMetaInDefaultPartition()
		idx, services, err := s.ResolvedExportedServices(ws, defaultMeta)
		require.NoError(t, err)
		require.Equal(t, tc.idx, idx)
		require.ElementsMatch(t, tc.expect, services)
	}

	t.Run("only exported services are included", func(t *testing.T) {
		tc := testCase{
			expect: []*pbconfigentry.ResolvedExportedService{
				{
					Service: "db",
					Consumers: &pbconfigentry.Consumers{
						Peers: []string{"east", "west"},
					},
				},
				{
					Service: "cache",
					Consumers: &pbconfigentry.Consumers{
						Peers: []string{"east"},
					},
				},
			},
			idx: 4,
		}

		run(t, tc)
	})

	t.Run("wild card includes all services", func(t *testing.T) {
		entry := &structs.ExportedServicesConfigEntry{
			Name: "default",
			Services: []structs.ExportedService{
				{
					Name: "*",
					Consumers: []structs.ServiceConsumer{
						{Peer: "west"},
					},
				},
			},
		}

		err := s.EnsureConfigEntry(c.Next(), entry)
		require.NoError(t, err)

		tc := testCase{
			expect: []*pbconfigentry.ResolvedExportedService{
				{
					Service: "db",
					Consumers: &pbconfigentry.Consumers{
						Peers: []string{"west"},
					},
				},
				{
					Service: "cache",
					Consumers: &pbconfigentry.Consumers{
						Peers: []string{"west"},
					},
				},
				{
					Service: "frontend",
					Consumers: &pbconfigentry.Consumers{
						Peers: []string{"west"},
					},
				},
				{
					Service: "backend",
					Consumers: &pbconfigentry.Consumers{
						Peers: []string{"west"},
					},
				},
			},
			idx: c.Last(),
		}

		run(t, tc)
	})

	t.Run("deleting the config entry clears the services", func(t *testing.T) {
		defaultMeta := structs.DefaultEnterpriseMetaInDefaultPartition()
		err := s.DeleteConfigEntry(c.Next(), structs.ExportedServices, "default", nil)
		require.NoError(t, err)

		idx, result, err := s.ResolvedExportedServices(nil, defaultMeta)
		require.NoError(t, err)
		require.Equal(t, c.Last(), idx)
		require.Nil(t, result)
	})
}
