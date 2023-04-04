module "main" {

  source = "../.."

  location     = "North Europe"
  rg_name      = "${var.name}-rg"
  vm_name      = var.name
  vm_subnet_id = var.vm_subnet_id
  vm_size      = "Standard_DS1_v2"
}
