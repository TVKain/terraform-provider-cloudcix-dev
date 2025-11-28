// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeGPUDataSourceModel struct {
	ID      types.Int64                                                `tfsdk:"id" path:"id,required"`
	Content customfield.NestedObject[ComputeGPUContentDataSourceModel] `tfsdk:"content" json:"content,computed"`
}

type ComputeGPUContentDataSourceModel struct {
	ID        types.Int64                                                         `tfsdk:"id" json:"id,computed"`
	Created   types.String                                                        `tfsdk:"created" json:"created,computed"`
	Instance  customfield.NestedObject[ComputeGPUContentInstanceDataSourceModel]  `tfsdk:"instance" json:"instance,computed"`
	Name      types.String                                                        `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                         `tfsdk:"project_id" json:"project_id,computed"`
	Specs     customfield.NestedObjectList[ComputeGPUContentSpecsDataSourceModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                                         `tfsdk:"state" json:"state,computed"`
	Updated   types.String                                                        `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                        `tfsdk:"uri" json:"uri,computed"`
}

type ComputeGPUContentInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type ComputeGPUContentSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
