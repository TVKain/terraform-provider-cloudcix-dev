// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeInstanceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"content": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ComputeInstanceContentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Compute Instance record",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Compute Instance record was created.",
						Computed:    true,
					},
					"grace_period": schema.Int64Attribute{
						Description: "Number of days after a user sets the state of the Compute Instance Resource to Scrub (8) before it is\nexecuted by robot.",
						Computed:    true,
					},
					"interfaces": schema.ListNestedAttribute{
						Description: "Array of the interfaces for the Compute Instance",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ComputeInstanceContentInterfacesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"gateway": schema.BoolAttribute{
									Description: "Indicates if this interface functions as the network gateway for the Compute instance.",
									Computed:    true,
								},
								"ipv4_addresses": schema.ListNestedAttribute{
									Description: "An array of the IPv4 addresses on the Interface.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[ComputeInstanceContentInterfacesIpv4AddressesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.Int64Attribute{
												Description: "The ID of the IP Address record on the Interface.",
												Computed:    true,
											},
											"address": schema.StringAttribute{
												Description: "The IP address on the  Interface.",
												Computed:    true,
											},
											"nat": schema.BoolAttribute{
												Description: "Flag indicating if the IP Address is NATted to a Public IP Address",
												Computed:    true,
											},
											"public_ip": schema.StringAttribute{
												Description: "The Public IP address that the address is NATted to.",
												Computed:    true,
											},
										},
									},
								},
								"ipv6_addresses": schema.ListNestedAttribute{
									Description: "An array of the IPv6 addresses on the Interface.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[ComputeInstanceContentInterfacesIpv6AddressesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.Int64Attribute{
												Description: "The ID of the IP Address record on the Interface.",
												Computed:    true,
											},
											"address": schema.StringAttribute{
												Description: "The IP address on the Interface.",
												Computed:    true,
											},
										},
									},
								},
								"vlan": schema.Int64Attribute{
									Description: "The VLAN assigned to this Interface",
									Computed:    true,
								},
							},
						},
					},
					"metadata": schema.SingleNestedAttribute{
						Description: "The metadata details of the Compute Instance",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ComputeInstanceContentMetadataDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"dns": schema.StringAttribute{
								Description: `DNS server IP addresses for the Compute Instance.  Returned if the type is "lxd".`,
								Computed:    true,
							},
							"dns4": schema.StringAttribute{
								Description: `Primary DNS server IPv4 address for the VM.  Returned if the type is "hyperv".`,
								Computed:    true,
							},
							"dns6": schema.StringAttribute{
								Description: `Primary DNS server IPv6 address for the VM. Returned if the type is "hyperv".`,
								Computed:    true,
							},
							"instance_type": schema.StringAttribute{
								Description: `The name of the instance type. It will be either "container" or "vm". Returned if the type is "lxd".`,
								Computed:    true,
							},
							"userdata": schema.StringAttribute{
								Description: `Configuration file to be used by Cloud Init. Returned if the type is "lxd".`,
								Computed:    true,
							},
						},
					},
					"name": schema.StringAttribute{
						Description: "The human-friendly name given to this Compute Instance",
						Computed:    true,
					},
					"project_id": schema.Int64Attribute{
						Description: "The id of the Project that this Compute Instance belongs to",
						Computed:    true,
					},
					"specs": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ComputeInstanceContentSpecsDataSourceModel](ctx),
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
					"state": schema.Int64Attribute{
						Description: "The current state of the Compute Instance",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of the Compute Instance",
						Computed:    true,
					},
					"updated": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Compute Instance record was last updated.",
						Computed:    true,
					},
					"uri": schema.StringAttribute{
						Description: "URL that can be used to run methods in the API associated with the Compute Instance.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *ComputeInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeInstanceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
