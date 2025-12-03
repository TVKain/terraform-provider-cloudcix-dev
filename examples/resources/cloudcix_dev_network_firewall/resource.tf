resource "cloudcix-dev_network_firewall" "example_network_firewall" {
  project_id = 1
  name = "Allow Traffic from Ireland"
  rules = [{
    allow = true
    description = "description"
    destination = "destination"
    group_name = "@IE_v4"
    inbound = true
    port = "port"
    protocol = "protocol"
    source = "source"
    zone = {

    }
  }]
  type = "geo"
}
