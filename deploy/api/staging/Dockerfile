FROM registry.gitlab.com/salt-and-pepper/crox-and-t-survey/backend-services/base:latest as builder

ENV APP_USER apps
ENV APP_SERVICE api
ENV APP_HOME /app

WORKDIR $APP_HOME

USER $USER

COPY --chown=$APP_USER . .

ENV GO111MODULE on

RUN go mod vendor
RUN go mod verify

RUN go build $APP_HOME/services/$APP_SERVICE/main.go

FROM gcr.io/distroless/base-debian10

ENV APP_HOME /app
WORKDIR $APP_HOME

COPY --from=builder $APP_HOME/main $APP_HOME/main

EXPOSE 80 443

CMD ["./main"]
