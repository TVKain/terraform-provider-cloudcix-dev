// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_router"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestNetworkRouterDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_router.NetworkRouterDataSourceModel)(nil)
	schema := network_router.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
