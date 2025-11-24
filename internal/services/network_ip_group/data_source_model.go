// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkIPGroupDataSourceModel struct {
	Pk      types.Int64                                                    `tfsdk:"pk" path:"pk,required"`
	Content customfield.NestedObject[NetworkIPGroupContentDataSourceModel] `tfsdk:"content" json:"content,computed"`
}

type NetworkIPGroupContentDataSourceModel struct {
	ID      types.Int64                    `tfsdk:"id" json:"id,computed"`
	Cidrs   customfield.List[types.String] `tfsdk:"cidrs" json:"cidrs,computed"`
	Created types.String                   `tfsdk:"created" json:"created,computed"`
	Name    types.String                   `tfsdk:"name" json:"name,computed"`
	Type    types.String                   `tfsdk:"type" json:"type,computed"`
	Updated types.String                   `tfsdk:"updated" json:"updated,computed"`
	Uri     types.String                   `tfsdk:"uri" json:"uri,computed"`
	Version types.Int64                    `tfsdk:"version" json:"version,computed"`
}
