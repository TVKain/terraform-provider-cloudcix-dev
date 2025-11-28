// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkFirewallDataSourceModel struct {
	ID      types.Int64                                                     `tfsdk:"id" path:"id,required"`
	Content customfield.NestedObject[NetworkFirewallContentDataSourceModel] `tfsdk:"content" json:"content,computed"`
}

type NetworkFirewallContentDataSourceModel struct {
	ID        types.Int64                                                              `tfsdk:"id" json:"id,computed"`
	Created   types.String                                                             `tfsdk:"created" json:"created,computed"`
	Name      types.String                                                             `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                              `tfsdk:"project_id" json:"project_id,computed"`
	Rules     customfield.NestedObjectList[NetworkFirewallContentRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
	Specs     customfield.NestedObjectList[NetworkFirewallContentSpecsDataSourceModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                                              `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                             `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                             `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                             `tfsdk:"uri" json:"uri,computed"`
}

type NetworkFirewallContentRulesDataSourceModel struct {
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

type NetworkFirewallContentSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
