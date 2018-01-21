package fakes

import compute "google.golang.org/api/compute/v1"

type NetworksClient struct {
	ListNetworksCall struct {
		CallCount int
		Receives  struct {
			Filter string
		}
		Returns struct {
			Output *compute.NetworkList
			Error  error
		}
	}

	DeleteNetworkCall struct {
		CallCount int
		Receives  struct {
			Network string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *NetworksClient) ListNetworks(filter string) (*compute.NetworkList, error) {
	n.ListNetworksCall.CallCount++
	n.ListNetworksCall.Receives.Filter = filter

	return n.ListNetworksCall.Returns.Output, n.ListNetworksCall.Returns.Error
}

func (n *NetworksClient) DeleteNetwork(network string) error {
	n.DeleteNetworkCall.CallCount++
	n.DeleteNetworkCall.Receives.Network = network

	return n.DeleteNetworkCall.Returns.Error
}
