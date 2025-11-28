// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProjectContentDataSourceEnvelope struct {
	Content ProjectDataSourceModel `json:"content,computed"`
}

type ProjectDataSourceModel struct {
	ID         types.Int64  `tfsdk:"id" path:"id,required"`
	AddressID  types.Int64  `tfsdk:"address_id" json:"address_id,computed"`
	Closed     types.Bool   `tfsdk:"closed" json:"closed,computed"`
	Created    types.String `tfsdk:"created" json:"created,computed"`
	ManagerID  types.Int64  `tfsdk:"manager_id" json:"manager_id,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
	Note       types.String `tfsdk:"note" json:"note,computed"`
	RegionID   types.Int64  `tfsdk:"region_id" json:"region_id,computed"`
	ResellerID types.Int64  `tfsdk:"reseller_id" json:"reseller_id,computed"`
	Updated    types.String `tfsdk:"updated" json:"updated,computed"`
	Uri        types.String `tfsdk:"uri" json:"uri,computed"`
}
