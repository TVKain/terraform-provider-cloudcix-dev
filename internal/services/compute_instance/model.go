// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeInstanceContentEnvelope struct {
	Content ComputeInstanceModel `json:"content"`
}

type ComputeInstanceModel struct {
	ID          types.Int64                        `tfsdk:"id" json:"id,computed"`
	ProjectID   types.Int64                        `tfsdk:"project_id" json:"project_id,required"`
	Type        types.String                       `tfsdk:"type" json:"type,optional"`
	Metadata    *ComputeInstanceMetadataModel      `tfsdk:"metadata" json:"metadata,required"`
	Specs       *[]*ComputeInstanceSpecsModel      `tfsdk:"specs" json:"specs,required,no_refresh"`
	GracePeriod types.Int64                        `tfsdk:"grace_period" json:"grace_period,optional"`
	Name        types.String                       `tfsdk:"name" json:"name,optional"`
	State       types.String                       `tfsdk:"state" json:"state,optional"`
	Interfaces  *[]*ComputeInstanceInterfacesModel `tfsdk:"interfaces" json:"interfaces,optional"`
	Created     types.String                       `tfsdk:"created" json:"created,computed"`
	Updated     types.String                       `tfsdk:"updated" json:"updated,computed"`
	Uri         types.String                       `tfsdk:"uri" json:"uri,computed"`
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
