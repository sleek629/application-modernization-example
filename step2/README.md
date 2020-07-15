# Step 2: Golang + MySQL

In this step, Golang + MySQL is used.

## Environment

- OS : Ubuntu 20.04 LTS on GCP
- Golang : 1.13.12
- MySQL : 8.0.20

## Setup

Run this command on Ubuntu.

```bash
# Install Golang v1.13.12
wget https://golang.org/dl/go1.13.12.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.13.12.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
mkdir $HOME/go
echo "export GOPATH=$HOME/go"

git clone https://github.com/sleek629/application-modernization-example.git
cd application-modernization-example/step2/src
# Install the required Go modules
go mod tidy

# Install MySQL server
sudo apt update
sudo apt install -y mysql-server
sudo mysql -u root < init.sql

# Run http server using Golang
go run main.go
```

 ## Access

If you use GCP, check machine's IP address by running following command on cloud shell.

```bash
gcloud compute instances list
```

If you don't use cloud services, simply running `ip addr`.



Then, access to http://IP:8080.

