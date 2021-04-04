#!/bin/bash

# Debug env
/app/beetle-actions version
/app/beetle version

# Load github actions data as env vars
/app/beetle-actions trigger

if [ $BEETLE_WATCH_DEPLOYMENT ]
then
    # Watch deployment progress
    /app/beetle deploy --api_url $BEETLE_API_URL --api_key $BEETLE_API_KEY --cluster $BEETLE_CLUSTER_NAME --namespace $BEETLE_NAMESPACE_NAME -a $BEETLE_APP_ID -s $BEETLE_DEPLOYMENT_STRATEGY -v $BEETLE_APP_VERSION --max_surge $BEETLE_MAX_SURGE --max_unavailable $BEETLE_MAX_UNAVAILABLE -w

else
    # Trigger deployment request
    /app/beetle deploy --api_url $BEETLE_API_URL --api_key $BEETLE_API_KEY --cluster $BEETLE_CLUSTER_NAME --namespace $BEETLE_NAMESPACE_NAME -a $BEETLE_APP_ID -s $BEETLE_DEPLOYMENT_STRATEGY -v $BEETLE_APP_VERSION --max_surge $BEETLE_MAX_SURGE --max_unavailable $BEETLE_MAX_UNAVAILABLE
fi
