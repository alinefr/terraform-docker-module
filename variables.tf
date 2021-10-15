variable "image" {
  description = "Specify the image to start the container from. Can either be a repository/tag or a partial image ID"
  type        = string
}
variable "container_name" {
  description = "Custom container name"
  type        = string
  default     = null
}
variable "hostname" {
  description = "Set docker hostname"
  type        = string
  default     = null
}
variable "working_dir" {
  description = "Working directory inside the container"
  type        = string
  default     = null
}
variable "restart_policy" {
  description = "Restart policy. Default: no"
  type        = string
  default     = "no"
}
variable "privileged" {
  description = "Give extended privileges to this container"
  type        = bool
  default     = false
}
variable "network_mode" {
  description = "Specify a custom network mode"
  type        = string
  default     = null
}
variable "dns" {
  description = "Set custom dns servers for the container"
  type        = list(string)
  default     = null
}
variable "command" {
  description = "Override the default command"
  type        = list(string)
  default     = null
}
variable "ports" {
  description = "Expose ports"
  type = list(object({
    internal = number
    external = number
    protocol = string
  }))
  default = null
}
variable "named_volumes" {
  description = "Mount named volumes"
  type = map(object({
    container_path = string
    read_only      = bool
    create         = bool
  }))
  default = {}
}
variable "host_paths" {
  description = "Mount host paths"
  type = map(object({
    container_path = string
    read_only      = bool
  }))
  default = {}
}
variable "volumes_from_containers" {
  description = "Mount volumes from another container"
  type        = list(any)
  default     = null
}
variable "devices" {
  description = "Device mappings"
  type = map(object({
    container_path = string
    permissions    = string
  }))
  default = {}
}
variable "capabilities" {
  description = "Add or drop container capabilities"
  type = object({
    add  = list(string)
    drop = list(string)
  })
  default = null
}
variable "networks_advanced" {
  description = "Advanced network options for the container"
  type = object({
    name         = string
    aliases      = list(string)
    ipv4_address = string
    ipv6_address = string
  })
  default = null
}
variable "healthcheck" {
  description = "Test to check if container is healthy"
  type = object({
    interval     = string
    retries      = number
    start_period = string
    test         = list(string)
    timeout      = string
  })
  default = null
}
variable "environment" {
  description = "Add environment variables"
  type        = map(string)
  default     = null
}
variable "docker_networks" {
  description = "List of custom networks to create"
  type = map(object({
    ipam_config = object({
      aux_address = map(string)
      gateway     = string
      subnet      = string
    })
  }))
  default = {}
}
variable "create_before_destroy" {
  description = "Create container before destroy"
  type        = bool
  default     = false
}