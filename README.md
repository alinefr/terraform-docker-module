# Terraform Docker Module

## Overview

Terraform docker module is a module to help docker maintenance over terraform. 
It should replace other means of docker maintenance like docker-compose.

There are several advantages of maintaining docker on terraform.

* Infrastructure as code. You don't need to manually ssh into servers
* CI/CD. Many CI tools offers some way to automate terraform execution.
* Remote execution. You don't need to manually ssh into servers.

This module uses under the hood [Docker Provider](https://www.terraform.io/docs/providers/docker/index.html).

## Aditional documentation

Full documentation for [Terraform Docker Module](https://github.com/alinefr/terraform-docker-module/blob/master/modules/docker/README.md).

