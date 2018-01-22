package compute

import (
	"fmt"
	"strings"

	gcp "google.golang.org/api/compute/v1"
)

type networksClient interface {
	ListNetworks() (*gcp.NetworkList, error)
	DeleteNetwork(network string) error
}

type Networks struct {
	client networksClient
	logger logger
}

func NewNetworks(client networksClient, logger logger) Networks {
	return Networks{
		client: client,
		logger: logger,
	}
}

func (n Networks) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	networks, err := n.client.ListNetworks()
	if err != nil {
		return delete, fmt.Errorf("Listing networks: %s", err)
	}

	for _, network := range networks.Items {
		if !strings.Contains(network.Name, filter) {
			continue
		}

		proceed := n.logger.Prompt(fmt.Sprintf("Are you sure you want to delete network %s?", network.Name))
		if !proceed {
			continue
		}

		delete[network.Name] = ""
	}

	return delete, nil
}

func (n Networks) Delete(networks map[string]string) {
	for name, _ := range networks {
		err := n.client.DeleteNetwork(name)

		if err != nil {
			n.logger.Printf("ERROR deleting network %s: %s\n", name, err)
		} else {
			n.logger.Printf("SUCCESS deleting network %s\n", name)
		}
	}
}
