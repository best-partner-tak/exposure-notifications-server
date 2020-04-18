#!/bin/sh

export PROJECT_ID=$(gcloud config get-value core/project)

# environment variables for running the services
export DATASTORE_PROJECT_ID=$PROJECT_ID
export DIAGNOSIS_KMS_KEY="projects/$PROJECT_ID/locations/us/keyRings/us-db-keys/cryptoKeys/global-diagnosiskeys"

# wipeout variables
export TTL_DURATION="14d"

# local application credentials - you need to get your own credentials
export GOOGLE_APPLICATION_CREDENTIALS="$(pwd)/local/sa.json"

if [ ! -f "$GOOGLE_APPLICATION_CREDENTIALS" ]; then
    echo "$GOOGLE_APPLICATION_CREDENTIALS does not exist. \
Use https://console.cloud.google.com/iam-admin/serviceaccounts/create?project=$PROJECT_ID to create a service account \
with Datastore->Cloud Datastore User, then create a key and download the JSON file and store it at \
$GOOGLE_APPLICATION_CREDENTIALS"
    exit -1
fi

echo "Project ID:  $PROJECT_ID"
echo "Credentials: $GOOGLE_APPLICATION_CREDENTIALS"
