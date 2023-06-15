module "docker" {
  source = "../.."

  image             = var.image
  container_name    = var.container_name
  hostname          = var.hostname
  restart_policy    = var.restart_policy
  working_dir       = var.working_dir
  privileged        = var.privileged
  network_mode      = var.network_mode
  dns               = var.dns
  entrypoint        = var.entrypoint
  command           = var.command
  capabilities      = var.capabilities
  environment       = var.environment
  docker_networks   = var.docker_networks
  ports             = var.ports
  named_volumes     = var.named_volumes
  host_paths        = var.host_paths
  devices           = var.devices
  networks_advanced = var.networks_advanced
}
