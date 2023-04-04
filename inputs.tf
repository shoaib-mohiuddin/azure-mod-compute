variable "location" {
  description = "Location of the resources"
  type        = string
  default     = ""
}

variable "rg_name" {
  description = "name of the resource group"
  type        = string
  default     = ""
}

variable "vm_name" {
  description = "name of the virtual machine"
  type        = string
  default     = "webserver"
}

variable "vm_size" {
  description = "size of the virtual machine"
  type        = string
}

variable "vm_subnet_id" {
  description = "subnet id to deploy the vm"
  type        = string
  default     = ""
}

variable "tags" {
  description = "Map of tags that will be applied to the VNet"
  type        = map(string)
  default     = {}
}
