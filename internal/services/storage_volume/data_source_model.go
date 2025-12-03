// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StorageVolumeContentDataSourceEnvelope struct {
	Content StorageVolumeDataSourceModel `json:"content,computed"`
}

type StorageVolumeDataSourceModel struct {
	ID              types.Int64                                                               `tfsdk:"id" path:"pk,computed"`
	Pk              types.Int64                                                               `tfsdk:"pk" path:"pk,required"`
	Created         types.String                                                              `tfsdk:"created" json:"created,computed"`
	Name            types.String                                                              `tfsdk:"name" json:"name,computed"`
	ProjectID       types.Int64                                                               `tfsdk:"project_id" json:"project_id,computed"`
	State           types.String                                                              `tfsdk:"state" json:"state,computed"`
	Type            types.String                                                              `tfsdk:"type" json:"type,computed"`
	Updated         types.String                                                              `tfsdk:"updated" json:"updated,computed"`
	Uri             types.String                                                              `tfsdk:"uri" json:"uri,computed"`
	ContraInstances customfield.NestedObjectList[StorageVolumeContraInstancesDataSourceModel] `tfsdk:"contra_instances" json:"contra_instances,computed"`
	Instance        customfield.NestedObject[StorageVolumeInstanceDataSourceModel]            `tfsdk:"instance" json:"instance,computed"`
	Metadata        customfield.NestedObject[StorageVolumeMetadataDataSourceModel]            `tfsdk:"metadata" json:"metadata,computed"`
	Specs           customfield.NestedObjectList[StorageVolumeSpecsDataSourceModel]           `tfsdk:"specs" json:"specs,computed"`
}

type StorageVolumeContraInstancesDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type StorageVolumeInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type StorageVolumeMetadataDataSourceModel struct {
	AttachInstanceIDs customfield.List[types.Int64] `tfsdk:"attach_instance_ids" json:"attach_instance_ids,computed"`
	DetachInstanceIDs customfield.List[types.Int64] `tfsdk:"detach_instance_ids" json:"detach_instance_ids,computed"`
	MountPath         types.String                  `tfsdk:"mount_path" json:"mount_path,computed"`
}

type StorageVolumeSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
