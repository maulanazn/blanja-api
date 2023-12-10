FROM golang:1.21.3-alpine3.18 as builder

LABEL author="Maulana Zulkifli N"
LABEL email="maulanarpl08@gmail.com"

ARG APP_DIR=/blanja-api

ARG CONFIG_DIR=/config
ARG LOG_DIR=/log
ARG MIDDLEWARE_DIR=/middleware
ARG PUBLIC_DIR=/public
ARG ADDR_DIR=/src/address
ARG ATTR_DIR=/src/attribute
ARG AUTOLOAD_DIR=/src/autoload
ARG ORDER_DIR=/src/order_details
ARG PRODUCT_DIR=/src/product
ARG USER_DIR=/src/users
ARG WISHLIST_DIR=/src/wishlist
ARG TABLEDATA_DIR=/tabledata
ARG HELPER_DIR=/util

EXPOSE 3000

WORKDIR ${APP_DIR}${CONFIG_DIR}
COPY config/* ${APP_DIR}${CONFIG_DIR}/
RUN go mod tidy

WORKDIR ${APP_DIR}${LOG_DIR}
COPY log/* ${APP_DIR}${LOG_DIR}/

WORKDIR ${APP_DIR}${MIDDLEWARE_DIR}
COPY middleware/* ${APP_DIR}${MIDDLEWARE_DIR}/
RUN go mod tidy

WORKDIR ${APP_DIR}${PUBLIC_DIR}
COPY public/* ${APP_DIR}${PUBLIC_DIR}/

WORKDIR ${APP_DIR}${ADDR_DIR}
COPY src/address ${APP_DIR}${ADDR_DIR}/

WORKDIR ${APP_DIR}${ATTR_DIR}
COPY src/attribute ${APP_DIR}${ATTR_DIR}/

WORKDIR ${APP_DIR}${AUTOLOAD_DIR}
COPY src/autoload ${APP_DIR}${AUTOLOAD_DIR}/

WORKDIR ${APP_DIR}${ORDER_DIR}
COPY src/order_details ${APP_DIR}${ORDER_DIR}/

WORKDIR ${APP_DIR}${PRODUCT_DIR}
COPY src/product ${APP_DIR}${PRODUCT_DIR}/

WORKDIR ${APP_DIR}${USER_DIR}
COPY src/users ${APP_DIR}${USER_DIR}/

WORKDIR ${APP_DIR}${WISHLIST_DIR}
COPY src/wishlist ${APP_DIR}${WISHLIST_DIR}/

WORKDIR ${APP_DIR}${TABLEDATA_DIR}
COPY tabledata/* ${APP_DIR}${TABLEDATA_DIR}/

WORKDIR ${APP_DIR}${HELPER_DIR}
COPY util/* ${APP_DIR}${HELPER_DIR}/
RUN go mod tidy

WORKDIR ${APP_DIR}
COPY main.go go.work go.work.sum .env .env.example LICENSE README.md blanja.yaml blanja.yaml.example ./
RUN go build -o ${APP_DIR}/main main.go

FROM alpine:latest

WORKDIR /blanja-api/

COPY --from=builder /blanja-api/main ./

CMD /blanja-api/main
