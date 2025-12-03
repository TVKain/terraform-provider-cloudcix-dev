// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeBackupContentEnvelope struct {
	Content ComputeBackupModel `json:"content"`
}

type ComputeBackupModel struct {
	ID         types.Int64                                           `tfsdk:"id" json:"id,computed"`
	InstanceID types.Int64                                           `tfsdk:"instance_id" json:"instance_id,required"`
	ProjectID  types.Int64                                           `tfsdk:"project_id" json:"project_id,required"`
	Name       types.String                                          `tfsdk:"name" json:"name,optional"`
	Type       types.String                                          `tfsdk:"type" json:"type,optional"`
	Created    types.String                                          `tfsdk:"created" json:"created,computed"`
	State      types.String                                          `tfsdk:"state" json:"state,computed"`
	Updated    types.String                                          `tfsdk:"updated" json:"updated,computed"`
	Uri        types.String                                          `tfsdk:"uri" json:"uri,computed"`
	Instance   customfield.NestedObject[ComputeBackupInstanceModel]  `tfsdk:"instance" json:"instance,computed"`
	Specs      customfield.NestedObjectList[ComputeBackupSpecsModel] `tfsdk:"specs" json:"specs,computed"`
}

func (m ComputeBackupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeBackupModel) MarshalJSONForUpdate(state ComputeBackupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ComputeBackupInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type ComputeBackupSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
