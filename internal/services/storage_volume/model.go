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
	InstanceID      types.Int64                                                     `tfsdk:"instance_id" json:"instance_id,optional,no_refresh"`
	Type            types.String                                                    `tfsdk:"type" json:"type,optional"`
	Specs           *[]*StorageVolumeSpecsModel                                     `tfsdk:"specs" json:"specs,required"`
	Name            types.String                                                    `tfsdk:"name" json:"name,optional"`
	State           types.String                                                    `tfsdk:"state" json:"state,optional,no_refresh"`
	Metadata        *StorageVolumeMetadataModel                                     `tfsdk:"metadata" json:"metadata,optional"`
	Created         types.String                                                    `tfsdk:"created" json:"created,computed"`
	Updated         types.String                                                    `tfsdk:"updated" json:"updated,computed"`
	Uri             types.String                                                    `tfsdk:"uri" json:"uri,computed"`
	ContraInstances customfield.NestedObjectList[StorageVolumeContraInstancesModel] `tfsdk:"contra_instances" json:"contra_instances,computed"`
	Instance        customfield.NestedObject[StorageVolumeInstanceModel]            `tfsdk:"instance" json:"instance,computed"`
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

type StorageVolumeMetadataModel struct {
	AttachInstanceIDs *[]types.Int64 `tfsdk:"attach_instance_ids" json:"attach_instance_ids,optional"`
	DetachInstanceIDs *[]types.Int64 `tfsdk:"detach_instance_ids" json:"detach_instance_ids,optional"`
	MountPath         types.String   `tfsdk:"mount_path" json:"mount_path,optional"`
}

type StorageVolumeContraInstancesModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type StorageVolumeInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}
