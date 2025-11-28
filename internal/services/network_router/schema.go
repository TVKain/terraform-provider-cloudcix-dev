// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*NetworkRouterResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Router Resource record",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the User's Project into which this Network Router should be added.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Network Router to create. Valid options are:\n- \"router\"\n A virtual route that manages IP forwarding, and participate in routing decisions\n for the Project.\n- \"static_route\"\n  Maps a destination network to a nexthop IP, enabling deterministic packet forwarding.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the Network Router. If not sent, it will default to current name.",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "Change the state of the Network Router, triggering the CloudCIX Robot to perform the requested action.\n\nAvailable state transitions:\n\nFrom running state, you can transition to:\n- update_running - Apply pending configuration changes while keeping the router operational\n- delete - Mark the router for deletion (requires all other project resources to be deleted first)\n\nFrom delete_queue state, you can transition to:\n- restart - Restore a router that was previously marked for deletion\n\nNote: To delete a router, all other resources in the project must first be in one of these states:\ndelete, delete_queue, or deleting.",
				Optional:    true,
			},
			"metadata": schema.SingleNestedAttribute{
				Description: `Metadata for the Netork Routers of the type "static_route".`,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"destination": schema.StringAttribute{
						Description: "CIDR notation of the destination address range of the target network of the static route.\nThe destination cannot be updated.",
						Optional:    true,
					},
					"nat": schema.BoolAttribute{
						Description: "Flag indicating if traffic from the destination can be routed to the Public internet via the\nProject's Router.",
						Optional:    true,
					},
					"nexthop": schema.StringAttribute{
						Description: "An IP address from one of the networks configured on the Router in the Project to forward the\npacket to. The nexthop cannot be updated.",
						Optional:    true,
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				Description: "Networks for the Netork Routers of the type \"router\".\n\nAn array of the list of networks defined on the Network Router.\nExisting networks (vlan property is not None) can have their names updated but IPv4/IPv6 ranges and VLAN\ncannot be modified. To create a new network on the Network Router, append an object to the list with an\n`ipv4` key for an available RFC 1918 address range. The `ipv6` and `vlan` values will be generated based\non what is available in the region.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4": schema.StringAttribute{
							Description: "The IPv4 address range of the network",
							Optional:    true,
						},
						"ipv6": schema.StringAttribute{
							Description: "The IPv6 address range of the network",
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the network",
							Optional:    true,
						},
						"vlan": schema.Int64Attribute{
							Description: "The VLAN ID of the network.",
							Optional:    true,
						},
					},
				},
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Router Resource record was created.",
				Computed:    true,
			},
			"grace_period": schema.Int64Attribute{
				Description: "Number of days after a user sets the state of the Router to Scrub (8) before it is executed by robot.\nThe default value is 7 days for a Router.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Router Resource record was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "URL that can be used to run methods in the API associated with the Network Routers instance.",
				Computed:    true,
			},
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Router Resource",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[NetworkRouterSpecsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"quantity": schema.Int64Attribute{
							Description: "How many units of a billable entity that a Resource utilises",
							Computed:    true,
						},
						"sku_name": schema.StringAttribute{
							Description: "An identifier for a billable entity that a Resource utilises",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *NetworkRouterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkRouterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
