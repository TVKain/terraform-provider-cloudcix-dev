// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ProjectResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Project.",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"region_id": schema.Int64Attribute{
				Description:   "The Address ID of the CloudCIX region that the Project will be deployed in.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the Project. Must be unique within an Address' Project collection.",
				Required:    true,
			},
			"note": schema.StringAttribute{
				Description: "An optional note providing a description of what the Project is used for.",
				Optional:    true,
			},
			"address_id": schema.Int64Attribute{
				Description: "The ID of the Project address.",
				Computed:    true,
			},
			"closed": schema.BoolAttribute{
				Description: "A flag stating whether or not the Project is classified as closed. A Project is classified as closed when\nall the infrastructure in it is in a Closed (99) state.",
				Computed:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date that the Project entry was created",
				Computed:    true,
			},
			"manager_id": schema.Int64Attribute{
				Description: "The ID of the User that manages the Project",
				Computed:    true,
			},
			"reseller_id": schema.Int64Attribute{
				Description: "The Address ID that will send the bill for the Project to the customer.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "The date that the Project entry was last updated",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "The absolute URL of the Project that can be used to perform `Read`, `Update` and `Delete`",
				Computed:    true,
			},
		},
	}
}

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ProjectResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
