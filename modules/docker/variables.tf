# Remote docker image
variable "docker_image" { type = string }
# Docker container name
variable "docker_name" { type = string }
# Working dir
variable "working_dir" {
  type    = string
  default = null
}
# Restart policy
variable "restart" {
  type    = string
  default = "unless-stopped"
}
# If we need full privileges
variable "privileged" {
  type    = bool
  default = false
}
# Network mode
variable "network_mode" {
  type    = string
  default = null
}
# Custom dns for container
variable "dns" {
  type    = list(string)
  default = null
}
# Custom command
variable "command" {
  type    = list(string)
  default = null
}
# Exported ports
variable "ports" {
  type = list(object({
    internal = number
    external = number
    protocol = string
  }))
  default = null
}
# Mounted volumes
variable "volumes" {
  type = list(object({
    volume_name    = string
    container_path = string
    host_path      = string
    read_only      = bool
  }))
  default = null
}
# Exported devices
variable "devices" {
  type = list(object({
    host_path      = string
    container_path = string
    permissions    = string
  }))
  default = null
}
# Additional capabilities
variable "capabilities" {
  type = object({
    add  = list(string)
    drop = list(string)
  })
  default = null
}
# Custom network
variable "networks_advanced" {
  type = object({
    name         = string
    ipv4_address = string
  })
  default = null
}
# Healthcheck
variable "healthcheck" {
  type = object({
    interval     = string
    retries      = number
    start_period = string
    test         = list(string)
    timeout      = string
  })
  default = null
}
# Environment variables
variable "env" {
  type    = list(string)
  default = null
}
# Docker volumes to create
variable "docker_volumes" {
  type    = list(string)
  default = null
}
# Docker networks to create
variable "docker_networks" {
  type = list(object({
    name = string
    ipam_config = object({
      aux_address = map(string)
      gateway     = string
      subnet      = string
    })
  }))
  default = null
}
