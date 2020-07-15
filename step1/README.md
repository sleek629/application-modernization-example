# Step 1: LAMP stack

In this step, LAMP stack is used.

## Environment

- OS : Ubuntu 20.04 LTS on GCP
- Web server : Apache v2.4.41
- MySQL : 8.0.20
- PHP : 7.4.3

## Setup

Run this command Ubuntu.

```bash
# Install LAMP stack
sudo apt update
sudo apt install tasksel
sudo tasksel install lamp-server
git clone https://github.com/sleek629/application-modernization-example.git
cd application-modernization-example/step1/src

# Place php file to document root
sudo mv index.php /var/www/html/
sudo rm /var/www/html/index.html

# Initialize the MySQL database
sudo mysql -u root < init.sql
```

 ## Access

If you use GCP, check machine's IP address by running following command on cloud shell.

```bash
gcloud compute instances list
```

If you don't use cloud services, simply running `ip addr`.



Then, access to http://IP.