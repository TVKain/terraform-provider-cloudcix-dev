// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
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
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the Project which this Storage Volume should be in.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"instance_id": schema.Int64Attribute{
				Description:   "Required if type is \"hyperv\".\n\nThe ID of a Compute Instance with the type \"hyperv\" the Storage Volume is to be mounted to.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The user-friendly name for the Storage Volume type. If not sent and the type is \"cephfs\", it will default\nto the name 'Ceph'. If not sent and the type is \"hyperv\", it will default to the name 'Storage HyperV'.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Storage Volume to create. Valid options are:\n- \"cephfs\"\n  A Ceph file system volume which can be mounted to one or more Compute Instances of the type \"lxd\"\n- \"hyperv\"\n  A volume which can be mounted as a secondary drive to a Compute Instance of the type \"hyperv\"",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Storage Volume was created.",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "The current state of the Storage Volume",
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
						"state": schema.StringAttribute{
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
					"state": schema.StringAttribute{
						Description: `The current state of the "hyperv" Compute Instance the "hyperv" Storage Volume is attached to.`,
						Computed:    true,
					},
				},
			},
			"metadata": schema.SingleNestedAttribute{
				Description: `The metadata object of "ceph" Storage Volumes`,
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StorageVolumeMetadataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"attach_instance_ids": schema.ListAttribute{
						Description: `List of IDs of "lxd" Compute instances which the "ceph" Storage Volume should be attached to.`,
						Computed:    true,
						CustomType:  customfield.NewListType[types.Int64](ctx),
						ElementType: types.Int64Type,
					},
					"detach_instance_ids": schema.ListAttribute{
						Description: `List of IDs of "lxd" Compute instances which the "ceph" Storage Volume should be detached from.`,
						Computed:    true,
						CustomType:  customfield.NewListType[types.Int64](ctx),
						ElementType: types.Int64Type,
					},
					"mount_path": schema.StringAttribute{
						Description: `The mpunt path of the "ceph" Storage Volume on the "lxd" Compute instances.`,
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
