package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type FirewallsClient struct {
	ListFirewallsCall struct {
		CallCount int
		Receives  struct {
			Filter string
		}
		Returns struct {
			Output *gcpcompute.FirewallList
			Error  error
		}
	}

	DeleteFirewallCall struct {
		CallCount int
		Receives  struct {
			Firewall string
		}
		Returns struct {
			Error error
		}
	}
}

func (c *FirewallsClient) ListFirewalls(filter string) (*gcpcompute.FirewallList, error) {
	c.ListFirewallsCall.CallCount++
	c.ListFirewallsCall.Receives.Filter = filter

	return c.ListFirewallsCall.Returns.Output, c.ListFirewallsCall.Returns.Error
}

func (c *FirewallsClient) DeleteFirewall(firewall string) error {
	c.DeleteFirewallCall.CallCount++
	c.DeleteFirewallCall.Receives.Firewall = firewall

	return c.DeleteFirewallCall.Returns.Error
}
