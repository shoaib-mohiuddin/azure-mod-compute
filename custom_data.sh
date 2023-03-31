#! /bin/bash
sudo apt-get update
sudo apt-get install -y apache2
#echo <h1>Hello, World!</h1> > /var/www/html/index.html
echo \<center\>\<h1\>Hello, World!\</h1\>\<br/\>\</center\> > /var/www/html/index.html
sudo service start apache2
sudo service enable apache2
