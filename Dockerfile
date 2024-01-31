# Build Stage
FROM golang:1.21.6-bullseye AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build \
-ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
-X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
-o main

# Deploy stage

FROM gcr.io/distroless/base-debian11 AS deploy

COPY --from=build /app/main /main
# COPY --from=build /app/.env.dev /.env
# COPY --from=build /app/.env.uat /.env
# COPY --from=build /app/.env.prod /.env

EXPOSE 8080

USER nonroot:nonroot

CMD [ "/main" ]