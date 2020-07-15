# Step 2.5: GAE(Golang) + Cloud SQL

In this step, GAE + Cloud SQL is used.

## Environment

- PaaS : Google App Engine (GAE) Go 1.13 standard
- Database : MySQL v5.7 on Cloud SQL

## Setup

Run this command on Cloud shell to setup Cloud SQL.

```bash
# Create database in Tokyo region
gcloud sql instances create word-database --tier=db-n1-standard-1 --region=asia-northeast1
# Connect to database (If asked for a password, press enter)
gcloud sql connect word-database --user root
# Run command in init.sql on mysql console

# Run following command in bash
git clone https://github.com/sleek629/application-modernization-example.git
cd application-modernization-example/step2.5/src
gcloud sql instances describe word-database --flatten=connectionName
# Open app.yaml and replace CONNECTION_NAME with output above (PROJECTNAME:REGION:DATABASE_NAME)
vi app.yaml
# Deploy application
gcloud app deploy
```



 ## Access

You can access the application in a browser by simply running following command on your local PC.

```bash
gcloud app browse
```



