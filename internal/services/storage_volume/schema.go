// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*StorageVolumeResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Storage Volume record",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the Project which this Storage Volume should be in.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"instance_id": schema.Int64Attribute{
				Description:   "Required if type is \"hyperv\".\n\nThe ID of a Compute Instance with the type \"hyperv\" the Storage Volume is to be mounted to.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Storage Volume to create. Valid options are:\n- \"cephfs\"\n  A Ceph file system volume which can be mounted to one or more Compute Instances of the type \"lxd\"\n- \"hyperv\"\n  A volume which can be mounted as a secondary drive to a Compute Instance of the type \"hyperv\"",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"specs": schema.ListNestedAttribute{
				Description: "List of specs (SKUs) for the Storage Volume drive.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"quantity": schema.Int64Attribute{
							Description: "The quantity (GB) of the SKU to configure the Storage Volume drive with.",
							Optional:    true,
						},
						"sku_name": schema.StringAttribute{
							Description: "The name of the SKU for the Storage Volume drive.",
							Optional:    true,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the Storage Volume type. If not sent and the type is \"cephfs\", it will default\nto the name 'Ceph'. If not sent and the type is \"hyperv\", it will default to the name 'Storage HyperV'.",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "Change the state of the Storage Volume, triggering the CloudCIX Robot to perform the requested action.\nUsers can only request state changes from certain current states:\n\n- running -> update_running or delete",
				Optional:    true,
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "Required if type is \"cephfs\".\n\nMetadata for the Storage Volume drive with the type \"cephfs\".",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"attach_instance_ids": schema.ListAttribute{
						Description: "A list of IDs for running or stopped Compute Instances with the type \"lxd\" in the project to\nmount this Ceph file system volume to. If not sent, it will default to an empty list.",
						Optional:    true,
						ElementType: types.Int64Type,
					},
					"detach_instance_ids": schema.ListAttribute{
						Description: "A list of IDs for running or stopped Compute Instances with the type \"lxd\" in the project to\nunmount this Ceph file system volume from. If not sent, it will default to an empty list.",
						Optional:    true,
						ElementType: types.Int64Type,
					},
					"mount_path": schema.StringAttribute{
						Description: "The mount path for the Ceph file system volume inside the LXC instance.",
						Optional:    true,
					},
				},
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Storage Volume was created.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Storage Volume was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "URL that can be used to run methods in the API associated with the Storage Volumes instance.",
				Computed:    true,
			},
			"contra_instances": schema.ListNestedAttribute{
				Description: "A list of Compute Instances the Storage Volume is mounted to.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[StorageVolumeContraInstancesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "The ID of the Compute Instance the Storage Volume is attached to.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The user-friendly name given to the Compute Instance the Storage Volume is attached to.",
							Computed:    true,
						},
						"state": schema.Int64Attribute{
							Description: "The current state of the Compute Instance the Storage Volume is attached to.",
							Computed:    true,
						},
					},
				},
			},
			"instance": schema.SingleNestedAttribute{
				Description: `The "hyperv" Compute Instance the "hyperv" Storage Volume is attached to.`,
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StorageVolumeInstanceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: `The ID of the "hyperv" Compute Instance the "hyperv" Storage Volume is attached to.`,
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: `The user-friendly name of the "hyperv" Compute Instance the "hyperv" Storage Volume is attached to.`,
						Computed:    true,
					},
					"state": schema.Int64Attribute{
						Description: `The current state of the "hyperv" Compute Instance the "hyperv" Storage Volume is attached to.`,
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *StorageVolumeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *StorageVolumeResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
