// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProjectContentEnvelope struct {
	Content ProjectModel `json:"content"`
}

type ProjectModel struct {
	ID         types.Int64  `tfsdk:"id" json:"id,computed"`
	Name       types.String `tfsdk:"name" json:"name,required"`
	RegionID   types.Int64  `tfsdk:"region_id" json:"region_id,required"`
	Note       types.String `tfsdk:"note" json:"note,optional"`
	AddressID  types.Int64  `tfsdk:"address_id" json:"address_id,computed"`
	Closed     types.Bool   `tfsdk:"closed" json:"closed,computed"`
	Created    types.String `tfsdk:"created" json:"created,computed"`
	ManagerID  types.Int64  `tfsdk:"manager_id" json:"manager_id,computed"`
	ResellerID types.Int64  `tfsdk:"reseller_id" json:"reseller_id,computed"`
	Updated    types.String `tfsdk:"updated" json:"updated,computed"`
	Uri        types.String `tfsdk:"uri" json:"uri,computed"`
}

func (m ProjectModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ProjectModel) MarshalJSONForUpdate(state ProjectModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
