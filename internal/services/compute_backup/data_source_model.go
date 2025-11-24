// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeBackupDataSourceModel struct {
	Pk      types.Int64                                                   `tfsdk:"pk" path:"pk,required"`
	Content customfield.NestedObject[ComputeBackupContentDataSourceModel] `tfsdk:"content" json:"content,computed"`
}

type ComputeBackupContentDataSourceModel struct {
	ID        types.Int64                                                            `tfsdk:"id" json:"id,computed"`
	Created   types.String                                                           `tfsdk:"created" json:"created,computed"`
	Instance  customfield.NestedObject[ComputeBackupContentInstanceDataSourceModel]  `tfsdk:"instance" json:"instance,computed"`
	Name      types.String                                                           `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                            `tfsdk:"project_id" json:"project_id,computed"`
	Specs     customfield.NestedObjectList[ComputeBackupContentSpecsDataSourceModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                                            `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                           `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                           `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                           `tfsdk:"uri" json:"uri,computed"`
}

type ComputeBackupContentInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type ComputeBackupContentSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
