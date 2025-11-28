// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkRouterDataSourceModel struct {
	ID      types.Int64                                                   `tfsdk:"id" path:"id,required"`
	Content customfield.NestedObject[NetworkRouterContentDataSourceModel] `tfsdk:"content" json:"content,computed"`
}

type NetworkRouterContentDataSourceModel struct {
	ID          types.Int64                                                               `tfsdk:"id" json:"id,computed"`
	Created     types.String                                                              `tfsdk:"created" json:"created,computed"`
	GracePeriod types.Int64                                                               `tfsdk:"grace_period" json:"grace_period,computed"`
	Metadata    customfield.NestedObject[NetworkRouterContentMetadataDataSourceModel]     `tfsdk:"metadata" json:"metadata,computed"`
	Name        types.String                                                              `tfsdk:"name" json:"name,computed"`
	Networks    customfield.NestedObjectList[NetworkRouterContentNetworksDataSourceModel] `tfsdk:"networks" json:"networks,computed"`
	ProjectID   types.Int64                                                               `tfsdk:"project_id" json:"project_id,computed"`
	Specs       customfield.NestedObjectList[NetworkRouterContentSpecsDataSourceModel]    `tfsdk:"specs" json:"specs,computed"`
	State       types.Int64                                                               `tfsdk:"state" json:"state,computed"`
	Type        types.String                                                              `tfsdk:"type" json:"type,computed"`
	Updated     types.String                                                              `tfsdk:"updated" json:"updated,computed"`
	Uri         types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type NetworkRouterContentMetadataDataSourceModel struct {
	Ipv4Address   customfield.NestedObject[NetworkRouterContentMetadataIpv4AddressDataSourceModel] `tfsdk:"ipv4_address" json:"ipv4_address,computed"`
	Ipv4AddressID types.Int64                                                                      `tfsdk:"ipv4_address_id" json:"ipv4_address_id,computed"`
	Ipv6Address   customfield.NestedObject[NetworkRouterContentMetadataIpv6AddressDataSourceModel] `tfsdk:"ipv6_address" json:"ipv6_address,computed"`
	Ipv6AddressID types.Int64                                                                      `tfsdk:"ipv6_address_id" json:"ipv6_address_id,computed"`
}

type NetworkRouterContentMetadataIpv4AddressDataSourceModel struct {
	ID         types.Int64                                                                              `tfsdk:"id" json:"id,computed"`
	Address    types.String                                                                             `tfsdk:"address" json:"address,computed"`
	Created    types.String                                                                             `tfsdk:"created" json:"created,computed"`
	Name       types.String                                                                             `tfsdk:"name" json:"name,computed"`
	Notes      types.String                                                                             `tfsdk:"notes" json:"notes,computed"`
	PublicIP   customfield.NestedObject[NetworkRouterContentMetadataIpv4AddressPublicIPDataSourceModel] `tfsdk:"public_ip" json:"public_ip,computed"`
	PublicIPID types.Int64                                                                              `tfsdk:"public_ip_id" json:"public_ip_id,computed"`
	SubnetID   types.Int64                                                                              `tfsdk:"subnet_id" json:"subnet_id,computed"`
	Updated    types.String                                                                             `tfsdk:"updated" json:"updated,computed"`
}

type NetworkRouterContentMetadataIpv4AddressPublicIPDataSourceModel struct {
	ID      types.Int64  `tfsdk:"id" json:"id,computed"`
	Address types.String `tfsdk:"address" json:"address,computed"`
}

type NetworkRouterContentMetadataIpv6AddressDataSourceModel struct {
	ID         types.Int64                                                                              `tfsdk:"id" json:"id,computed"`
	Address    types.String                                                                             `tfsdk:"address" json:"address,computed"`
	Created    types.String                                                                             `tfsdk:"created" json:"created,computed"`
	Name       types.String                                                                             `tfsdk:"name" json:"name,computed"`
	Notes      types.String                                                                             `tfsdk:"notes" json:"notes,computed"`
	PublicIP   customfield.NestedObject[NetworkRouterContentMetadataIpv6AddressPublicIPDataSourceModel] `tfsdk:"public_ip" json:"public_ip,computed"`
	PublicIPID types.Int64                                                                              `tfsdk:"public_ip_id" json:"public_ip_id,computed"`
	SubnetID   types.Int64                                                                              `tfsdk:"subnet_id" json:"subnet_id,computed"`
	Updated    types.String                                                                             `tfsdk:"updated" json:"updated,computed"`
}

type NetworkRouterContentMetadataIpv6AddressPublicIPDataSourceModel struct {
	ID      types.Int64  `tfsdk:"id" json:"id,computed"`
	Address types.String `tfsdk:"address" json:"address,computed"`
}

type NetworkRouterContentNetworksDataSourceModel struct {
	Destination types.String `tfsdk:"destination" json:"destination,computed"`
	Ipv4        types.String `tfsdk:"ipv4" json:"ipv4,computed"`
	Ipv6        types.String `tfsdk:"ipv6" json:"ipv6,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Nat         types.Bool   `tfsdk:"nat" json:"nat,computed"`
	Nexthop     types.String `tfsdk:"nexthop" json:"nexthop,computed"`
	Vlan        types.Int64  `tfsdk:"vlan" json:"vlan,computed"`
}

type NetworkRouterContentSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
