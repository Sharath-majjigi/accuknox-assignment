FROM golang:1.17-alpine as build-env

ENV APP_NAME accuknox-assignment
ENV CMD_PATH main.go

COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN CGO_ENABLED=O go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

FROM alpine:3.14

ENV APP_NAME accuknox-assignment

COPY --from=build-env /$APP_NAME .

EXPOSE 8088

CMD ./$APP_NAME