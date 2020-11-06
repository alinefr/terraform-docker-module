terraform {
  required_providers {
    docker = {
      source  = "terraform-providers/docker"
      version = "~> 2.7"
    }
  }
  required_version = ">= 0.13"
}
