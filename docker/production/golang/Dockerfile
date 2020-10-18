FROM golang:1.15-alpine AS build
ADD ./ /src

RUN apk add make gcc musl-dev && cd /src && make clean all

FROM alpine
WORKDIR /app
COPY --from=build /src/app /app/
ENTRYPOINT ./app
