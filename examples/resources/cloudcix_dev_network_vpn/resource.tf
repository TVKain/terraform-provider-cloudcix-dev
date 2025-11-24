resource "cloudcix-dev_network_vpn" "example_network_vpn" {
  project_id = 1
  metadata = {
    child_sas = [{
      local_ts = "10.0.0.0/24"
      remote_ts = "172.16.10.0/27"
    }]
    ike_authentication = "SHA256"
    ike_dh_groups = "Group 24"
    ike_encryption = "256 bit AES-CBC"
    ike_gateway_type = "hostname"
    ike_gateway_value = "hostname.example.com"
    ike_lifetime = 0
    ike_pre_shared_key = "R4nd0mKey!"
    ike_version = "v2-only"
    ipsec_authentication = "SHA256"
    ipsec_encryption = "AES 256 GCM"
    ipsec_establish_time = "On Traffic"
    ipsec_lifetime = 3000
    ipsec_pfs_groups = "Group 24"
    traffic_selector = true
  }
  name = "Cork Office to Dublin Office"
  type = "site-to-site"
}
