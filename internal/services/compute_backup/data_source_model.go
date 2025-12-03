// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeBackupContentDataSourceEnvelope struct {
	Content ComputeBackupDataSourceModel `json:"content,computed"`
}

type ComputeBackupDataSourceModel struct {
	ID        types.Int64                                                     `tfsdk:"id" path:"id,required"`
	Created   types.String                                                    `tfsdk:"created" json:"created,computed"`
	Name      types.String                                                    `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                     `tfsdk:"project_id" json:"project_id,computed"`
	State     types.String                                                    `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                    `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                    `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                    `tfsdk:"uri" json:"uri,computed"`
	Instance  customfield.NestedObject[ComputeBackupInstanceDataSourceModel]  `tfsdk:"instance" json:"instance,computed"`
	Specs     customfield.NestedObjectList[ComputeBackupSpecsDataSourceModel] `tfsdk:"specs" json:"specs,computed"`
}

type ComputeBackupInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type ComputeBackupSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
