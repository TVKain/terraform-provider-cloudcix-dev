// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_snapshot

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ComputeSnapshotResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Compute Snapshots record",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"instance_id": schema.Int64Attribute{
				Description:   "The id of the Compute Instance the Compute Snapshot is to be taken of.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"project_id": schema.Int64Attribute{
				Description:   "The ID of the User's Project into which this new Compute Snapshots should be added.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description:   "The type of Compute Snapshot to create. Valid options are:\n- \"hyperv\"\n- \"lxd\"",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the Compute Snapshot Resource. If not sent, it will default to the name\n\"Compute Snapshot\"",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "Change the state of the Compute Snapshot, triggering the CloudCIX Robot to perform the requested action.\nUsers can only request state changes from certain current states:\n\n- running -> update_running or delete",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Compute Snapshots record was created.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Compute Snapshots record was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "URL that can be used to run methods in the API associated with the Compute Snapshots instance.",
				Computed:    true,
			},
			"instance": schema.SingleNestedAttribute{
				Description: "The Compute Instance the Compute Snapshot record is of.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeSnapshotInstanceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Compute Instance the Compute Snapshot is of.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name of the Compute Instance the Compute Snapshot is of.",
						Computed:    true,
					},
					"state": schema.Int64Attribute{
						Description: "The current state of the Compute Instance the Compute Snapshot is of.",
						Computed:    true,
					},
				},
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "The metadata details of the The metadata details of the \"hyperv\" Compute Snapshot. Returned if the type\nis \"hyperv\".",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeSnapshotMetadataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Description: `Indicates if the "hyperv" Compute Snapshot is currently active.`,
						Computed:    true,
					},
					"remove_subtree": schema.BoolAttribute{
						Description: `Indicates if the "hyperv" Compute Snapshot should remove the subtree when deleted.`,
						Computed:    true,
					},
				},
			},
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Compute Snapshots",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ComputeSnapshotSpecsModel](ctx),
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

func (r *ComputeSnapshotResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ComputeSnapshotResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
