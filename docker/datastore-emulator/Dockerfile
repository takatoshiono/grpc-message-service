FROM google/cloud-sdk:alpine

ENV DATASTORE_DATA_DIR /opt/datastore
ENV DATASTORE_PORT 8081

RUN apk --update --no-cache add openjdk8-jre
RUN gcloud components install beta cloud-datastore-emulator

COPY start.sh .

VOLUME ${DATASTORE_DATA_DIR}

EXPOSE ${DATASTORE_PORT}
ENTRYPOINT ["./start.sh"]
