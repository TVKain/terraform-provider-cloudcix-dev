resource "cloudcix-dev_compute_backup" "example_compute_backup" {
  instance_id = 456
  project_id = 1
  name = "critical-windows-backup"
  type = "hyperv"
}
