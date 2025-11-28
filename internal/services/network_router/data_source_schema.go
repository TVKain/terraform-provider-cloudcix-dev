// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*NetworkRouterDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"content": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[NetworkRouterContentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Router Resource record",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Router Resource record was created.",
						Computed:    true,
					},
					"grace_period": schema.Int64Attribute{
						Description: "Number of days after a user sets the state of the Router to Scrub (8) before it is executed by robot.\nThe default value is 7 days for a Router.",
						Computed:    true,
					},
					"metadata": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[NetworkRouterContentMetadataDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ipv4_address": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[NetworkRouterContentMetadataIpv4AddressDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										Description: "The ID of the IPAddress record.",
										Computed:    true,
									},
									"address": schema.StringAttribute{
										Description: "The IP address of the IPAddress record.",
										Computed:    true,
									},
									"created": schema.StringAttribute{
										Description: "Timestamp, in ISO format, of when the IPAddress record was created.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "A verbose name given to the IPAddress record.",
										Computed:    true,
									},
									"notes": schema.StringAttribute{
										Description: "The note attached to IPAddress that made it.",
										Computed:    true,
									},
									"public_ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[NetworkRouterContentMetadataIpv4AddressPublicIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.Int64Attribute{
												Description: "The ID of the Public IPAddress record.",
												Computed:    true,
											},
											"address": schema.StringAttribute{
												Description: "The actual address of the Public IPAddress record.",
												Computed:    true,
											},
										},
									},
									"public_ip_id": schema.Int64Attribute{
										Description: "The ID of the Public IPAddress record.",
										Computed:    true,
									},
									"subnet_id": schema.Int64Attribute{
										Description: "The ID of the Subnet record.",
										Computed:    true,
									},
									"updated": schema.StringAttribute{
										Description: "Timestamp, in ISO format, of when the IPAddress record was last updated.",
										Computed:    true,
									},
								},
							},
							"ipv4_address_id": schema.Int64Attribute{
								Description: "The ID of the assigned public IPv4 address for the Router.",
								Computed:    true,
							},
							"ipv6_address": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[NetworkRouterContentMetadataIpv6AddressDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										Description: "The ID of the IPAddress record.",
										Computed:    true,
									},
									"address": schema.StringAttribute{
										Description: "The IP address of the IPAddress record.",
										Computed:    true,
									},
									"created": schema.StringAttribute{
										Description: "Timestamp, in ISO format, of when the IPAddress record was created.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "A verbose name given to the IPAddress record.",
										Computed:    true,
									},
									"notes": schema.StringAttribute{
										Description: "The note attached to IPAddress that made it.",
										Computed:    true,
									},
									"public_ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[NetworkRouterContentMetadataIpv6AddressPublicIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.Int64Attribute{
												Description: "The ID of the Public IPAddress record.",
												Computed:    true,
											},
											"address": schema.StringAttribute{
												Description: "The actual address of the Public IPAddress record.",
												Computed:    true,
											},
										},
									},
									"public_ip_id": schema.Int64Attribute{
										Description: "The ID of the Public IPAddress record.",
										Computed:    true,
									},
									"subnet_id": schema.Int64Attribute{
										Description: "The ID of the Subnet record.",
										Computed:    true,
									},
									"updated": schema.StringAttribute{
										Description: "Timestamp, in ISO format, of when the IPAddress record was last updated.",
										Computed:    true,
									},
								},
							},
							"ipv6_address_id": schema.Int64Attribute{
								Description: "The ID of the assigned public IPv6 address for the Router.",
								Computed:    true,
							},
						},
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name given to this Router Resource instance",
						Computed:    true,
					},
					"networks": schema.ListNestedAttribute{
						Description: "An array of the list of networks defined on the Router",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[NetworkRouterContentNetworksDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"destination": schema.StringAttribute{
									Description: "The destination address range of the target network of the static route. Returned if the\ntype is \"static_route\".",
									Computed:    true,
								},
								"ipv4": schema.StringAttribute{
									Description: `The IPv4 address range of the network. Returned if the type is "router".`,
									Computed:    true,
								},
								"ipv6": schema.StringAttribute{
									Description: `The IPv6 address range of the network. Returned if the type is "router".`,
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: `The name of the network. Returned if the type is "router".`,
									Computed:    true,
								},
								"nat": schema.BoolAttribute{
									Description: "Flag indicating if traffic from the destination can route to the Public internet. Returned if\nthe type is \"static_route\".",
									Computed:    true,
								},
								"nexthop": schema.StringAttribute{
									Description: "An IP address from one of the networks configured on the Router in the Project to forward the\npacket to. Returned if the type is \"static_route\".",
									Computed:    true,
								},
								"vlan": schema.Int64Attribute{
									Description: `The VLAN ID of the network. Returned if the type is "router".`,
									Computed:    true,
								},
							},
						},
					},
					"project_id": schema.Int64Attribute{
						Description: "The id of the Project that this Router Resource belongs to",
						Computed:    true,
					},
					"specs": schema.ListNestedAttribute{
						Description: "An array of the specs for the Router Resource",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[NetworkRouterContentSpecsDataSourceModel](ctx),
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
						Description: "The current state of the Router Resource",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of the Network Router",
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
				},
			},
		},
	}
}

func (d *NetworkRouterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkRouterDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
