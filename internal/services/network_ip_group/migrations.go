// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*NetworkIPGroupResource)(nil)

func (r *NetworkIPGroupResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
