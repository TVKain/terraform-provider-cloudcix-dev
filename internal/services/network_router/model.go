// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkRouterModel struct {
	ID        types.Int64                                         `tfsdk:"id" path:"id,optional"`
	ProjectID types.Int64                                         `tfsdk:"project_id" json:"project_id,required,no_refresh"`
	Type      types.String                                        `tfsdk:"type" json:"type,optional,no_refresh"`
	Name      types.String                                        `tfsdk:"name" json:"name,optional,no_refresh"`
	State     types.String                                        `tfsdk:"state" json:"state,optional,no_refresh"`
	Metadata  *NetworkRouterMetadataModel                         `tfsdk:"metadata" json:"metadata,optional,no_refresh"`
	Networks  *[]*NetworkRouterNetworksModel                      `tfsdk:"networks" json:"networks,optional,no_refresh"`
	Content   customfield.NestedObject[NetworkRouterContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m NetworkRouterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkRouterModel) MarshalJSONForUpdate(state NetworkRouterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NetworkRouterMetadataModel struct {
	Destination types.String `tfsdk:"destination" json:"destination,optional"`
	Nat         types.Bool   `tfsdk:"nat" json:"nat,optional"`
	Nexthop     types.String `tfsdk:"nexthop" json:"nexthop,optional"`
}

type NetworkRouterNetworksModel struct {
	Ipv4 types.String `tfsdk:"ipv4" json:"ipv4,optional"`
	Ipv6 types.String `tfsdk:"ipv6" json:"ipv6,optional"`
	Name types.String `tfsdk:"name" json:"name,optional"`
	Vlan types.Int64  `tfsdk:"vlan" json:"vlan,optional"`
}

type NetworkRouterContentModel struct {
	ID          types.Int64                                                     `tfsdk:"id" json:"id,computed"`
	Created     types.String                                                    `tfsdk:"created" json:"created,computed"`
	GracePeriod types.Int64                                                     `tfsdk:"grace_period" json:"grace_period,computed"`
	Metadata    customfield.NestedObject[NetworkRouterContentMetadataModel]     `tfsdk:"metadata" json:"metadata,computed"`
	Name        types.String                                                    `tfsdk:"name" json:"name,computed"`
	Networks    customfield.NestedObjectList[NetworkRouterContentNetworksModel] `tfsdk:"networks" json:"networks,computed"`
	ProjectID   types.Int64                                                     `tfsdk:"project_id" json:"project_id,computed"`
	Specs       customfield.NestedObjectList[NetworkRouterContentSpecsModel]    `tfsdk:"specs" json:"specs,computed"`
	State       types.Int64                                                     `tfsdk:"state" json:"state,computed"`
	Type        types.String                                                    `tfsdk:"type" json:"type,computed"`
	Updated     types.String                                                    `tfsdk:"updated" json:"updated,computed"`
	Uri         types.String                                                    `tfsdk:"uri" json:"uri,computed"`
}

type NetworkRouterContentMetadataModel struct {
	Ipv4Address   customfield.NestedObject[NetworkRouterContentMetadataIpv4AddressModel] `tfsdk:"ipv4_address" json:"ipv4_address,computed"`
	Ipv4AddressID types.Int64                                                            `tfsdk:"ipv4_address_id" json:"ipv4_address_id,computed"`
	Ipv6Address   customfield.NestedObject[NetworkRouterContentMetadataIpv6AddressModel] `tfsdk:"ipv6_address" json:"ipv6_address,computed"`
	Ipv6AddressID types.Int64                                                            `tfsdk:"ipv6_address_id" json:"ipv6_address_id,computed"`
}

type NetworkRouterContentMetadataIpv4AddressModel struct {
	ID         types.Int64                                                                    `tfsdk:"id" json:"id,computed"`
	Address    types.String                                                                   `tfsdk:"address" json:"address,computed"`
	Created    types.String                                                                   `tfsdk:"created" json:"created,computed"`
	Name       types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Notes      types.String                                                                   `tfsdk:"notes" json:"notes,computed"`
	PublicIP   customfield.NestedObject[NetworkRouterContentMetadataIpv4AddressPublicIPModel] `tfsdk:"public_ip" json:"public_ip,computed"`
	PublicIPID types.Int64                                                                    `tfsdk:"public_ip_id" json:"public_ip_id,computed"`
	SubnetID   types.Int64                                                                    `tfsdk:"subnet_id" json:"subnet_id,computed"`
	Updated    types.String                                                                   `tfsdk:"updated" json:"updated,computed"`
}

type NetworkRouterContentMetadataIpv4AddressPublicIPModel struct {
	ID      types.Int64  `tfsdk:"id" json:"id,computed"`
	Address types.String `tfsdk:"address" json:"address,computed"`
}

type NetworkRouterContentMetadataIpv6AddressModel struct {
	ID         types.Int64                                                                    `tfsdk:"id" json:"id,computed"`
	Address    types.String                                                                   `tfsdk:"address" json:"address,computed"`
	Created    types.String                                                                   `tfsdk:"created" json:"created,computed"`
	Name       types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Notes      types.String                                                                   `tfsdk:"notes" json:"notes,computed"`
	PublicIP   customfield.NestedObject[NetworkRouterContentMetadataIpv6AddressPublicIPModel] `tfsdk:"public_ip" json:"public_ip,computed"`
	PublicIPID types.Int64                                                                    `tfsdk:"public_ip_id" json:"public_ip_id,computed"`
	SubnetID   types.Int64                                                                    `tfsdk:"subnet_id" json:"subnet_id,computed"`
	Updated    types.String                                                                   `tfsdk:"updated" json:"updated,computed"`
}

type NetworkRouterContentMetadataIpv6AddressPublicIPModel struct {
	ID      types.Int64  `tfsdk:"id" json:"id,computed"`
	Address types.String `tfsdk:"address" json:"address,computed"`
}

type NetworkRouterContentNetworksModel struct {
	Destination types.String `tfsdk:"destination" json:"destination,computed"`
	Ipv4        types.String `tfsdk:"ipv4" json:"ipv4,computed"`
	Ipv6        types.String `tfsdk:"ipv6" json:"ipv6,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Nat         types.Bool   `tfsdk:"nat" json:"nat,computed"`
	Nexthop     types.String `tfsdk:"nexthop" json:"nexthop,computed"`
	Vlan        types.Int64  `tfsdk:"vlan" json:"vlan,computed"`
}

type NetworkRouterContentSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
