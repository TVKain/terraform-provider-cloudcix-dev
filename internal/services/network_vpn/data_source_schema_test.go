// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_vpn_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_vpn"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestNetworkVpnDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_vpn.NetworkVpnDataSourceModel)(nil)
	schema := network_vpn.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
