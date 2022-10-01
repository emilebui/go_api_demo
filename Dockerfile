FROM golang:alpine

ARG app_name=app
RUN mkdir -p /opt/${app_name}
ADD ./go.mod /opt/${app_name}
ADD ./go.sum /opt/${app_name}
ADD ca-certificates.crt /etc/ssl/certs/

RUN apk add build-base

WORKDIR /opt/${app_name}

ENV http_proxy=http://192.168.5.8:3128
ENV https_proxy=http://192.168.5.8:3128

ADD ./main.go /opt/${app_name}
ADD ./db_conn /opt/${app_name}/db_conn
ADD ./helpers /opt/${app_name}/helpers
ADD ./load_config /opt/${app_name}/load_config
ADD ./logging /opt/${app_name}/logging
ADD ./models /opt/${app_name}/models
ADD ./proto /opt/${app_name}/proto
ADD ./proto_gen /opt/${app_name}/proto_gen
ADD ./service /opt/${app_name}/service

RUN go mod download && \
unset http_proxy && \
unset https_proxy && \
CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o go_app .

FROM alpine

ARG app_name=app
ENV TZ=Asia/Ho_Chi_Minh

WORKDIR /app/go_api_demo
COPY --from=0 /opt/${app_name}/go_app /app/go_api_demo/go_app
COPY --from=0 /opt/${app_name}/load_config/error_messages.json /app/go_api_demo/load_config/error_messages.json
COPY --from=0 /opt/${app_name}/load_config/config.yaml /app/go_api_demo/load_config/config.yaml

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["/app/go_api_demo/go_app"]

