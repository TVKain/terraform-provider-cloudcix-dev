// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeGPUContentEnvelope struct {
	Content ComputeGPUModel `json:"content"`
}

type ComputeGPUModel struct {
	ID        types.Int64                                        `tfsdk:"id" path:"id,required"`
	Name      types.String                                       `tfsdk:"name" json:"name,optional"`
	State     types.String                                       `tfsdk:"state" json:"state,optional"`
	Created   types.String                                       `tfsdk:"created" json:"created,computed"`
	ProjectID types.Int64                                        `tfsdk:"project_id" json:"project_id,computed"`
	Updated   types.String                                       `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                       `tfsdk:"uri" json:"uri,computed"`
	Instance  customfield.NestedObject[ComputeGPUInstanceModel]  `tfsdk:"instance" json:"instance,computed"`
	Specs     customfield.NestedObjectList[ComputeGPUSpecsModel] `tfsdk:"specs" json:"specs,computed"`
}

func (m ComputeGPUModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeGPUModel) MarshalJSONForUpdate(state ComputeGPUModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ComputeGPUInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type ComputeGPUSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
