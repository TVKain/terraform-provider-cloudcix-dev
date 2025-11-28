// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*StorageVolumeDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Storage Volume was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name given to this Storage Volume",
				Computed:    true,
			},
			"project_id": schema.Int64Attribute{
				Description: "The ID of the Project that this Storage Volume belongs to",
				Computed:    true,
			},
			"state": schema.Int64Attribute{
				Description: "The current state of the Storage Volume",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the Storage Volume",
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
				CustomType:  customfield.NewNestedObjectListType[StorageVolumeContraInstancesDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[StorageVolumeInstanceDataSourceModel](ctx),
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
			"metadata": schema.SingleNestedAttribute{
				Description: `The metadata object of "ceph" Storage Volumes`,
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[StorageVolumeMetadataDataSourceModel](ctx),
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
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Storage Volume",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[StorageVolumeSpecsDataSourceModel](ctx),
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

func (d *StorageVolumeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StorageVolumeDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
