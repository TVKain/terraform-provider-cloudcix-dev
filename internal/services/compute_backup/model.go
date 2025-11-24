// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeBackupModel struct {
	Pk         types.Int64                                         `tfsdk:"pk" path:"pk,optional"`
	InstanceID types.Int64                                         `tfsdk:"instance_id" json:"instance_id,required,no_refresh"`
	ProjectID  types.Int64                                         `tfsdk:"project_id" json:"project_id,required,no_refresh"`
	Type       types.String                                        `tfsdk:"type" json:"type,optional,no_refresh"`
	Name       types.String                                        `tfsdk:"name" json:"name,optional,no_refresh"`
	State      types.String                                        `tfsdk:"state" json:"state,optional,no_refresh"`
	Content    customfield.NestedObject[ComputeBackupContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m ComputeBackupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeBackupModel) MarshalJSONForUpdate(state ComputeBackupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ComputeBackupContentModel struct {
	ID        types.Int64                                                  `tfsdk:"id" json:"id,computed"`
	Created   types.String                                                 `tfsdk:"created" json:"created,computed"`
	Instance  customfield.NestedObject[ComputeBackupContentInstanceModel]  `tfsdk:"instance" json:"instance,computed"`
	Name      types.String                                                 `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                  `tfsdk:"project_id" json:"project_id,computed"`
	Specs     customfield.NestedObjectList[ComputeBackupContentSpecsModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                                  `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                 `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                 `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                 `tfsdk:"uri" json:"uri,computed"`
}

type ComputeBackupContentInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type ComputeBackupContentSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
