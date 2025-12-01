// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkRouterContentEnvelope struct {
	Content NetworkRouterModel `json:"content"`
}

type NetworkRouterModel struct {
	ID          types.Int64                                              `tfsdk:"id" json:"id,computed"`
	ProjectID   types.Int64                                              `tfsdk:"project_id" json:"project_id,required"`
	Type        types.String                                             `tfsdk:"type" json:"type,optional"`
	Name        types.String                                             `tfsdk:"name" json:"name,optional"`
	State       types.String                                             `tfsdk:"state" json:"state,optional,no_refresh"`
	Metadata    *NetworkRouterMetadataModel                              `tfsdk:"metadata" json:"metadata,optional"`
	Networks    customfield.NestedObjectList[NetworkRouterNetworksModel] `tfsdk:"networks" json:"networks,computed_optional"`
	Created     types.String                                             `tfsdk:"created" json:"created,computed"`
	GracePeriod types.Int64                                              `tfsdk:"grace_period" json:"grace_period,computed"`
	Updated     types.String                                             `tfsdk:"updated" json:"updated,computed"`
	Uri         types.String                                             `tfsdk:"uri" json:"uri,computed"`
	Specs       customfield.NestedObjectList[NetworkRouterSpecsModel]    `tfsdk:"specs" json:"specs,computed"`
}

func (m NetworkRouterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkRouterModel) MarshalJSONForUpdate(state NetworkRouterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NetworkRouterMetadataModel struct {
	Destination types.String `tfsdk:"destination" json:"destination,optional,no_refresh"`
	Nat         types.Bool   `tfsdk:"nat" json:"nat,optional,no_refresh"`
	Nexthop     types.String `tfsdk:"nexthop" json:"nexthop,optional,no_refresh"`
}

type NetworkRouterNetworksModel struct {
	Ipv4 types.String `tfsdk:"ipv4" json:"ipv4,optional"`
	Ipv6 types.String `tfsdk:"ipv6" json:"ipv6,computed"`
	Name types.String `tfsdk:"name" json:"name,optional"`
	Vlan types.Int64  `tfsdk:"vlan" json:"vlan,computed"`
}

type NetworkRouterSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
