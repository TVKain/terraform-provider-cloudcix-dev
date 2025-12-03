resource "cloudcix-dev_network_router" "example_network_router" {
  networks = [{
    ipv4 = "10.0.1.0/24"
    ipv6 = "ipv6"
    name = "web-tier"
    vlan = 0
  }, {
    ipv4 = "10.0.2.0/24"
    ipv6 = "ipv6"
    name = "app-tier"
    vlan = 0
  }, {
    ipv4 = "10.0.3.0/24"
    ipv6 = "ipv6"
    name = "db-tier"
    vlan = 0
  }]
  project_id = 1
  metadata = {
    destination = "destination"
    nat = true
    nexthop = "nexthop"
  }
  name = "Public Website Router"
  type = "router"
}
