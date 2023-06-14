module "docker" {
  source = "../.."

  image             = "nginx"
  container_name    = var.container_name
  restart_policy    = "always"
  environment       = var.environment
  docker_networks   = var.docker_networks
  ports             = var.ports
  named_volumes     = var.named_volumes
  host_paths        = var.host_paths
  devices           = var.devices
  networks_advanced = var.networks_advanced
}
