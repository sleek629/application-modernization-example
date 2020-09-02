# Step 6.5 App using Vuejs in frontend

In this step, Vue is used in frontend.

The application is devided into 4 components.

- web (Vue): This is front-end application. This application gets data from wordservice using gRPC client.
- wordservice (Golang): This is back-end application using gRPC.
- db: This is database. We can choose to use MySQL or On-Memory databases.
- proxy: This is used to redirects gRPC request from web to wordservice.

## Environment

- OS : Container-Optimized OS 81-12871.1190.0 stable on GCP
- Docker: 19.03.6
- Docker-compose: 1.26.2 (Using Docker)
- Golang : 1.13.12
- node: 14.9.0
- MySQL : 8.0

## Setup

Run this command on Container-Optimized OS.

```bash
# Clone source files
git clone https://github.com/sleek629/application-modernization-example.git
cd application-modernization-example/step6.5/src
# Change IP address to instance global IP in docker-compose file
sed -i "s/INSTANCE_IP/`curl 'http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip' -H 'Metadata-Flavor: Google'`/g" docker-compose-*.yml
```

### Using MySQL

```
# Run http server using Docker compose
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "$PWD:$PWD" -w="$PWD" docker/compose:1.26.2 -f docker-compose-1.yml up
```

### Using On-Memory DB

If you start application this way, you don't have to prepare MySQL DB.

```
# Run http server using Docker compose
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "$PWD:$PWD" -w="$PWD" docker/compose:1.26.2 -f docker-compose-2.yml up
```

 ## Access

If you use GCP, check machine's IP address by running following command on cloud shell.

After that, you have to allow connection to tcp:8080.

```bash
gcloud compute instances list
gcloud compute firewall-rules create allow-http-8080 --allow=tcp:8080 --network default --direction ingress --priority 1000
```

If you don't use cloud services, simply running `ip addr`.



Then, access to http://IP:8080.
