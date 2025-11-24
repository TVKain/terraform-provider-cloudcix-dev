// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_ip_group"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestNetworkIPGroupDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_ip_group.NetworkIPGroupDataSourceModel)(nil)
	schema := network_ip_group.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
