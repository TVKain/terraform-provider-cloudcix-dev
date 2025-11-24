resource "cloudcix-dev_network_ip_group" "example_network_ip_group" {
  cidrs = ["91.103.3.0/24", "90.103.2.36", "185.94.188.0/24"]
  name = "office-networks"
  version = 4
}
