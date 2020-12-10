variable "container_name" {
  type = string
}

variable "network_name" {
  type = string
}

variable "gateway" {
  type = string
}

variable "ipv4_address" {
  type = string
}

variable "subnet" {
  type = string
}

variable "ports" {
  type    = list
  default = null
}

variable "named_volumes" {
  type    = map
  default = {}
}

variable "host_paths" {
  type    = map
  default = {}
}

variable "devices" {
  type    = map
  default = {}
}
