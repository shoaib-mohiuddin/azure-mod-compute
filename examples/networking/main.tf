module "main" {

  source = "git::https://github.com/shoaib-mohiuddin/azure-mod-network"

  vnet_resource_group_name = "${var.name}-rg"
  location                 = "North Europe"
  address_space            = ["192.168.0.0/23"]
  subnet_address_prefixes  = ["192.168.0.0/27"]
  virtual_network_name     = "${var.name}-main-vnet"
  subnet_name              = "${var.name}-VmSubnet"
}
