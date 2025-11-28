// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeInstanceContentDataSourceEnvelope struct {
	Content ComputeInstanceDataSourceModel `json:"content,computed"`
}

type ComputeInstanceDataSourceModel struct {
	ID          types.Int64                                                            `tfsdk:"id" path:"id,required"`
	Created     types.String                                                           `tfsdk:"created" json:"created,computed"`
	GracePeriod types.Int64                                                            `tfsdk:"grace_period" json:"grace_period,computed"`
	Name        types.String                                                           `tfsdk:"name" json:"name,computed"`
	ProjectID   types.Int64                                                            `tfsdk:"project_id" json:"project_id,computed"`
	State       types.Int64                                                            `tfsdk:"state" json:"state,computed"`
	Type        types.String                                                           `tfsdk:"type" json:"type,computed"`
	Updated     types.String                                                           `tfsdk:"updated" json:"updated,computed"`
	Uri         types.String                                                           `tfsdk:"uri" json:"uri,computed"`
	Interfaces  customfield.NestedObjectList[ComputeInstanceInterfacesDataSourceModel] `tfsdk:"interfaces" json:"interfaces,computed"`
	Metadata    customfield.NestedObject[ComputeInstanceMetadataDataSourceModel]       `tfsdk:"metadata" json:"metadata,computed"`
	Specs       customfield.NestedObject[ComputeInstanceSpecsDataSourceModel]          `tfsdk:"specs" json:"specs,computed"`
}

type ComputeInstanceInterfacesDataSourceModel struct {
	Gateway       types.Bool                                                                          `tfsdk:"gateway" json:"gateway,computed"`
	Ipv4Addresses customfield.NestedObjectList[ComputeInstanceInterfacesIpv4AddressesDataSourceModel] `tfsdk:"ipv4_addresses" json:"ipv4_addresses,computed"`
	Ipv6Addresses customfield.NestedObjectList[ComputeInstanceInterfacesIpv6AddressesDataSourceModel] `tfsdk:"ipv6_addresses" json:"ipv6_addresses,computed"`
	Vlan          types.Int64                                                                         `tfsdk:"vlan" json:"vlan,computed"`
}

type ComputeInstanceInterfacesIpv4AddressesDataSourceModel struct {
	ID       types.Int64  `tfsdk:"id" json:"id,computed"`
	Address  types.String `tfsdk:"address" json:"address,computed"`
	Nat      types.Bool   `tfsdk:"nat" json:"nat,computed"`
	PublicIP types.String `tfsdk:"public_ip" json:"public_ip,computed"`
}

type ComputeInstanceInterfacesIpv6AddressesDataSourceModel struct {
	ID      types.Int64  `tfsdk:"id" json:"id,computed"`
	Address types.String `tfsdk:"address" json:"address,computed"`
}

type ComputeInstanceMetadataDataSourceModel struct {
	DNS          types.String `tfsdk:"dns" json:"dns,computed"`
	Dns4         types.String `tfsdk:"dns4" json:"dns4,computed"`
	Dns6         types.String `tfsdk:"dns6" json:"dns6,computed"`
	InstanceType types.String `tfsdk:"instance_type" json:"instance_type,computed"`
	Userdata     types.String `tfsdk:"userdata" json:"userdata,computed"`
}

type ComputeInstanceSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
