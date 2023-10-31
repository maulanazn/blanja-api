FROM golang:1.21.3 as builder

ARG APP_DIR=/userboilerplate-api

ARG CONFIG_DIR=/config
ARG ENTITY_DIR=/entity
ARG PUBLIC_DIR=/public
ARG REPOSITORY_DIR=/repository
ARG TABLEDATA_DIR=/tabledata
ARG TEST_DIR=/test
ARG TMP_DIR=/tmp
ARG WEBSERVER_DIR=/webserver
ARG CONTROLLER_DIR=/webserver/controller
ARG HELPER_DIR=/webserver/helper
ARG MIDDLEWARE_DIR=/webserver/middleware
ARG REQUEST_DIR=/webserver/request
ARG RESPONSE_DIR=/webserver/response
ARG SERVICE_DIR=/webserver/service

WORKDIR ${APP_DIR}${CONFIG_DIR}
COPY config/* ${APP_DIR}${CONFIG_DIR}/
RUN go get github.com/spf13/viper
RUN go get gorm.io/driver/postgres
RUN go get gorm.io/gorm

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
RUN go get github.com/joho/godotenv
RUN go get github.com/stretchr/testify/assert

WORKDIR ${APP_DIR}${TMP_DIR}
COPY tmp/* ${APP_DIR}${TMP_DIR}/

WORKDIR ${APP_DIR}${WEBSERVER_DIR}
COPY webserver/* ${APP_DIR}${WEBSERVER_DIR}/

WORKDIR ${APP_DIR}${CONTROLLER_DIR}
COPY webserver/controller/* ${APP_DIR}${CONTROLLER_DIR}/

WORKDIR ${APP_DIR}${HELPER_DIR}
COPY webserver/helper/* ${APP_DIR}${HELPER_DIR}/
RUN go get github.com/cloudinary/cloudinary-go/v2
RUN go get github.com/cloudinary/cloudinary-go/v2/api/uploader
RUN go get github.com/golang-jwt/jwt
RUN go get github.com/google/uuid
RUN go get golang.org/x/crypto/bcrypt

WORKDIR ${APP_DIR}${MIDDLEWARE_DIR}
COPY webserver/middleware/* ${APP_DIR}${MIDDLEWARE_DIR}/

WORKDIR ${APP_DIR}${REQUEST_DIR}
COPY webserver/request/* ${APP_DIR}${REQUEST_DIR}/
RUN go get github.com/go-ozzo/ozzo-validation
RUN go get github.com/go-ozzo/ozzo-validation/is

WORKDIR ${APP_DIR}${RESPONSE_DIR}
COPY webserver/response/* ${APP_DIR}${RESPONSE_DIR}/

WORKDIR ${APP_DIR}${SERVICE_DIR}
COPY webserver/service/* ${APP_DIR}${SERVICE_DIR}/
RUN go get github.com/albrow/forms

WORKDIR ${APP_DIR}
COPY main.go config.env .env.example .gitignore go.work go.work.sum LICENSE Makefile README.md ./
RUN go build -o ${APP_DIR}/main .

FROM alpine:latest

WORKDIR /userboilerplate-api/

COPY --from=builder /userboilerplate-api/userboilerplate-api ./

CMD /userboilerplate-api/userboilerplate-api
