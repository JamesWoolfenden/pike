resource "aws_fsx_file_cache" "example" {

  data_repository_association {
    data_repository_path           = "nfs://filer.domain.com"
    data_repository_subdirectories = ["test", "test2"]
    file_cache_path                = "/ns1"

    nfs {
      dns_ips = ["192.168.0.1", "192.168.0.2"]
      version = "NFS3"
    }
  }

  file_cache_type         = "LUSTRE"
  file_cache_type_version = "2.12"

  lustre_configuration {
    deployment_type = "CACHE_1"
    metadata_configuration {
      storage_capacity = 2400
    }
    per_unit_storage_throughput   = 1000
    weekly_maintenance_start_time = "2:05:00"
  }

  subnet_ids       = [data.aws_subnets.example.ids[0]]
  storage_capacity = 1200
}

data "aws_subnets" "example" {}
