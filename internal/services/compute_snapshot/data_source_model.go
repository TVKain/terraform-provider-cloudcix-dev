// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_snapshot

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeSnapshotContentDataSourceEnvelope struct {
	Content ComputeSnapshotDataSourceModel `json:"content,computed"`
}

type ComputeSnapshotDataSourceModel struct {
	ID        types.Int64                                                       `tfsdk:"id" path:"id,required"`
	Created   types.String                                                      `tfsdk:"created" json:"created,computed"`
	Name      types.String                                                      `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                       `tfsdk:"project_id" json:"project_id,computed"`
	State     types.String                                                      `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                      `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                      `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                      `tfsdk:"uri" json:"uri,computed"`
	Instance  customfield.NestedObject[ComputeSnapshotInstanceDataSourceModel]  `tfsdk:"instance" json:"instance,computed"`
	Metadata  customfield.NestedObject[ComputeSnapshotMetadataDataSourceModel]  `tfsdk:"metadata" json:"metadata,computed"`
	Specs     customfield.NestedObjectList[ComputeSnapshotSpecsDataSourceModel] `tfsdk:"specs" json:"specs,computed"`
}

type ComputeSnapshotInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type ComputeSnapshotMetadataDataSourceModel struct {
	Active        types.Bool `tfsdk:"active" json:"active,computed"`
	RemoveSubtree types.Bool `tfsdk:"remove_subtree" json:"remove_subtree,computed"`
}

type ComputeSnapshotSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
