// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeGPUContentDataSourceEnvelope struct {
	Content ComputeGPUDataSourceModel `json:"content,computed"`
}

type ComputeGPUDataSourceModel struct {
	ID        types.Int64                                                  `tfsdk:"id" path:"pk,computed"`
	Pk        types.Int64                                                  `tfsdk:"pk" path:"pk,required"`
	Created   types.String                                                 `tfsdk:"created" json:"created,computed"`
	Name      types.String                                                 `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                  `tfsdk:"project_id" json:"project_id,computed"`
	State     types.String                                                 `tfsdk:"state" json:"state,computed"`
	Updated   types.String                                                 `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                 `tfsdk:"uri" json:"uri,computed"`
	Instance  customfield.NestedObject[ComputeGPUInstanceDataSourceModel]  `tfsdk:"instance" json:"instance,computed"`
	Specs     customfield.NestedObjectList[ComputeGPUSpecsDataSourceModel] `tfsdk:"specs" json:"specs,computed"`
}

type ComputeGPUInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type ComputeGPUSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
