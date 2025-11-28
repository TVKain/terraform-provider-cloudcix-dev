// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ComputeInstanceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the Project which this Compute Intsance Resource should be in.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Compute Instance to create. Valid options are:\n- \"hyperv\"\n- \"lxd\"",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "Optional. The metadata required to configure in an Compute Instance.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"dns": schema.StringAttribute{
						Description: "Optional. A string containing IP Addresses, separated by commas, that represent the DNS servers\nthat the Compute Instance will use.",
						Optional:    true,
					},
					"instance_type": schema.StringAttribute{
						Description: "Optional, The Compute Instance instance type of the VM. Valid options are \"container\"\nor \"virtual-machine\". If not sent it will default to \"container\".",
						Optional:    true,
					},
					"userdata": schema.StringAttribute{
						Description: "Cloud Init allows Mime Multi-part messages, or files that start with a given set of strings. It is\na requirement to configure at minimum one user with a password or ssh key that is in the sudo group.\n\nReference: https://cloudinit.readthedocs.io/en/latest/explanation/format.html\n\nA hashed password can be generated using `openssl passwd -6 'yourpassword'`\n\nExample:\n```yaml\n#cloud-config\nusers:\n  - name: administrator\n    groups: sudo\n    shell: /bin/bash\n    lock_passwd: false\n    passwd: < HASHED PASWWORD >\n    ssh_authorized_keys:\n      - < HASHED PASWWORD >\nchpasswd:\n  expire: false\nssh_pwauth: true\n```",
						Optional:    true,
					},
				},
			},
			"specs": schema.ListNestedAttribute{
				Description: "List of specs (SKUs) for the Compute Instance resource.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"quantity": schema.Int64Attribute{
							Description: "The quantity of the SKU to configure the Compute Instance instance with.",
							Optional:    true,
						},
						"sku_name": schema.StringAttribute{
							Description: "The name of the SKU to configure on the Compute Instance instacne",
							Optional:    true,
						},
					},
				},
			},
			"grace_period": schema.Int64Attribute{
				Description: "The number of days after a Compute Intsance is closed before it is permanently deleted.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the Compute Intsance type. If not sent and the type is \"lxd\", it will default\nto the name 'LXD'. If not sent and the type is \"hyperv\", it will default to the name 'VM HyperV'.",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "Change the state of the Compute Instance, triggering the CloudCIX Robot to perform the requested action.\nUsers can only request state changes from certain current states, with specific allowed target states:\n\n- running        -> stop, delete, or update_running\n- stopped        -> restart, delete, or update_stopped\n- delete_queue   -> restart or stop",
				Optional:    true,
			},
			"interfaces": schema.ListNestedAttribute{
				Description: "Optional. A list of network interfaces that represent the interfaces that will be configured on the LXD\ninstance.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"gateway": schema.BoolAttribute{
							Description: "Flag representing if this interface will be the Gateway Interface to the Public\nInternet.",
							Optional:    true,
						},
						"ipv4_addresses": schema.ListNestedAttribute{
							Description: "A list of IPv4 address objects to be assigned to this interface. All addresses in this list\nmust be from the same network.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description: "An RFC1918 IPv4 address to be configured on this interface on\nthe Compute Instance instance.",
										Optional:    true,
									},
									"nat": schema.BoolAttribute{
										Description: "Optional, Flag indicating if this address should be NATted to a Public IP Address.\nIf not sent, it will default to False.",
										Optional:    true,
									},
								},
							},
						},
						"ipv6_addresses": schema.ListNestedAttribute{
							Description: "A list of IPv6 address objects to be assigned to this interface. All addresses in this list\nmust be from the same network as the `ipv4_addresses`.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description: "An IPv6 address to be configured on this interface on\nthe Compute Instance instance.",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"content": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ComputeInstanceContentModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[ComputeInstanceContentInterfacesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"gateway": schema.BoolAttribute{
									Description: "Indicates if this interface functions as the network gateway for the Compute instance.",
									Computed:    true,
								},
								"ipv4_addresses": schema.ListNestedAttribute{
									Description: "An array of the IPv4 addresses on the Interface.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[ComputeInstanceContentInterfacesIpv4AddressesModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[ComputeInstanceContentInterfacesIpv6AddressesModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[ComputeInstanceContentMetadataModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[ComputeInstanceContentSpecsModel](ctx),
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

func (r *ComputeInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ComputeInstanceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
