// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkIPGroupContentEnvelope struct {
	Content NetworkIPGroupModel `json:"content"`
}

type NetworkIPGroupModel struct {
	ID      types.Int64     `tfsdk:"id" json:"id,computed"`
	Name    types.String    `tfsdk:"name" json:"name,required"`
	Cidrs   *[]types.String `tfsdk:"cidrs" json:"cidrs,required"`
	Version types.Int64     `tfsdk:"version" json:"version,optional"`
	Created types.String    `tfsdk:"created" json:"created,computed"`
	Type    types.String    `tfsdk:"type" json:"type,computed"`
	Updated types.String    `tfsdk:"updated" json:"updated,computed"`
	Uri     types.String    `tfsdk:"uri" json:"uri,computed"`
}

func (m NetworkIPGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkIPGroupModel) MarshalJSONForUpdate(state NetworkIPGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
