resource "cloudcix-dev_compute_instance" "example_compute_instance" {
  metadata = {
    dns = "dns"
    instance_type = "instance_type"
    userdata = "userdata"
  }
  project_id = 1
  specs = [{
    quantity = 8
    sku_name = "RAM_001"
  }, {
    quantity = 4
    sku_name = "vCPU_002"
  }, {
    quantity = 100
    sku_name = "SSD_001"
  }, {
    quantity = 1
    sku_name = "MSServer2022"
  }]
  grace_period = 72
  interfaces = [{
    gateway = true
    ipv4_addresses = [{
      address = "10.0.0.10"
      nat = true
    }]
    ipv6_addresses = [{
      address = "2a02:2078:9:1234::20"
    }]
  }]
  name = "windows-server"
  type = "hyperv"
}
