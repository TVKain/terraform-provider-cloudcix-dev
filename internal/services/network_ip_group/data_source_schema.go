// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"context"

	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*NetworkIPGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network IP Group was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Network IP Group.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the Network IP Group",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network IP Group was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "The absolute URL of the Network IP Group record that can be used to perform `Read`, `Update` and `Delete`",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The IP Version of the CIDRs in the group.",
				Computed:    true,
			},
			"cidrs": schema.ListAttribute{
				Description: "An array of CIDR addresses in the Network IP Group.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *NetworkIPGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkIPGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
