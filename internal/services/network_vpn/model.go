// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_vpn

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkVpnModel struct {
	Pk        types.Int64                                      `tfsdk:"pk" path:"pk,optional"`
	ProjectID types.Int64                                      `tfsdk:"project_id" json:"project_id,required,no_refresh"`
	Type      types.String                                     `tfsdk:"type" json:"type,optional,no_refresh"`
	Name      types.String                                     `tfsdk:"name" json:"name,optional,no_refresh"`
	State     types.String                                     `tfsdk:"state" json:"state,optional,no_refresh"`
	Metadata  *NetworkVpnMetadataModel                         `tfsdk:"metadata" json:"metadata,optional,no_refresh"`
	Content   customfield.NestedObject[NetworkVpnContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m NetworkVpnModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NetworkVpnModel) MarshalJSONForUpdate(state NetworkVpnModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NetworkVpnMetadataModel struct {
	ChildSas            *[]*NetworkVpnMetadataChildSasModel `tfsdk:"child_sas" json:"child_sas,optional"`
	IkeAuthentication   types.String                        `tfsdk:"ike_authentication" json:"ike_authentication,optional"`
	IkeDhGroups         types.String                        `tfsdk:"ike_dh_groups" json:"ike_dh_groups,optional"`
	IkeEncryption       types.String                        `tfsdk:"ike_encryption" json:"ike_encryption,optional"`
	IkeGatewayType      types.String                        `tfsdk:"ike_gateway_type" json:"ike_gateway_type,optional"`
	IkeGatewayValue     types.String                        `tfsdk:"ike_gateway_value" json:"ike_gateway_value,optional"`
	IkeLifetime         types.Int64                         `tfsdk:"ike_lifetime" json:"ike_lifetime,optional"`
	IkePreSharedKey     types.String                        `tfsdk:"ike_pre_shared_key" json:"ike_pre_shared_key,optional"`
	IkeVersion          types.String                        `tfsdk:"ike_version" json:"ike_version,optional"`
	IpsecAuthentication types.String                        `tfsdk:"ipsec_authentication" json:"ipsec_authentication,optional"`
	IpsecEncryption     types.String                        `tfsdk:"ipsec_encryption" json:"ipsec_encryption,optional"`
	IpsecEstablishTime  types.String                        `tfsdk:"ipsec_establish_time" json:"ipsec_establish_time,optional"`
	IpsecLifetime       types.Int64                         `tfsdk:"ipsec_lifetime" json:"ipsec_lifetime,optional"`
	IpsecPfsGroups      types.String                        `tfsdk:"ipsec_pfs_groups" json:"ipsec_pfs_groups,optional"`
	TrafficSelector     types.Bool                          `tfsdk:"traffic_selector" json:"traffic_selector,optional"`
}

type NetworkVpnMetadataChildSasModel struct {
	LocalTs  types.String `tfsdk:"local_ts" json:"local_ts,optional"`
	RemoteTs types.String `tfsdk:"remote_ts" json:"remote_ts,optional"`
}

type NetworkVpnContentModel struct {
	ID        types.Int64                                               `tfsdk:"id" json:"id,computed"`
	Created   types.String                                              `tfsdk:"created" json:"created,computed"`
	Metadata  customfield.NestedObject[NetworkVpnContentMetadataModel]  `tfsdk:"metadata" json:"metadata,computed"`
	Name      types.String                                              `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                               `tfsdk:"project_id" json:"project_id,computed"`
	Specs     customfield.NestedObjectList[NetworkVpnContentSpecsModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                               `tfsdk:"state" json:"state,computed"`
	Type      types.String                                              `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                              `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                              `tfsdk:"uri" json:"uri,computed"`
}

type NetworkVpnContentMetadataModel struct {
	ChildSas            customfield.NestedObjectList[NetworkVpnContentMetadataChildSasModel] `tfsdk:"child_sas" json:"child_sas,computed"`
	IkeAuthentication   types.String                                                         `tfsdk:"ike_authentication" json:"ike_authentication,computed"`
	IkeDhGroups         types.String                                                         `tfsdk:"ike_dh_groups" json:"ike_dh_groups,computed"`
	IkeEncryption       types.String                                                         `tfsdk:"ike_encryption" json:"ike_encryption,computed"`
	IkeGatewayType      types.String                                                         `tfsdk:"ike_gateway_type" json:"ike_gateway_type,computed"`
	IkeGatewayValue     types.String                                                         `tfsdk:"ike_gateway_value" json:"ike_gateway_value,computed"`
	IkeLifetime         types.Int64                                                          `tfsdk:"ike_lifetime" json:"ike_lifetime,computed"`
	IkeLocalIdentifier  types.String                                                         `tfsdk:"ike_local_identifier" json:"ike_local_identifier,computed"`
	IkePreSharedKey     types.String                                                         `tfsdk:"ike_pre_shared_key" json:"ike_pre_shared_key,computed"`
	IkeRemoteIdentifier types.String                                                         `tfsdk:"ike_remote_identifier" json:"ike_remote_identifier,computed"`
	IkeVersion          types.Int64                                                          `tfsdk:"ike_version" json:"ike_version,computed"`
	IpsecAuthentication types.String                                                         `tfsdk:"ipsec_authentication" json:"ipsec_authentication,computed"`
	IpsecEncryption     types.String                                                         `tfsdk:"ipsec_encryption" json:"ipsec_encryption,computed"`
	IpsecEstablishTime  types.String                                                         `tfsdk:"ipsec_establish_time" json:"ipsec_establish_time,computed"`
	IpsecLifetime       types.Int64                                                          `tfsdk:"ipsec_lifetime" json:"ipsec_lifetime,computed"`
	IpsecPfsGroups      types.String                                                         `tfsdk:"ipsec_pfs_groups" json:"ipsec_pfs_groups,computed"`
	StifNumber          types.Int64                                                          `tfsdk:"stif_number" json:"stif_number,computed"`
	TrafficSelector     types.Bool                                                           `tfsdk:"traffic_selector" json:"traffic_selector,computed"`
}

type NetworkVpnContentMetadataChildSasModel struct {
	LocalTs  types.String `tfsdk:"local_ts" json:"local_ts,computed"`
	RemoteTs types.String `tfsdk:"remote_ts" json:"remote_ts,computed"`
}

type NetworkVpnContentSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
