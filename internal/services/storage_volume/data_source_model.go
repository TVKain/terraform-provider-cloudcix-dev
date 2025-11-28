// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume

import (
	"github.com/TVKain/terraform-provider-cloudcix-dev/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StorageVolumeDataSourceModel struct {
	ID      types.Int64                                                   `tfsdk:"id" path:"id,required"`
	Content customfield.NestedObject[StorageVolumeContentDataSourceModel] `tfsdk:"content" json:"content,computed"`
}

type StorageVolumeContentDataSourceModel struct {
	ID              types.Int64                                                                      `tfsdk:"id" json:"id,computed"`
	ContraInstances customfield.NestedObjectList[StorageVolumeContentContraInstancesDataSourceModel] `tfsdk:"contra_instances" json:"contra_instances,computed"`
	Created         types.String                                                                     `tfsdk:"created" json:"created,computed"`
	Instance        customfield.NestedObject[StorageVolumeContentInstanceDataSourceModel]            `tfsdk:"instance" json:"instance,computed"`
	Metadata        customfield.NestedObject[StorageVolumeContentMetadataDataSourceModel]            `tfsdk:"metadata" json:"metadata,computed"`
	Name            types.String                                                                     `tfsdk:"name" json:"name,computed"`
	ProjectID       types.Int64                                                                      `tfsdk:"project_id" json:"project_id,computed"`
	Specs           customfield.NestedObjectList[StorageVolumeContentSpecsDataSourceModel]           `tfsdk:"specs" json:"specs,computed"`
	State           types.Int64                                                                      `tfsdk:"state" json:"state,computed"`
	Type            types.String                                                                     `tfsdk:"type" json:"type,computed"`
	Updated         types.String                                                                     `tfsdk:"updated" json:"updated,computed"`
	Uri             types.String                                                                     `tfsdk:"uri" json:"uri,computed"`
}

type StorageVolumeContentContraInstancesDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type StorageVolumeContentInstanceDataSourceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.Int64  `tfsdk:"state" json:"state,computed"`
}

type StorageVolumeContentMetadataDataSourceModel struct {
	AttachInstanceIDs customfield.List[types.Int64] `tfsdk:"attach_instance_ids" json:"attach_instance_ids,computed"`
	DetachInstanceIDs customfield.List[types.Int64] `tfsdk:"detach_instance_ids" json:"detach_instance_ids,computed"`
	MountPath         types.String                  `tfsdk:"mount_path" json:"mount_path,computed"`
}

type StorageVolumeContentSpecsDataSourceModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,computed"`
}
