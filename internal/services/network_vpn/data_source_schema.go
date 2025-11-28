// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_vpn

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*NetworkVpnDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"content": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[NetworkVpnContentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Network VPN record",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Network VPN record was created.",
						Computed:    true,
					},
					"metadata": schema.SingleNestedAttribute{
						Description: "The metadata for the configuration of the IKE and IPSec phases of the Network VPN.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[NetworkVpnContentMetadataDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"child_sas": schema.ListNestedAttribute{
								Description: "An array of Child SAs (Security Associations) of the routes for the Network VPN.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[NetworkVpnContentMetadataChildSasDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"local_ts": schema.StringAttribute{
											Description: "The CIDR notation of a subnet configured on the Network Router in the same Project",
											Computed:    true,
										},
										"remote_ts": schema.StringAttribute{
											Description: "CIDR notation of a subnet on the Customer side of the Network VPN.",
											Computed:    true,
										},
									},
								},
							},
							"ike_authentication": schema.StringAttribute{
								Description: "Authentication algorithms for the IKE phase.",
								Computed:    true,
							},
							"ike_dh_groups": schema.StringAttribute{
								Description: "Diffie-Helmen groups for the IKE phase.",
								Computed:    true,
							},
							"ike_encryption": schema.StringAttribute{
								Description: "Encryption algorithms for the IKE phase.",
								Computed:    true,
							},
							"ike_gateway_type": schema.StringAttribute{
								Description: "The gateway type (public_ip or hostname ) for the IKE phase.",
								Computed:    true,
							},
							"ike_gateway_value": schema.StringAttribute{
								Description: "The gateway value (IP or hostname) for the IKE phase.",
								Computed:    true,
							},
							"ike_lifetime": schema.Int64Attribute{
								Description: "The lifetime of the IKE phase in seconds.",
								Computed:    true,
							},
							"ike_local_identifier": schema.StringAttribute{
								Description: "The local identifier of the IKE phase in the region.",
								Computed:    true,
							},
							"ike_pre_shared_key": schema.StringAttribute{
								Description: "The pre shared key of the IKE phase.",
								Computed:    true,
							},
							"ike_remote_identifier": schema.StringAttribute{
								Description: "The remote identifier of the IKE phase on the customers side of the VPN.",
								Computed:    true,
							},
							"ike_version": schema.Int64Attribute{
								Description: "The version of the IKE phase",
								Computed:    true,
							},
							"ipsec_authentication": schema.StringAttribute{
								Description: "Authentication algorithms for the IPSec phase.",
								Computed:    true,
							},
							"ipsec_encryption": schema.StringAttribute{
								Description: "Encryption algorithms for the IPSec phase.",
								Computed:    true,
							},
							"ipsec_establish_time": schema.StringAttribute{
								Description: "The establish time for the IPSec phase.",
								Computed:    true,
							},
							"ipsec_lifetime": schema.Int64Attribute{
								Description: "The lifetime of the IPSec phase in seconds.",
								Computed:    true,
							},
							"ipsec_pfs_groups": schema.StringAttribute{
								Description: "Perfect Forward Secrecy groups for the IPSec phase.",
								Computed:    true,
							},
							"stif_number": schema.Int64Attribute{
								Description: "STIF number for the Network VPN.",
								Computed:    true,
							},
							"traffic_selector": schema.BoolAttribute{
								Description: "Flag for it traffic selectors are enabled.",
								Computed:    true,
							},
						},
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name given to this Network VPN instance",
						Computed:    true,
					},
					"project_id": schema.Int64Attribute{
						Description: "The id of the Project that this Network VPN belongs to",
						Computed:    true,
					},
					"specs": schema.ListNestedAttribute{
						Description: "An array of the specs for the Network VPN",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[NetworkVpnContentSpecsDataSourceModel](ctx),
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
						Description: "The current state of the Network VPN",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of the Network VPN",
						Computed:    true,
					},
					"updated": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Network VPN record was last updated.",
						Computed:    true,
					},
					"uri": schema.StringAttribute{
						Description: "URL that can be used to run methods in the API associated with the Network VPN instance.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *NetworkVpnDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkVpnDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
