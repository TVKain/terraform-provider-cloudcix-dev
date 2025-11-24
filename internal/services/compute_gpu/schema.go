// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ComputeGPUResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"pk": schema.Int64Attribute{
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the GPU Resource.",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "Change the state of the Compute GPU, triggering the CloudCIX Robot to perform the requested action.\n\nTo detach a Compute GPU from an lxd Compute Instance, send request to change the state to delete",
				Optional:    true,
			},
			"content": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ComputeGPUContentModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Compute GPU record",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Compute GPU record was created.",
						Computed:    true,
					},
					"instance": schema.SingleNestedAttribute{
						Description: `The "lxd" Compute Instance the Compute GPU is attached to.`,
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ComputeGPUContentInstanceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Description: `The ID of the "lxd" Compute Instance the Compute GPU is attached to.`,
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: `The user-friendly name of the "lxd" Compute Instance the Compute GPU is attached to.`,
								Computed:    true,
							},
							"state": schema.Int64Attribute{
								Description: `The current state of the "lxd" Compute Instance the Compute GPU is attached to.`,
								Computed:    true,
							},
						},
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name given to this Compute GPU",
						Computed:    true,
					},
					"project_id": schema.Int64Attribute{
						Description: "The id of the Project that this Compute GPU belongs to",
						Computed:    true,
					},
					"specs": schema.ListNestedAttribute{
						Description: "An array of the specs for the Compute GPU",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ComputeGPUContentSpecsModel](ctx),
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
						Description: "The current state of the Compute GPU",
						Computed:    true,
					},
					"updated": schema.StringAttribute{
						Description: "Timestamp, in ISO format, of when the Compute GPU record was last updated.",
						Computed:    true,
					},
					"uri": schema.StringAttribute{
						Description: "URL that can be used to run methods in the API associated with the Compute GPU instance.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *ComputeGPUResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ComputeGPUResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
