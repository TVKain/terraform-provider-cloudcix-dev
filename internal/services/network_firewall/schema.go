// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*NetworkFirewallResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the Project which this Network Firewall should be created in. Each project can have\nexactly ONE project firewall and ONE geo firewall maximum.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Network Firewall to create. Each project can have exactly ONE of each type.\nValid options are:\n- \"geo\"\n  A Geofilter Firewall to allow or block traffic from/to specific countries using global\n  IP Address Groups (member_id = 0) that contain country IP ranges.\n- \"project\"\n  A Project Firewall with fine-grained rules for specific source/destination IPs, ports,\n  and protocols. Can reference your member's IP Address Groups using '@groupname' syntax.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the Network Firewall type. If not sent, it will default to current name.",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "Change the state of the Network Firewall, triggering the CloudCIX Robot to perform the requested action.\nUsers can only request state changes from certain current states:\n\n- running -> update_running or delete",
				Optional:    true,
			},
			"rules": schema.ListNestedAttribute{
				Description: "CRITICAL WARNING: This completely replaces ALL existing firewall rules. Any rules not included\nin this update will be permanently deleted. You must include the complete list of all rules\nyou want to keep, both existing and new ones.\n\nA list of the rules to be configured in the Network Firewall type. They will be applied in the\norder they are sent.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow": schema.BoolAttribute{
							Description: "Optional. A flag that states whether traffic matching this rule should be allowed to pass\nthrough the Network Firewall Type. If not sent, it wlll default to False.",
							Optional:    true,
						},
						"description": schema.StringAttribute{
							Description: "Optional description of what the rule is for.",
							Optional:    true,
						},
						"destination": schema.StringAttribute{
							Description: "Required if type is \"project\".\n\nA Subnet, IP Address or the name of a Project Network IP Group with `@` prepended\nthat indicates what the destination of traffic should be in order to match this rule.\n\nIt can also be just a `*` character, which will indicate that any destination is allowed.\n\nBoth source and destination must use the same IP Version.",
							Optional:    true,
						},
						"group_name": schema.StringAttribute{
							Description: "Required if type is \"geo\".\n\nThe name of a Geo Network IP Group with `@` prepended.",
							Optional:    true,
						},
						"inbound": schema.BoolAttribute{
							Description: "Optional. Flag indicating the direction of traffic this rule applies to:\n- true: Inbound rule (traffic coming INTO your project/network)\n- false: Outbound rule (traffic going OUT FROM your project/network)\n\nIf not sent, it will default to False (outbound rule).",
							Optional:    true,
						},
						"port": schema.StringAttribute{
							Description: "Required if type is \"project\".\n\nA string that indicates what the destination port of traffic should be in order to match this\nrule. The range for valid ports are between 1 - 65535 inclusive.\n\nAllowed Values:\n    `22`: Only port 22 is allowed\n    `20-25`: Ports between 20 and 25 inclusive are allowed\n    `22,24-25,444`: Combination of one or more single port and single port ranges\n                    with comma separated are allowed\n    ``: No port is required if the protocol is 'any' or 'icmp'",
							Optional:    true,
						},
						"protocol": schema.StringAttribute{
							Description: "Required if type is \"project\".\n\nA string that indicates what protocol traffic should be using in order to match this rule.\nThe supported protocols are;\n    - 'icmp'\n    - 'tcp'\n    - 'udp'\n\nThe special case protocol 'any' is allowed and allows any protocol through.",
							Optional:    true,
						},
						"source": schema.StringAttribute{
							Description: "Required if type is \"project\".\n\nA Subnet, IP Address or the name of a Project Network IP Group with `@` prepended\nthat indicates what the destination of traffic should be in order to match this rule.\n\nIt can also be just a `*` character, which will indicate that any source is allowed.\n\nBoth source and destination must use the same IP Version.",
							Optional:    true,
						},
					},
				},
			},
			"content": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[NetworkFirewallContentModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Network Firewall record",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Network Firewall record was created.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name given to this Network Firewall instance",
						Computed:    true,
					},
					"project_id": schema.Int64Attribute{
						Description: "The id of the Project that this Network Firewall belongs to",
						Computed:    true,
					},
					"rules": schema.ListNestedAttribute{
						Description: "List of rules for this Network Firewall.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[NetworkFirewallContentRulesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"allow": schema.BoolAttribute{
									Description: "True to allow traffic, False to deny.",
									Computed:    true,
								},
								"description": schema.StringAttribute{
									Description: `Optional description of the rule. Returned if the type is "project".`,
									Computed:    true,
								},
								"destination": schema.StringAttribute{
									Description: `Destination address or subnet. Use * for any. Returned if the type is "project".`,
									Computed:    true,
								},
								"group_name": schema.StringAttribute{
									Description: `The name of the Geo IP Address Group. Returned if the type is "geo".`,
									Computed:    true,
								},
								"inbound": schema.BoolAttribute{
									Description: "True if the rule applies to inbound traffic.",
									Computed:    true,
								},
								"order": schema.Int64Attribute{
									Description: `Order of rule evaluation (lower runs first). Returned if the type is "project".`,
									Computed:    true,
								},
								"port": schema.StringAttribute{
									Description: "Port or port range (e.g. 80, 443, 1000-2000). Not required for ICMP or ANY.\nReturned if the type is \"project\".",
									Computed:    true,
								},
								"protocol": schema.StringAttribute{
									Description: `Network protocol (any, icmp, tcp, udp). Returned if the type is "project".`,
									Computed:    true,
								},
								"source": schema.StringAttribute{
									Description: `Source address or subnet. Use * for any. Returned if the type is "project".`,
									Computed:    true,
								},
								"version": schema.Int64Attribute{
									Description: `IP version (4 or 6). Returned if the type is "project".`,
									Computed:    true,
								},
							},
						},
					},
					"specs": schema.ListNestedAttribute{
						Description: "An array of the specs for the Network Firewall",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[NetworkFirewallContentSpecsModel](ctx),
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
					"state": schema.Int64Attribute{
						Description: "The current state of the Network Firewall",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of the Network Firewall",
						Computed:    true,
					},
					"updated": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Network Firewall record was last updated.",
						Computed:    true,
					},
					"uri": schema.StringAttribute{
						Description: "URL that can be used to run methods in the API associated with the Network Firewall instance.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *NetworkFirewallResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkFirewallResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
