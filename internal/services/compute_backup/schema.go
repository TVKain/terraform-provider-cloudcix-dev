// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ComputeBackupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Compute Backups record",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"instance_id": schema.Int64Attribute{
				Description:   "The id of the Compute Instance the Compute Backup is to be taken of.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the User's Project into which this new Compute Backups should be added.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The user-friendly name for the Compute Backup. If not sent, it will default to the name\n\"Backup HyperV\" or \"Backup LXD\" depending on the type chosen.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Compute Backup to create. Valid options are:\n- \"hyperv\"\n- \"lxd\"",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Compute Backups record was created.",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "The current state of the Compute Backups",
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
			"instance": schema.SingleNestedAttribute{
				Description: "The Compute Instance the Compute Backup record is of.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeBackupInstanceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Compute Instance the Compute Backup is of.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name of the Compute Instance the Compute Backup is of.",
						Computed:    true,
					},
					"state": schema.StringAttribute{
						Description: "The current state of the Compute Instance the Compute Backup is of.",
						Computed:    true,
					},
				},
			},
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Compute Backups",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ComputeBackupSpecsModel](ctx),
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

func (r *ComputeBackupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ComputeBackupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
