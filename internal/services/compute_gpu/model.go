// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeGPUModel struct {
	Pk      types.Int64                                      `tfsdk:"pk" path:"pk,required"`
	Name    types.String                                     `tfsdk:"name" json:"name,optional,no_refresh"`
	State   types.String                                     `tfsdk:"state" json:"state,optional,no_refresh"`
	Content customfield.NestedObject[ComputeGPUContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m ComputeGPUModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeGPUModel) MarshalJSONForUpdate(state ComputeGPUModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ComputeGPUContentModel struct {
	ID        types.Int64                                               `tfsdk:"id" json:"id,computed"`
	Created   types.String                                              `tfsdk:"created" json:"created,computed"`
	Instance  customfield.NestedObject[ComputeGPUContentInstanceModel]  `tfsdk:"instance" json:"instance,computed"`
	Name      types.String                                              `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                               `tfsdk:"project_id" json:"project_id,computed"`
	Specs     customfield.NestedObjectList[ComputeGPUContentSpecsModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                               `tfsdk:"state" json:"state,computed"`
	Updated   types.String                                              `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                              `tfsdk:"uri" json:"uri,computed"`
}

type ComputeGPUContentInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type ComputeGPUContentSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
