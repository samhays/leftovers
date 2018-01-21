package gcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"golang.org/x/oauth2/google"
	gcpcompute "google.golang.org/api/compute/v1"
)

type logger interface {
	Printf(m string, a ...interface{})
	Println(m string)
	Prompt(m string) bool
}

type resource interface {
	Delete(string) error
}

type Deleter struct {
	resources []resource
}

func (d Deleter) Delete(filter string) error {
	for _, r := range d.resources {
		if err := r.Delete(filter); err != nil {
			return err
		}
	}
	return nil
}

func NewDeleter(logger logger, serviceAccountKey string) (Deleter, error) {
	if serviceAccountKey == "" {
		return Deleter{}, errors.New("Missing BBL_GCP_SERVICE_ACCOUNT_KEY.")
	}

	key, err := ioutil.ReadFile(serviceAccountKey)
	if err != nil {
		return Deleter{}, errors.New(fmt.Sprintf("Reading %s: %s", serviceAccountKey, err))
	}

	p := struct {
		ProjectId string `json:"project_id"`
	}{}
	if err := json.Unmarshal(key, &p); err != nil {
		return Deleter{}, errors.New(fmt.Sprintf("Unmarshalling account key for project id: %s", err))
	}

	logger.Println(fmt.Sprintf("Cleaning gcp project: %s.", p.ProjectId))

	config, err := google.JWTConfigFromJSON(key, gcpcompute.ComputeScope)
	if err != nil {
		return Deleter{}, errors.New(fmt.Sprintf("Creating jwt config: %s", err))
	}

	service, err := gcpcompute.New(config.Client(context.Background()))
	if err != nil {
		return Deleter{}, errors.New(fmt.Sprintf("Creating gcp client: %s", err))
	}

	client := compute.NewClient(p.ProjectId, service, logger)

	regions, err := client.ListRegions()
	if err != nil {
		return Deleter{}, errors.New(fmt.Sprintf("Listing regions: %s", err))
	}

	zones, err := client.ListZones()
	if err != nil {
		return Deleter{}, errors.New(fmt.Sprintf("Listing zones: %s", err))
	}

	return Deleter{
		resources: []resource{
			compute.NewFirewalls(client, logger),
			compute.NewForwardingRules(client, logger, regions),
			compute.NewTargetPools(client, logger, regions),
			compute.NewInstances(client, logger, zones),
			compute.NewInstanceGroups(client, logger, zones),
			compute.NewHttpHealthChecks(client, logger),
			compute.NewHttpsHealthChecks(client, logger),
			compute.NewBackendServices(client, logger),
			compute.NewDisks(client, logger, zones),
			compute.NewNetworks(client, logger),
			compute.NewAddresses(client, logger, regions),
		},
	}, nil
}
