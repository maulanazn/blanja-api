FROM golang:1.21.3-alpine3.18 as builder

LABEL author="Maulana Zulkifli N"
LABEL email="maulanarpl08@gmail.com"

ARG APP_DIR=/userboilerplate-api

ARG CONFIG_DIR=/config
ARG ENTITY_DIR=/entity
ARG PUBLIC_DIR=/public
ARG REPOSITORY_DIR=/repository
ARG TABLEDATA_DIR=/tabledata
ARG TEST_DIR=/test
ARG TMP_DIR=/tmp
ARG CONTROLLER_DIR=/controller
ARG HELPER_DIR=/helper
ARG MIDDLEWARE_DIR=/middleware
ARG REQUEST_DIR=/request
ARG RESPONSE_DIR=/response
ARG SERVICE_DIR=/service

EXPOSE 3000

WORKDIR ${APP_DIR}${CONFIG_DIR}
COPY config/* ${APP_DIR}${CONFIG_DIR}/
RUN go mod download

WORKDIR ${APP_DIR}${ENTITY_DIR}
COPY entity/* ${APP_DIR}${ENTITY_DIR}/

WORKDIR ${APP_DIR}${PUBLIC_DIR}
COPY public/* ${APP_DIR}${PUBLIC_DIR}/

WORKDIR ${APP_DIR}${REPOSITORY_DIR}
COPY repository/* ${APP_DIR}${REPOSITORY_DIR}/

WORKDIR ${APP_DIR}${TABLEDATA_DIR}
COPY tabledata/* ${APP_DIR}${TABLEDATA_DIR}/

WORKDIR ${APP_DIR}${TEST_DIR}
COPY test/* ${APP_DIR}${TEST_DIR}/
RUN go mod download

WORKDIR ${APP_DIR}${TMP_DIR}
COPY tmp/* ${APP_DIR}${TMP_DIR}/

WORKDIR ${APP_DIR}${CONTROLLER_DIR}
COPY controller/* ${APP_DIR}${CONTROLLER_DIR}/

WORKDIR ${APP_DIR}${HELPER_DIR}
COPY helper/* ${APP_DIR}${HELPER_DIR}/
RUN go mod download

WORKDIR ${APP_DIR}${MIDDLEWARE_DIR}
COPY middleware/* ${APP_DIR}${MIDDLEWARE_DIR}/

WORKDIR ${APP_DIR}${REQUEST_DIR}
COPY request/* ${APP_DIR}${REQUEST_DIR}/
RUN go mod download

WORKDIR ${APP_DIR}${RESPONSE_DIR}
COPY response/* ${APP_DIR}${RESPONSE_DIR}/

WORKDIR ${APP_DIR}${SERVICE_DIR}
COPY service/* ${APP_DIR}${SERVICE_DIR}/
RUN go mod download

WORKDIR ${APP_DIR}
COPY main.go go.work go.work.sum .env LICENSE Makefile README.md ./
RUN go build -o ${APP_DIR}/main main.go

FROM alpine:latest

WORKDIR /userboilerplate-api/

COPY --from=builder /userboilerplate-api/main ./

CMD /userboilerplate-api/main

