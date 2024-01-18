#!/bin/bash

TAG=europe-west4-docker.pkg.dev/work-app-411215/main/workapp-dev

docker build . --tag=$TAG

docker push $TAG

gcloud run deploy workapp-dev \
    --project work-app-411215 \
    --image $TAG \
    --region europe-west4
