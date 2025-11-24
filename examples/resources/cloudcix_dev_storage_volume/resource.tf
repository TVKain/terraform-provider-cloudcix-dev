resource "cloudcix-dev_storage_volume" "example_storage_volume" {
  project_id = 1
  specs = [{
    quantity = 500
    sku_name = "CEPH_001"
  }]
  instance_id = 0
  name = "shared-data-volume"
  type = "cephfs"
}
