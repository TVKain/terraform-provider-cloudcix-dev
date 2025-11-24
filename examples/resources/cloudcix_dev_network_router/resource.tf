resource "cloudcix-dev_network_router" "example_network_router" {
  project_id = 1
  metadata = {
    destination = "destination"
    nat = true
    nexthop = "nexthop"
  }
  name = "Public Website Router"
  networks = [{
    ipv4 = "10.0.1.0/24"
    name = "web-tier"
  }, {
    ipv4 = "10.0.2.0/24"
    name = "app-tier"
  }, {
    ipv4 = "10.0.3.0/24"
    name = "db-tier"
  }]
  type = "router"
}
