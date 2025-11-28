// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_snapshot

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeSnapshotContentEnvelope struct {
	Content ComputeSnapshotModel `json:"content"`
}

type ComputeSnapshotModel struct {
	ID         types.Int64                                             `tfsdk:"id" json:"id,computed"`
	InstanceID types.Int64                                             `tfsdk:"instance_id" json:"instance_id,required,no_refresh"`
	ProjectID  types.Int64                                             `tfsdk:"project_id" json:"project_id,required"`
	Type       types.String                                            `tfsdk:"type" json:"type,optional"`
	Name       types.String                                            `tfsdk:"name" json:"name,optional"`
	State      types.String                                            `tfsdk:"state" json:"state,optional,no_refresh"`
	Created    types.String                                            `tfsdk:"created" json:"created,computed"`
	Updated    types.String                                            `tfsdk:"updated" json:"updated,computed"`
	Uri        types.String                                            `tfsdk:"uri" json:"uri,computed"`
	Instance   customfield.NestedObject[ComputeSnapshotInstanceModel]  `tfsdk:"instance" json:"instance,computed"`
	Metadata   customfield.NestedObject[ComputeSnapshotMetadataModel]  `tfsdk:"metadata" json:"metadata,computed"`
	Specs      customfield.NestedObjectList[ComputeSnapshotSpecsModel] `tfsdk:"specs" json:"specs,computed"`
}

func (m ComputeSnapshotModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeSnapshotModel) MarshalJSONForUpdate(state ComputeSnapshotModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ComputeSnapshotInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type ComputeSnapshotMetadataModel struct {
	Active        types.Bool `tfsdk:"active" json:"active,computed"`
	RemoveSubtree types.Bool `tfsdk:"remove_subtree" json:"remove_subtree,computed"`
}

type ComputeSnapshotSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
