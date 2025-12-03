// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*NetworkFirewallDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
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
			"state": schema.StringAttribute{
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
			"rules": schema.ListNestedAttribute{
				Description: "List of rules for this Network Firewall.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[NetworkFirewallRulesDataSourceModel](ctx),
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
						"zone": schema.StringAttribute{
							Description: "The zone in the firewall that the rule is applied to.",
							Computed:    true,
						},
					},
				},
			},
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Network Firewall",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[NetworkFirewallSpecsDataSourceModel](ctx),
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

func (d *NetworkFirewallDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkFirewallDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
