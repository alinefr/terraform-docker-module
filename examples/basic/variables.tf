variable "container_name" {
  type = string
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

variable "environment" {
  type    = map(string)
  default = null
}

variable "networks_advanced" {
  type    = list(any)
  default = []
}

variable "docker_networks" {
  type    = map(any)
  default = {}
}
