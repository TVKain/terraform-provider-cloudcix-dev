// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_vpn

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*NetworkVpnResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Network VPN record",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the User's Project into which this Network VPN should be added.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Network VPN to create. Valid options are:\n- \"site-to-site\"",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the Network VPN. If not sent, it will default to the name 'VPNS2S'",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "Change the state of the Network VPN, triggering the CloudCIX Robot to perform the requested action.\nUsers can only request state changes from certain current states:\n\n- running -> update_running or delete",
				Optional:    true,
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "Optional. The metadata required to configure the Network VPN instance.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"child_sas": schema.ListNestedAttribute{
						Description: "An array of Child SAs (Security Associations) to create the initial routes for the VPN.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"local_ts": schema.StringAttribute{
									Description: "The local Subnet in Network VPN's Project that will be configured as part of this route\nin a Child SA.",
									Optional:    true,
								},
								"remote_ts": schema.StringAttribute{
									Description: "CIDR notation of the Remote Subnet on the Customer side of the Network VPN that should\nbe given access through the VPN.\n\nNote:\n- The sent address range can overlap with the subnets configured on the Router in the\nProject for the VPNS2S.\n- The sent address range cannot overlap with a remote subnet of another VPN in the\nsame Project.",
									Optional:    true,
								},
							},
						},
					},
					"ike_authentication": schema.StringAttribute{
						Description: "Optional. A string containing a comma separated array of authentication algorithms for the IKE\nphase of the Network VPN. If not sent, it will default to `SHA384`.\n\nThe IKE phase authentication algorithms supported are;\n- `SHA1`\n- `SHA256`\n- `SHA384`\n\nPlease ensure that each entry in the array matches one of the above strings exactly.\nDuplicate entries will be ignored.",
						Optional:    true,
					},
					"ike_dh_groups": schema.StringAttribute{
						Description: "Optional. A string containing a comma separated array of Diffie-Helmen groups for the IKE phase\nof the Network VPN. If not sent, it will default to `Group 24`.\n\nThe IKE phase Diffie-Helmen groups supported are;\n- `Group 1`\n- `Group 2`\n- `Group 5`\n- `Group 19`\n- `Group 20`\n- `Group 24`\n\nPlease ensure that each entry in the array matches one of the above strings exactly.\nDuplicate entries will be ignored.",
						Optional:    true,
					},
					"ike_encryption": schema.StringAttribute{
						Description: "Optional. A string containing a comma separated array of encryption algorithms for the IKE phase of\nthe Network VPN. If not sent, it will default to \"256 bit AES-CBC\".\n\nThe IKE phase encryption algorithms supported are;\n- `128 bit AES-CBC`\n- `192 bit AES-CBC`\n- `256 bit AES-CBC`\n\nPlease ensure that each entry in the array matches one of the above strings exactly.\nDuplicate entries will be ignored.",
						Optional:    true,
					},
					"ike_gateway_type": schema.StringAttribute{
						Description: "The type of data that is stored in the `ike_gateway_value` field. It can only either \"public_ip\" or\n\"hostname\". If not sent, it will default to \"public_ip\".",
						Optional:    true,
					},
					"ike_gateway_value": schema.StringAttribute{
						Description: "The value used as the IKE gateway for the Network VPN. The type for this value depends on what type\nwas sent for the \"ike_gateway_type\".\n\nFor \"public_ip\", this value must be a string containing an IP address.\nFor \"hostname\", this value must be a valid hostname.",
						Optional:    true,
					},
					"ike_lifetime": schema.Int64Attribute{
						Description: "Optional. The lifetime of the IKE phase in seconds. Must be a value between 180 and 86400 inclusive.\nIf not sent, it will default to 28800.",
						Optional:    true,
					},
					"ike_pre_shared_key": schema.StringAttribute{
						Description: "The pre shared key to use for setting up the IKE phase of the Network VPN.\n\nNote that the pre shared key cannot contain any of the following special characters;\n- `\"`\n- `'`\n- `@`\n- `+`\n- `-`\n- `/`\n- `\\`\n- `|`\n- `=`",
						Optional:    true,
					},
					"ike_version": schema.StringAttribute{
						Description: "Optional. String value of the chosen version for the IKE phase. If not sent, it will default to\n\"v2-only\".\n\nThe IKE phase versions supported are;\n- `v1-only`\n- `v2-only`\n\nPlease ensure the sent string matches one of these exactly.",
						Optional:    true,
					},
					"ipsec_authentication": schema.StringAttribute{
						Description: "Optional. A string containing a comma separated array of authentication algorithms for the IPSec\nphase of the Site-to-Site Network VPN. If not sent, it will default to \"SHA256\".\n\nThe IPSec phase authentication algorithms supported are;\n- `SHA1`\n- `SHA256`\n\nPlease ensure that each entry in the array matches one of the above strings exactly.\nDuplicate entries will be ignored.",
						Optional:    true,
					},
					"ipsec_encryption": schema.StringAttribute{
						Description: "Optional. A string containing a comma separated array of encryption algorithms for the IPSEC phase\nof the Network VPN. If not sent, it will default to \"AES 256\".\n\nThe IPSEC phase encryption algorithms supported are;\n- `AES 128`\n- `AES 192`\n- `AES 256`\n- `AES 128 GCM`\n- `AES 192 GCM`\n- `AES 256 GCM`\n\nPlease ensure that each entry in the array matches one of the above strings exactly.\nDuplicate entries will be ignored.",
						Optional:    true,
					},
					"ipsec_establish_time": schema.StringAttribute{
						Description: "Optional. String value of the chosen establish_time for the IPSec phase. If not sent, it will\ndefault to \"Immediately\".\n\nThe IPSec phase establish time values supported are;\n- `Immediately`\n- `On Traffic`\n\nPlease ensure the sent string matches one of these exactly.",
						Optional:    true,
					},
					"ipsec_lifetime": schema.Int64Attribute{
						Description: "Optional. The lifetime of the IPSec phase in seconds. It be a value between 180 and 86400 inclusive.\nIf not sent, it will default to 3600.",
						Optional:    true,
					},
					"ipsec_pfs_groups": schema.StringAttribute{
						Description: "Optional. A string containing a comma separated array of Perfect Forward Secrecy (PFS) groups for\nthe IPSec phase of the Network VPN. If not sent, it will default to \"Group 20\".\n\nThe IPSec phase PFS groups supported are;\n- `Group 1`\n- `Group 2`\n- `Group 5`\n- `Group 14`\n- `Group 19`\n- `Group 20`\n- `Group 24`\n\nPlease ensure that each entry in the array matches one of the above strings exactly.\nDuplicate entries will be ignored.",
						Optional:    true,
					},
					"traffic_selector": schema.BoolAttribute{
						Description: "Optional. Boolean value stating if traffic selectors are to be used in configuring vpn tunnel.\nIf not sent, it will default to false.\n\nBy default, 0.0.0.0/0 will be used for the default local and remote CIDRs.\nIf true, then each of the local and remote CIDRs will be added to the configuration negotiation\nwith peer.",
						Optional:    true,
					},
				},
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network VPN record was created.",
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
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Network VPN",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[NetworkVpnSpecsModel](ctx),
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

func (r *NetworkVpnResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkVpnResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
