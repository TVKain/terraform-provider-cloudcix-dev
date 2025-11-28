// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeBackupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"content": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ComputeBackupContentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Compute Backups record",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Compute Backups record was created.",
						Computed:    true,
					},
					"instance": schema.SingleNestedAttribute{
						Description: "The Compute Instance the Compute Backup record is of.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ComputeBackupContentInstanceDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Description: "The ID of the Compute Instance the Compute Backup is of.",
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: "The user-friendly name of the Compute Instance the Compute Backup is of.",
								Computed:    true,
							},
							"state": schema.Int64Attribute{
								Description: "The current state of the Compute Instance the Compute Backup is of.",
								Computed:    true,
							},
						},
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name given to this Compute Backups instance",
						Computed:    true,
					},
					"project_id": schema.Int64Attribute{
						Description: "The id of the Project that this Compute Backups belongs to",
						Computed:    true,
					},
					"specs": schema.ListNestedAttribute{
						Description: "An array of the specs for the Compute Backups",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ComputeBackupContentSpecsDataSourceModel](ctx),
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
						Description: "The current state of the Compute Backups",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of the Compute Backups",
						Computed:    true,
					},
					"updated": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Compute Backups record was last updated.",
						Computed:    true,
					},
					"uri": schema.StringAttribute{
						Description: "URL that can be used to run methods in the API associated with the Compute Backups instance.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *ComputeBackupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeBackupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
