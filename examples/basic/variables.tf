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
