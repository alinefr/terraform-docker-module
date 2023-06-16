variable "image" {
  type = string
}

variable "container_name" {
  type = string
}

variable "hostname" {
  type    = string
  default = null
}

variable "restart_policy" {
  type    = string
  default = null
}

variable "working_dir" {
  type    = string
  default = null
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

variable "entrypoint" {
  type    = list(string)
  default = null
}

variable "command" {
  type    = list(string)
  default = null
}

variable "ports" {
  type    = list(any)
  default = null
}

variable "named_volumes" {
  type    = map(any)
  default = {}
}

variable "host_paths" {
  type    = map(any)
  default = {}
}

variable "devices" {
  type    = map(any)
  default = {}
}

variable "capabilities" {
  type    = map(any)
  default = null
}

variable "environment" {
  type    = map(string)
  default = null
}

variable "networks_advanced" {
  type    = list(any)
  default = []
}

variable "docker_networks" {
  type    = list(any)
  default = []
}
