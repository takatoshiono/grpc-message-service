#!/bin/sh

gcloud beta emulators datastore start \
    --data-dir=${DATASTORE_DATA_DIR} \
    --host-port=0.0.0.0:${DATASTORE_PORT}
