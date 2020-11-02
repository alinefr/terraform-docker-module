variable "docker_image" { type = string }
variable "docker_name" { type = string }
variable "working_dir" {
  type    = string
  default = null
}
variable "restart" {
  type    = string
  default = "unless-stopped"
}
variable "privileged" {
  type    = bool
  default = false
}
variable "network_mode" {
  type    = string
  default = null
}
variable "dns" {
  type    = list(string)
  default = null
}
variable "command" {
  type    = list(string)
  default = null
}
variable "ports" {
  type = list(object({
    internal = number
    external = number
    protocol = string
  }))
  default = null
}
variable "volumes" {
  type = list(object({
    volume_name    = string
    container_path = string
    host_path      = string
    read_only      = bool
  }))
  default = null
}
variable "devices" {
  type = list(object({
    host_path      = string
    container_path = string
    permissions    = string
  }))
  default = null
}
variable "capabilities" {
  type = object({
    add  = list(string)
    drop = list(string)
  })
  default = null
}
variable "networks_advanced" {
  type = object({
    name         = string
    ipv4_address = string
  })
  default = null
}
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
variable "env" {
  type    = list(string)
  default = null
}
variable "docker_volumes" {
  type    = list(string)
  default = null
}
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
