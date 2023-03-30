output "webserver_public_ip" {
  value = azurerm_linux_virtual_machine.webserver.public_ip_address
}
