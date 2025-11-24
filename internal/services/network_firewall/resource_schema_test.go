// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_firewall"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestNetworkFirewallModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_firewall.NetworkFirewallModel)(nil)
	schema := network_firewall.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
