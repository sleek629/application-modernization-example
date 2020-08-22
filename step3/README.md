# Step 3: App on Docker

In this step, Docker is used to build Golang + MySQL environment.

## Environment

- OS : Container-Optimized OS 81-12871.1160.0 stable on GCP
- Docker: 19.03.6
- Docker-compose: 1.26.2 (Using Docker)
- Golang : 1.13.12
- MySQL : 8.0

## Setup

Run this command on Container-Optimized OS.

```bash
# Clone source files
git clone https://github.com/sleek629/application-modernization-example.git
cd application-modernization-example/step3/src

# Run http server using Docker compose
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "$PWD:$PWD" -w="$PWD" docker/compose:1.26.2 up
```

 ## Access

If you use GCP, check machine's IP address by running following command on cloud shell.

```bash
gcloud compute instances list
```

If you don't use cloud services, simply running `ip addr`.



Then, access to http://IP:8080.

