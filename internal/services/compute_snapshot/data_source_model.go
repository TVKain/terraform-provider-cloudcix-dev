// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_snapshot

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeSnapshotDataSourceModel struct {
	Pk      types.Int64                                                     `tfsdk:"pk" path:"pk,required"`
	Content customfield.NestedObject[ComputeSnapshotContentDataSourceModel] `tfsdk:"content" json:"content,computed"`
}

type ComputeSnapshotContentDataSourceModel struct {
	ID        types.Int64                                                              `tfsdk:"id" json:"id,computed"`
	Created   types.String                                                             `tfsdk:"created" json:"created,computed"`
	Instance  customfield.NestedObject[ComputeSnapshotContentInstanceDataSourceModel]  `tfsdk:"instance" json:"instance,computed"`
	Metadata  customfield.NestedObject[ComputeSnapshotContentMetadataDataSourceModel]  `tfsdk:"metadata" json:"metadata,computed"`
	Name      types.String                                                             `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                              `tfsdk:"project_id" json:"project_id,computed"`
	Specs     customfield.NestedObjectList[ComputeSnapshotContentSpecsDataSourceModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                                              `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                             `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                             `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                             `tfsdk:"uri" json:"uri,computed"`
}

type ComputeSnapshotContentInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type ComputeSnapshotContentMetadataDataSourceModel struct {
	Active        types.Bool `tfsdk:"active" json:"active,computed"`
	RemoveSubtree types.Bool `tfsdk:"remove_subtree" json:"remove_subtree,computed"`
}

type ComputeSnapshotContentSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
