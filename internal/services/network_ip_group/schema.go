// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*NetworkIPGroupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Network IP Goup record",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description: "The name to be given to the new IP Address Group. Used to identify the group when creating\nfirewall rules or geo-filters. Must start with a letter and contain only letters, numbers,\nunderscores, and hyphens.",
				Required:    true,
			},
			"cidrs": schema.ListAttribute{
				Description: "An array of CIDR addresses in the IP Address Group. Can include individual IP addresses\n(e.g., \"91.103.3.36\") or network ranges (e.g., \"90.103.2.0/24\"). All addresses must match\nthe specified IP version. Use these groups in firewall rules to allow/block traffic from\nmultiple locations with a single rule.",
				Required:    true,
				ElementType: types.StringType,
			},
			"version": schema.Int64Attribute{
				Description: "The IP version of the IP Address Group Objects in the IP Address Group. Accepted versions are 4 and 6.\nIf not sent, it will default to 4.",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network IP Group was created.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the Network IP Group",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network IP Group was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "The absolute URL of the Network IP Group record that can be used to perform `Read`, `Update` and `Delete`",
				Computed:    true,
			},
		},
	}
}

func (r *NetworkIPGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkIPGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
