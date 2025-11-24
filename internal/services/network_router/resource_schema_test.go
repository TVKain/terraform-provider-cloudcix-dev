// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router_test

import (
	"context"
	"testing"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/services/network_router"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/test_helpers"
)

func TestNetworkRouterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_router.NetworkRouterModel)(nil)
	schema := network_router.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
