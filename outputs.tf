output "volumes" {
  value = docker_container.default.volumes
}

output "devices" {
  value = docker_container.default.devices
}

output "environment" {
  value = docker_container.default.env
}
