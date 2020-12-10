module "docker" {
  source = "../.."

  image          = "nginx"
  container_name = var.container_name
  restart_policy = "always"
  docker_networks = {
    (var.network_name) = {
      ipam_config = {
        aux_address = {}
        gateway     = var.gateway
        subnet      = var.subnet
      }
    }
  }
  ports         = var.ports
  named_volumes = var.named_volumes
  host_paths    = var.host_paths
  devices       = var.devices
  networks_advanced = {
    name         = var.network_name
    ipv4_address = var.ipv4_address
    ipv6_address = null
    aliases      = null
  }
}
