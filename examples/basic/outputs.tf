output "image_name" {
  value = module.docker.image_name
}

output "container_name" {
  value = module.docker.container_name
}

output "hostname" {
  value = module.docker.hostname
}

output "working_dir" {
  value = module.docker.working_dir
}

output "restart" {
  value = module.docker.restart
}

output "privileged" {
  value = module.docker.privileged
}

output "network_mode" {
  value = module.docker.network_mode
}

output "dns" {
  value = module.docker.dns
}

output "entrypoint" {
  value = module.docker.entrypoint
}

output "command" {
  value = module.docker.command
}

output "ports" {
  value = module.docker.ports
}

output "volumes" {
  value = module.docker.volumes
}

output "docker_volumes" {
  value = module.docker.docker_volumes
}

output "devices" {
  value = module.docker.devices
}

output "capabilities" {
  value = module.docker.capabilities
}

output "networks_advanced" {
  value = module.docker.networks_advanced
}

output "healthcheck" {
  value = module.docker.healthcheck
}

output "environment" {
  value = module.docker.environment
}

output "docker_networks" {
  value = module.docker.docker_networks
}
