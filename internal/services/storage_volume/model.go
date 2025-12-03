// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/apijson"
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StorageVolumeContentEnvelope struct {
	Content StorageVolumeModel `json:"content"`
}

type StorageVolumeModel struct {
	ID              types.Int64                                                     `tfsdk:"id" json:"id,computed"`
	ProjectID       types.Int64                                                     `tfsdk:"project_id" json:"project_id,required"`
	Specs           *[]*StorageVolumeSpecsModel                                     `tfsdk:"specs" json:"specs,required"`
	InstanceID      types.Int64                                                     `tfsdk:"instance_id" json:"instance_id,optional"`
	Name            types.String                                                    `tfsdk:"name" json:"name,optional"`
	Type            types.String                                                    `tfsdk:"type" json:"type,optional"`
	Created         types.String                                                    `tfsdk:"created" json:"created,computed"`
	State           types.String                                                    `tfsdk:"state" json:"state,computed"`
	Updated         types.String                                                    `tfsdk:"updated" json:"updated,computed"`
	Uri             types.String                                                    `tfsdk:"uri" json:"uri,computed"`
	ContraInstances customfield.NestedObjectList[StorageVolumeContraInstancesModel] `tfsdk:"contra_instances" json:"contra_instances,computed"`
	Instance        customfield.NestedObject[StorageVolumeInstanceModel]            `tfsdk:"instance" json:"instance,computed"`
	Metadata        customfield.NestedObject[StorageVolumeMetadataModel]            `tfsdk:"metadata" json:"metadata,computed"`
}

func (m StorageVolumeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StorageVolumeModel) MarshalJSONForUpdate(state StorageVolumeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type StorageVolumeSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,optional"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,optional"`
}

type StorageVolumeContraInstancesModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type StorageVolumeInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type StorageVolumeMetadataModel struct {
	AttachInstanceIDs customfield.List[types.Int64] `tfsdk:"attach_instance_ids" json:"attach_instance_ids,computed"`
	DetachInstanceIDs customfield.List[types.Int64] `tfsdk:"detach_instance_ids" json:"detach_instance_ids,computed"`
	MountPath         types.String                  `tfsdk:"mount_path" json:"mount_path,computed"`
}
