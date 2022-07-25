resource "aws_vpc" "example" {
    cidr_block = "10.0.0.0/16"
    instance_tenancy = "default"
    //ipv4_ipam_pool_id = 
    //ipv4_netmask_length = 
    # ipv6_ipam_pool_id = 
    # ipv6_netmask_length = 
    # ipv6_cidr_block = 
    # ipv6_cidr_block_network_border_group = 
    # enable_dns_support = 
    # enable_dns_hostnames = 

    assign_generated_ipv6_cidr_block = true
    tags = {
      "test" = "pass"
    }
}