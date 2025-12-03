// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkFirewallContentEnvelope struct {
	Content NetworkFirewallModel `json:"content"`
}

type NetworkFirewallModel struct {
	ID        types.Int64                                             `tfsdk:"id" json:"id,computed"`
	ProjectID types.Int64                                             `tfsdk:"project_id" json:"project_id,required"`
	Name      types.String                                            `tfsdk:"name" json:"name,optional"`
	Type      types.String                                            `tfsdk:"type" json:"type,optional"`
	Rules     *[]*NetworkFirewallRulesModel                           `tfsdk:"rules" json:"rules,optional"`
	Created   types.String                                            `tfsdk:"created" json:"created,computed"`
	State     types.String                                            `tfsdk:"state" json:"state,computed"`
	Updated   types.String                                            `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                            `tfsdk:"uri" json:"uri,computed"`
	Specs     customfield.NestedObjectList[NetworkFirewallSpecsModel] `tfsdk:"specs" json:"specs,computed"`
}

func (m NetworkFirewallModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkFirewallModel) MarshalJSONForUpdate(state NetworkFirewallModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NetworkFirewallRulesModel struct {
	Allow       types.Bool           `tfsdk:"allow" json:"allow,optional"`
	Description types.String         `tfsdk:"description" json:"description,optional"`
	Destination types.String         `tfsdk:"destination" json:"destination,optional"`
	GroupName   types.String         `tfsdk:"group_name" json:"group_name,optional"`
	Inbound     types.Bool           `tfsdk:"inbound" json:"inbound,optional"`
	Port        types.String         `tfsdk:"port" json:"port,optional"`
	Protocol    types.String         `tfsdk:"protocol" json:"protocol,optional"`
	Source      types.String         `tfsdk:"source" json:"source,optional"`
	Zone        jsontypes.Normalized `tfsdk:"zone" json:"zone,optional"`
}

type NetworkFirewallSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
