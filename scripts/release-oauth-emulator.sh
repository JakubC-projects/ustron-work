#!/bin/bash

TAG=europe-west4-docker.pkg.dev/work-app-411215/main/oauth-emulator

docker build oauth-emulator --tag=$TAG

docker push $TAG

gcloud run deploy oauth-emulator \
    --project work-app-411215 \
    --image $TAG \
    --region europe-west4