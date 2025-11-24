resource "cloudcix-dev_compute_snapshot" "example_compute_snapshot" {
  instance_id = 101
  project_id = 1
  name = "before-windows-updates"
  type = "hyperv"
}
