#! /bin/bash
sudo apt-get update
sudo apt-get install -y apache2
echo <html><body><h1>Hello, World!<h1></body></html> > /var/www/html/index.html
sudo service start apache2
sudo service enable apache2
