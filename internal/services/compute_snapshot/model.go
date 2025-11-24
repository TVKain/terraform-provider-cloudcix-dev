// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_snapshot

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeSnapshotModel struct {
	Pk         types.Int64                                           `tfsdk:"pk" path:"pk,optional"`
	InstanceID types.Int64                                           `tfsdk:"instance_id" json:"instance_id,required,no_refresh"`
	ProjectID  types.Int64                                           `tfsdk:"project_id" json:"project_id,required,no_refresh"`
	Type       types.String                                          `tfsdk:"type" json:"type,optional,no_refresh"`
	Name       types.String                                          `tfsdk:"name" json:"name,optional,no_refresh"`
	State      types.String                                          `tfsdk:"state" json:"state,optional,no_refresh"`
	Content    customfield.NestedObject[ComputeSnapshotContentModel] `tfsdk:"content" json:"content,computed"`
}

func (m ComputeSnapshotModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeSnapshotModel) MarshalJSONForUpdate(state ComputeSnapshotModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ComputeSnapshotContentModel struct {
	ID        types.Int64                                                    `tfsdk:"id" json:"id,computed"`
	Created   types.String                                                   `tfsdk:"created" json:"created,computed"`
	Instance  customfield.NestedObject[ComputeSnapshotContentInstanceModel]  `tfsdk:"instance" json:"instance,computed"`
	Metadata  customfield.NestedObject[ComputeSnapshotContentMetadataModel]  `tfsdk:"metadata" json:"metadata,computed"`
	Name      types.String                                                   `tfsdk:"name" json:"name,computed"`
	ProjectID types.Int64                                                    `tfsdk:"project_id" json:"project_id,computed"`
	Specs     customfield.NestedObjectList[ComputeSnapshotContentSpecsModel] `tfsdk:"specs" json:"specs,computed"`
	State     types.Int64                                                    `tfsdk:"state" json:"state,computed"`
	Type      types.String                                                   `tfsdk:"type" json:"type,computed"`
	Updated   types.String                                                   `tfsdk:"updated" json:"updated,computed"`
	Uri       types.String                                                   `tfsdk:"uri" json:"uri,computed"`
}

type ComputeSnapshotContentInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type ComputeSnapshotContentMetadataModel struct {
	Active        types.Bool `tfsdk:"active" json:"active,computed"`
	RemoveSubtree types.Bool `tfsdk:"remove_subtree" json:"remove_subtree,computed"`
}

type ComputeSnapshotContentSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
