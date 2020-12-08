terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.7"
    }
  }
  required_version = ">= 0.13"
}
