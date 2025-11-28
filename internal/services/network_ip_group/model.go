// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkIPGroupModel struct {
	ID      types.Int64                                          `tfsdk:"id" path:"id,optional"`
	Name    types.String                                         `tfsdk:"name" json:"name,required,no_refresh"`
	Cidrs   *[]types.String                                      `tfsdk:"cidrs" json:"cidrs,required,no_refresh"`
	Version types.Int64                                          `tfsdk:"version" json:"version,optional,no_refresh"`
	Content customfield.NestedObject[NetworkIPGroupContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m NetworkIPGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkIPGroupModel) MarshalJSONForUpdate(state NetworkIPGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NetworkIPGroupContentModel struct {
	ID      types.Int64                    `tfsdk:"id" json:"id,computed"`
	Cidrs   customfield.List[types.String] `tfsdk:"cidrs" json:"cidrs,computed"`
	Created types.String                   `tfsdk:"created" json:"created,computed"`
	Name    types.String                   `tfsdk:"name" json:"name,computed"`
	Type    types.String                   `tfsdk:"type" json:"type,computed"`
	Updated types.String                   `tfsdk:"updated" json:"updated,computed"`
	Uri     types.String                   `tfsdk:"uri" json:"uri,computed"`
	Version types.Int64                    `tfsdk:"version" json:"version,computed"`
}
