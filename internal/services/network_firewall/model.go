// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkFirewallModel struct {
	ID        types.Int64                                           `tfsdk:"id" path:"id,optional"`
	ProjectID types.Int64                                           `tfsdk:"project_id" json:"project_id,required,no_refresh"`
	Type      types.String                                          `tfsdk:"type" json:"type,optional,no_refresh"`
	Name      types.String                                          `tfsdk:"name" json:"name,optional,no_refresh"`
	State     types.String                                          `tfsdk:"state" json:"state,optional,no_refresh"`
	Rules     *[]*NetworkFirewallRulesModel                         `tfsdk:"rules" json:"rules,optional,no_refresh"`
	Content   customfield.NestedObject[NetworkFirewallContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m NetworkFirewallModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkFirewallModel) MarshalJSONForUpdate(state NetworkFirewallModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NetworkFirewallRulesModel struct {
	Allow       types.Bool   `tfsdk:"allow" json:"allow,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Destination types.String `tfsdk:"destination" json:"destination,optional"`
	GroupName   types.String `tfsdk:"group_name" json:"group_name,optional"`
	Inbound     types.Bool   `tfsdk:"inbound" json:"inbound,optional"`
	Port        types.String `tfsdk:"port" json:"port,optional"`
	Protocol    types.String `tfsdk:"protocol" json:"protocol,optional"`
	Source      types.String `tfsdk:"source" json:"source,optional"`
}

type NetworkFirewallContentModel struct {
	ID        types.Int64                                                    `tfsdk:"id" json:"id,computed"`
	Created   types.String                                                   `tfsdk:"created" json:"created,computed"`
	Name      types.String                                                   `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                    `tfsdk:"project_id" json:"project_id,computed"`
	Rules     customfield.NestedObjectList[NetworkFirewallContentRulesModel] `tfsdk:"rules" json:"rules,computed"`
	Specs     customfield.NestedObjectList[NetworkFirewallContentSpecsModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                                    `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                   `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                   `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                   `tfsdk:"uri" json:"uri,computed"`
}

type NetworkFirewallContentRulesModel struct {
	Allow       types.Bool   `tfsdk:"allow" json:"allow,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Destination types.String `tfsdk:"destination" json:"destination,computed"`
	GroupName   types.String `tfsdk:"group_name" json:"group_name,computed"`
	Inbound     types.Bool   `tfsdk:"inbound" json:"inbound,computed"`
	Order       types.Int64  `tfsdk:"order" json:"order,computed"`
	Port        types.String `tfsdk:"port" json:"port,computed"`
	Protocol    types.String `tfsdk:"protocol" json:"protocol,computed"`
	Source      types.String `tfsdk:"source" json:"source,computed"`
	Version     types.Int64  `tfsdk:"version" json:"version,computed"`
}

type NetworkFirewallContentSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
