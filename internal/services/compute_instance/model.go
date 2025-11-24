// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeInstanceModel struct {
	Pk          types.Int64                                           `tfsdk:"pk" path:"pk,optional"`
	ProjectID   types.Int64                                           `tfsdk:"project_id" json:"project_id,required,no_refresh"`
	Type        types.String                                          `tfsdk:"type" json:"type,optional,no_refresh"`
	Metadata    *ComputeInstanceMetadataModel                         `tfsdk:"metadata" json:"metadata,required,no_refresh"`
	Specs       *[]*ComputeInstanceSpecsModel                         `tfsdk:"specs" json:"specs,required,no_refresh"`
	GracePeriod types.Int64                                           `tfsdk:"grace_period" json:"grace_period,optional,no_refresh"`
	Name        types.String                                          `tfsdk:"name" json:"name,optional,no_refresh"`
	State       types.String                                          `tfsdk:"state" json:"state,optional,no_refresh"`
	Interfaces  *[]*ComputeInstanceInterfacesModel                    `tfsdk:"interfaces" json:"interfaces,optional,no_refresh"`
	Content     customfield.NestedObject[ComputeInstanceContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m ComputeInstanceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeInstanceModel) MarshalJSONForUpdate(state ComputeInstanceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ComputeInstanceMetadataModel struct {
	DNS          types.String `tfsdk:"dns" json:"dns,optional"`
	InstanceType types.String `tfsdk:"instance_type" json:"instance_type,optional"`
	Userdata     types.String `tfsdk:"userdata" json:"userdata,optional"`
}

type ComputeInstanceSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,optional"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,optional"`
}

type ComputeInstanceInterfacesModel struct {
	Gateway       types.Bool                                      `tfsdk:"gateway" json:"gateway,optional"`
	Ipv4Addresses *[]*ComputeInstanceInterfacesIpv4AddressesModel `tfsdk:"ipv4_addresses" json:"ipv4_addresses,optional"`
	Ipv6Addresses *[]*ComputeInstanceInterfacesIpv6AddressesModel `tfsdk:"ipv6_addresses" json:"ipv6_addresses,optional"`
}

type ComputeInstanceInterfacesIpv4AddressesModel struct {
	Address types.String `tfsdk:"address" json:"address,optional"`
	Nat     types.Bool   `tfsdk:"nat" json:"nat,optional"`
}

type ComputeInstanceInterfacesIpv6AddressesModel struct {
	Address types.String `tfsdk:"address" json:"address,optional"`
}

type ComputeInstanceContentModel struct {
	ID          types.Int64                                                         `tfsdk:"id" json:"id,computed"`
	Created     types.String                                                        `tfsdk:"created" json:"created,computed"`
	GracePeriod types.Int64                                                         `tfsdk:"grace_period" json:"grace_period,computed"`
	Interfaces  customfield.NestedObjectList[ComputeInstanceContentInterfacesModel] `tfsdk:"interfaces" json:"interfaces,computed"`
	Metadata    customfield.NestedObject[ComputeInstanceContentMetadataModel]       `tfsdk:"metadata" json:"metadata,computed"`
	Name        types.String                                                        `tfsdk:"name" json:"name,computed"`
	ProjectID   types.Int64                                                         `tfsdk:"project_id" json:"project_id,computed"`
	Specs       customfield.NestedObject[ComputeInstanceContentSpecsModel]          `tfsdk:"specs" json:"specs,computed"`
	State       types.Int64                                                         `tfsdk:"state" json:"state,computed"`
	Type        types.String                                                        `tfsdk:"type" json:"type,computed"`
	Updated     types.String                                                        `tfsdk:"updated" json:"updated,computed"`
	Uri         types.String                                                        `tfsdk:"uri" json:"uri,computed"`
}

type ComputeInstanceContentInterfacesModel struct {
	Gateway       types.Bool                                                                       `tfsdk:"gateway" json:"gateway,computed"`
	Ipv4Addresses customfield.NestedObjectList[ComputeInstanceContentInterfacesIpv4AddressesModel] `tfsdk:"ipv4_addresses" json:"ipv4_addresses,computed"`
	Ipv6Addresses customfield.NestedObjectList[ComputeInstanceContentInterfacesIpv6AddressesModel] `tfsdk:"ipv6_addresses" json:"ipv6_addresses,computed"`
	Vlan          types.Int64                                                                      `tfsdk:"vlan" json:"vlan,computed"`
}

type ComputeInstanceContentInterfacesIpv4AddressesModel struct {
	ID       types.Int64  `tfsdk:"id" json:"id,computed"`
	Address  types.String `tfsdk:"address" json:"address,computed"`
	Nat      types.Bool   `tfsdk:"nat" json:"nat,computed"`
	PublicIP types.String `tfsdk:"public_ip" json:"public_ip,computed"`
}

type ComputeInstanceContentInterfacesIpv6AddressesModel struct {
	ID      types.Int64  `tfsdk:"id" json:"id,computed"`
	Address types.String `tfsdk:"address" json:"address,computed"`
}

type ComputeInstanceContentMetadataModel struct {
	DNS          types.String `tfsdk:"dns" json:"dns,computed"`
	Dns4         types.String `tfsdk:"dns4" json:"dns4,computed"`
	Dns6         types.String `tfsdk:"dns6" json:"dns6,computed"`
	InstanceType types.String `tfsdk:"instance_type" json:"instance_type,computed"`
	Userdata     types.String `tfsdk:"userdata" json:"userdata,computed"`
}

type ComputeInstanceContentSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
